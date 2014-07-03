package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateStudent ...
func CreateStudent(c *gin.Context) {
	s := new(m.Student)

	errs := binding.Bind(c.Req, s)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.Students.Store(s)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	s.ID = id

	c.JSON(201, &APIRes{"student": []m.Student{*s}})
	return
}

// GetStudent ...
func GetStudent(c *gin.Context) {

	id := c.Params.ByName("id")

	s := m.Student{}
	err := store.Students.FindByID(&s, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"student": []m.Student{s}})
	return
}

// UpdateStudent ...
func UpdateStudent(c *gin.Context) {

	id := c.Params.ByName("id")

	s := new(m.Student)

	errs := binding.Bind(c.Req, s)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	s.ID = id
	err := store.Students.Update(s, id)

	if err != nil {
		writeError(c.Writer, "Error updating Student", 500, err)
		return
	}

	c.JSON(200, &APIRes{"student": []m.Student{*s}})
	return
}

// GetAllStudents ...
func GetAllStudents(c *gin.Context) {
	students := []m.Student{}
	err := store.Students.FindAll(&students)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"student": students})
	return
}
