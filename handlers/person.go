package handlers

import (
	"errors"
	"log"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreatePerson ...
func CreatePerson(c *gin.Context) {
	p := new(m.Person)

	errs := binding.Bind(c.Req, p)
	if errs != nil {
		writeError(c.Writer, errs, 500, nil)
		return
	}

	id, err := store.People.Store(p)
	if err != nil {
		writeError(c.Writer, "Error creating Person", 500, err)
		return
	}

	p.ID = id

	c.JSON(200, &APIRes{"person": []m.Person{*p}})
	return
}

// UpdatePerson ...
func UpdatePerson(c *gin.Context) {

	id := c.Params.ByName("id")

	p := new(m.Person)

	errs := binding.Bind(c.Req, p)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	p.ID = id
	err := store.People.Update(p, id)

	if err != nil {
		writeError(c.Writer, "Error updating Person", 500, err)
		return
	}

	c.JSON(200, &APIRes{"person": []m.Person{*p}})
	return
}

// GetPerson ...
func GetPerson(c *gin.Context) {

	id := c.Params.ByName("id")

	p := m.Person{}
	err := store.People.FindByID(&p, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"person": []m.Person{p}})
	return
}

// GetAllPeople ...
func GetAllPeople(c *gin.Context) {
	log.Println(c.Req.URL.Query())
	log.Println(c.Get("user"))
	people, err := store.People.FindAll()
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	c.Error(errors.New("test"), "a")
	c.JSON(200, &APIRes{"person": people})
	return
}
