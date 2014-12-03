package handlers

import (
	"errors"

	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignmentType ...
func CreateAssignmentType(c *gin.Context) {
	a := new(m.AssignmentType)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	id, err := store.AssignmentTypes.Store(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = id

	c.JSON(201, &APIRes{"type": []m.AssignmentType{*a}})
	return
}

// GetAssignmentType ...
func GetAssignmentType(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.AssignmentType{}
	err := store.AssignmentTypes.FindByID(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"type": []m.AssignmentType{a}})
	return
}

// UpdateAssignmentType ...
func UpdateAssignmentType(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.AssignmentType)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	a.ID = id
	err := store.AssignmentTypes.Update(a, id)

	if err != nil {
		writeError(c.Writer, "Error updating AssignmentType", 500, err)
		return
	}

	c.JSON(200, &APIRes{"type": []m.AssignmentType{*a}})
	return
}

// GetAllAssignmentTypes ...
func GetAllAssignmentTypes(c *gin.Context) {
	types := []m.AssignmentType{}
	err := store.AssignmentTypes.FindAll(&types)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"type": types})
	return
}

// DeleteAssignmentType ...
func DeleteAssignmentType(c *gin.Context) {

	id := c.Params.ByName("id")

	err := store.AssignmentTypes.Delete(id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"type": []m.AssignmentType{}})
	return
}
