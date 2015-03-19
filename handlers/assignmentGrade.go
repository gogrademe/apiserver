package handlers

import (
	"errors"

	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignmentGrade ...
func CreateAssignmentGrade(c *gin.Context) {
	a := new(m.AssignmentGrade)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	ids, err := store.AssignmentGrades.Insert(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = ids[0]

	c.JSON(201, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAssignmentGrade ...
func GetAssignmentGrade(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.AssignmentGrade{}
	err := store.AssignmentGrades.One(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"grade": []m.AssignmentGrade{a}})
	return
}

// UpdateAssignmentGrade ...
func UpdateAssignmentGrade(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.AssignmentGrade)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	a.ID = id
	err := store.AssignmentGrades.Update(a, id)

	if err != nil {
		writeError(c.Writer, "Error updating AssignmentGrade", 500, err)
		return
	}

	c.JSON(200, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAllAssignmentGrades ...
func GetAllAssignmentGrades(c *gin.Context) {
	filter := map[string]string{}
	if c.Request.URL.Query().Get("assignmentId") != "" {
		filter["assignmentId"] = c.Request.URL.Query().Get("assignmentId")
	}

	grades := []m.AssignmentGradeAPIRes{}
	q := store.AssignmentGrades.Filter(filter)

	q = q.EqJoin("assignmentId", r.Table("assignments"))

	// FIXME: Fix this!
	q = q.Map(func(row r.Term) r.Term {
		return row.Field("left").Merge(map[string]interface{}{
			"assignment": row.Field("right"),
			// Specifically this part.
			"gradeAverage": 3,
		})
	})
	err := store.DB.All(&grades, q)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"grade": grades})
	return
}
