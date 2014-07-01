package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignmentGrade ...
func CreateAssignmentGrade(c *gin.Context) {
	a := new(m.AssignmentGrade)

	errs := binding.Bind(c.Req, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.AssignmentGrades.Store(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = id

	writeJSON(c.Writer, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAssignmentGrade ...
func GetAssignmentGrade(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.AssignmentGrade{}
	err := store.AssignmentGrades.FindByID(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	writeJSON(c.Writer, &APIRes{"grade": []m.AssignmentGrade{a}})
	return
}

// UpdateAssignmentGrade ...
func UpdateAssignmentGrade(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.AssignmentGrade)

	errs := binding.Bind(c.Req, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	a.ID = id
	err := store.AssignmentGrades.Update(a, id)

	if err != nil {
		writeError(c.Writer, "Error updating AssignmentGrade", 500, err)
		return
	}

	writeJSON(c.Writer, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAllAssignmentGrades ...
func GetAllAssignmentGrades(c *gin.Context) {
	grades := []m.AssignmentGrade{}
	err := store.Classes.FindAll(&grades)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	writeJSON(c.Writer, &APIRes{"grade": grades})
	return
}
