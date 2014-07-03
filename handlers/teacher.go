package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateTeacher ...
func CreateTeacher(c *gin.Context) {
	t := new(m.Teacher)

	errs := binding.Bind(c.Req, t)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.Teachers.Store(t)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	t.ID = id

	c.JSON(201, &APIRes{"teacher": []m.Teacher{*t}})
	return
}

// GetTeacher ...
func GetTeacher(c *gin.Context) {

	id := c.Params.ByName("id")

	t := m.Teacher{}
	err := store.Teachers.FindByID(&t, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"teacher": []m.Teacher{t}})
}

// UpdateTeacher ...
func UpdateTeacher(c *gin.Context) {

	id := c.Params.ByName("id")

	t := new(m.Teacher)

	errs := binding.Bind(c.Req, t)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	t.ID = id
	err := store.Teachers.Update(t, id)

	if err != nil {
		writeError(c.Writer, "Error updating Teacher", 500, err)
		return
	}

	c.JSON(200, &APIRes{"teacher": []m.Teacher{*t}})
	return
}

// GetAllTeachers ...
func GetAllTeachers(c *gin.Context) {
	teachers := []m.Teacher{}
	err := store.Teachers.FindAll(&teachers)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"teacher": teachers})
	return
}
