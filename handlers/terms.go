package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllTerms ...
func GetAllTerms(c *gin.Context) {
	terms := []m.Term{}
	err := store.Terms.FindAll(&terms)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"term": terms})
	return
}

//GetTerm ...
func GetTerm(c *gin.Context) {

	id := c.Params.ByName("id")
	term := m.Term{}
	err := store.Terms.FindByID(&term, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"term": []m.Term{term}})
	return
}

//CreateTerm ...
func CreateTerm(c *gin.Context) {
	term := new(m.Term)

	errs := binding.Bind(c.Req, term)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.Terms.Store(term)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	term.ID = id

	c.JSON(201, &APIRes{"term": []m.Term{*term}})
	return
}

//UpdateTerm ...
func UpdateTerm(c *gin.Context) {

	id := c.Params.ByName("id")

	term := new(m.Term)

	errs := binding.Bind(c.Req, term)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	term.ID = id
	err := store.Terms.Update(term, id)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"term": []m.Term{*term}})
	return
}
