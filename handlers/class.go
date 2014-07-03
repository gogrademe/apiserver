package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllClasses ...
func GetAllClasses(c *gin.Context) {
	classes := []m.Class{}
	err := store.Classes.FindAll(&classes)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"class": classes})
	return
}

// GetClass ...
func GetClass(c *gin.Context) {

	id := c.Params.ByName("id")
	class := m.Class{}
	err := store.Classes.FindByID(&class, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"class": []m.Class{class}})
	return
}

//CreateClass ...
func CreateClass(c *gin.Context) {
	class := new(m.Class)

	errs := binding.Bind(c.Req, class)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.Classes.Store(class)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	class.ID = id

	c.JSON(201, &APIRes{"class": []m.Class{*class}})
	return
}

//UpdateClass ...
func UpdateClass(c *gin.Context) {

	id := c.Params.ByName("id")

	class := new(m.Class)

	errs := binding.Bind(c.Req, class)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	class.ID = id
	err := store.Classes.Update(class, id)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"class": []m.Class{*class}})
	return
}
