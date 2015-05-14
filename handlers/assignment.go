package handlers

import (
	"errors"

	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// DeleteAssignment ...
func DeleteAssignment(c *gin.Context) {

	id := c.Params.ByName("id")

	_, err := store.DB.RunWrite(store.Assignments.Get(id).Delete())
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"assignment": []m.Assignment{}})
	return
}

// CreateAssignment ...
func CreateAssignment(c *gin.Context) {
	a := new(m.Assignment)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	ids, err := store.Courses.Insert(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = ids[0]

	c.JSON(201, &APIRes{"assignment": []m.Assignment{*a}})
	return
}

// GetAssignment ...
func GetAssignment(c *gin.Context) {
	var (
		id = c.Params.ByName("id")
		a  = m.Assignment{}
	)

	err := store.Assignments.One(&a, id)
	if err != nil {
		handleDBError(c.Writer, err)
		return
	}

	c.JSON(200, &APIRes{"assignment": []m.Assignment{a}})
	return
}

// UpdateAssignment ...
func UpdateAssignment(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.Assignment)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	a.ID = id
	err := store.Assignments.Update(a, id)

	if err != nil {
		writeError(c.Writer, "error updating Assignment", 500, err)
		return
	}

	c.JSON(200, &APIRes{"assignment": []m.Assignment{*a}})
	return
}

// GetAllAssignments ...
func GetAllAssignments(c *gin.Context) {
	filter := map[string]string{}
	if c.Request.URL.Query().Get("classId") != "" {
		filter["classId"] = c.Request.URL.Query().Get("classId")
	}
	if c.Request.URL.Query().Get("termId") != "" {
		filter["termId"] = c.Request.URL.Query().Get("termId")
	}
	if c.Request.URL.Query().Get("typeId") != "" {
		filter["typeId"] = c.Request.URL.Query().Get("typeId")
	}

	assignments := []m.AssignmentAPIRes{}

	query := store.Assignments.Filter(filter).OrderBy("dueDate", "name")
	query = query.EqJoin("groupId", r.Table("assignmentGroups"))
	query = query.Map(func(row r.Term) r.Term {
		return row.Field("left").Merge(map[string]interface{}{
			"group": row.Field("right"),
		})
	})
	err := store.DB.All(&assignments, query)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"assignment": assignments})
	return
}
