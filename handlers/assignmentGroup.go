package handlers

import (
	"errors"

	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignmentGroup ...
func CreateAssignmentGroup(c *gin.Context) {
	a := new(m.AssignmentGroup)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	id, err := store.AssignmentGroups.Store(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = id

	c.JSON(201, &APIRes{"assignmentGroup": []m.AssignmentGroup{*a}})
	return
}

// GetAssignmentGroup ...
func GetAssignmentGroup(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.AssignmentGroup{}
	err := store.AssignmentGroups.FindByID(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"assignmentGroup": []m.AssignmentGroup{a}})
	return
}

// UpdateAssignmentGroup ...
func UpdateAssignmentGroup(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.AssignmentGroup)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	a.ID = id
	err := store.AssignmentGroups.Update(a, id)

	if err != nil {
		writeError(c.Writer, "Error updating AssignmentGroup", 500, err)
		return
	}

	c.JSON(200, &APIRes{"assignmentGroup": []m.AssignmentGroup{*a}})
	return
}

// GetAllAssignmentGroups ...
func GetAllAssignmentGroups(c *gin.Context) {
	filter := map[string]string{
		"classId": c.Request.URL.Query().Get("classId"),
		"termId":  c.Request.URL.Query().Get("termId"),
	}

	types := []m.AssignmentGroup{}
	err := store.AssignmentGroups.Filter(&types, filter)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"assignmentGroup": types})
	return
}

// DeleteAssignmentGroup ...
func DeleteAssignmentGroup(c *gin.Context) {

	id := c.Params.ByName("id")

	err := store.AssignmentGroups.Delete(id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"assignmentGroup": []m.AssignmentGroup{}})
	return
}
