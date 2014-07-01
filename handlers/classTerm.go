package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllClassTerms ...
func GetAllClassTerms(c *gin.Context) {
	terms := []m.ClassTerm{}
	err := store.ClassTerms.FindAll(&terms)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	writeJSON(c.Writer, &APIRes{"classTerm": terms})
	return
}

//GetClassTerm ...
func GetClassTerm(c *gin.Context) {

	id := c.Params.ByName("id")
	term := m.ClassTerm{}
	err := store.ClassTerms.FindByID(&term, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	writeJSON(c.Writer, &APIRes{"classTerm": []m.ClassTerm{term}})
	return
}

//CreateClassTerm ...
func CreateClassTerm(c *gin.Context) {
	term := new(m.ClassTerm)

	errs := binding.Bind(c.Req, term)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.ClassTerms.Store(term)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	term.ID = id

	writeJSON(c.Writer, &APIRes{"classTerm": []m.ClassTerm{*term}})
	return
}

//UpdateClassTerm ...
func UpdateClassTerm(c *gin.Context) {

	id := c.Params.ByName("id")

	term := new(m.ClassTerm)

	errs := binding.Bind(c.Req, term)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	term.ID = id
	err := store.ClassTerms.Update(term, id)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	writeJSON(c.Writer, &APIRes{"classTerm": []m.ClassTerm{*term}})
	return
}
