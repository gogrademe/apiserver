package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllClasses ...
func GetAllClasses(c *gin.Context) {
	classes := []m.Class{}

	query := store.Classes.OrderBy(r.Asc("gradeLevel"), r.Asc("name"))
	err := store.DB.All(&classes, query)
	// err := store.Classes.FindAll(&classes)
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
	err := store.Classes.One(&class, id)
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

	errs := binding.Bind(c.Request, class)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	ids, err := store.Classes.Insert(class)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	class.ID = ids[0]

	c.JSON(201, &APIRes{"class": []m.Class{*class}})
	return
}

//UpdateClass ...
func UpdateClass(c *gin.Context) {

	id := c.Params.ByName("id")

	class := new(m.Class)

	errs := binding.Bind(c.Request, class)
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

// DeleteClass ...
func DeleteClass(c *gin.Context) {

	id := c.Params.ByName("id")

	_, err := store.DB.RunWrite(store.Classes.Get(id).Delete())
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"class": []m.Class{}})
	return
}
