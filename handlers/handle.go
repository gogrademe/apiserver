package handlers

import (
	"errors"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

func HandleDelete(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")

	_, err := store.DB.RunWrite(store.Assignments.Get(id).Delete())
	if err == store.ErrNotFound {
		rest.NotFound(w, r)
		return
	}
	if err != nil {
		rest.Error(w, serverError, 500)
		return
	}

	// w.WriteJson(&APIRes{"assignment": []m.Assignment{}})
	w.WriteHeader(http.StatusOK)
}

func HandleCreate(w rest.ResponseWriter, r *rest.Request) {
	a := new(m.Assignment)

	errs := binding.Bind(r.Request, a)
	if errs != nil {
		rest.Error(w, errs.Error(), StatusUnprocessable)
		return
	}

	ids, err := store.Courses.Insert(a)
	if err != nil {
		rest.Error(w, serverError, 500)
		return
	}
	a.ID = ids[0]

	w.WriteJson(&APIRes{"assignment": []m.Assignment{*a}})
	return
}

func HandleGet(c *gin.Context) {
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

func HandlePut(c *gin.Context) {
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

func HandleGetAll(c *gin.Context) {
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
