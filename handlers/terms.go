package handlers

import (
	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllTerms ...
func GetAllTerms(c *gin.Context) {
	terms := []m.Term{}

	query := store.Terms.OrderBy(r.Desc("schoolYear"), r.Asc("startDate"), r.Asc("endDate"))
	err := store.DB.All(&terms, query)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"term": terms})
	return
}

//GetTerm ...
func GetTerm(c *gin.Context) {

	id := c.Param("id")
	term := m.Term{}

	//err := store.Terms.FindByID(&term, id)
	err := store.Terms.One(&term, id)
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

	errs := binding.Bind(c.Request, term)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.Terms.Insert(term)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	term.ID = id[0]

	c.JSON(201, &APIRes{"term": []m.Term{*term}})
	return
}

//UpdateTerm ...
func UpdateTerm(c *gin.Context) {

	id := c.Param("id")

	term := new(m.Term)

	errs := binding.Bind(c.Request, term)
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
