package handlers

import (
	m "github.com/GoGradeMe/APIServer/model"
	"github.com/GoGradeMe/APIServer/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreatePerson ...
func CreatePerson(c *gin.Context) {
	p := new(m.Person)

	errs := binding.Bind(c.Request, p)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	ids, err := store.People.Insert(p)
	if err != nil {
		writeError(c.Writer, "Error creating Person", 500, err)
		return
	}

	p.ID = ids[0]

	c.JSON(201, &APIRes{"person": []m.Person{*p}})
	return
}

// UpdatePerson ...
func UpdatePerson(c *gin.Context) {

	id := c.Params.ByName("id")

	p := new(m.Person)

	errs := binding.Bind(c.Request, p)
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
	// err := store.People.FindByID(&p, id)
	err := store.People.One(&p, id)
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

	// p, err := store.People.FindAll()
	p := []m.Person{}
	query := store.People.OrderBy("firstName", "middleName", "lastName")
	err := store.DB.All(&p, query)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	c.JSON(200, &APIRes{"person": p})
	return
}

// DeletePerson ...
func DeletePerson(c *gin.Context) {

	id := c.Params.ByName("id")

	_, err := store.DB.RunWrite(store.People.Get(id).Delete())
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"person": []m.Person{}})
	return
}
