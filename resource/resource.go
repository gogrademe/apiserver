package resource

import (
	m "github.com/GoGradeMe/APIServer/model"
	"github.com/GoGradeMe/APIServer/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

type Resourcer interface {
	GetOne(c *gin.Context)
	GetAll(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
	Post(c *gin.Context)
}

type Resource struct{}

func (r *Resource) GetAll(c *gin.Context) {
	years := []m.SchoolYear{}

	query := store.SchoolYears.OrderBy(r.Desc("startYear"), r.Asc("endYear"))
	err := store.DB.All(&years, query)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"term": terms})
	return
}

func (r *Resource) GetOne(c *gin.Context) {

	id := c.Params.ByName("id")
	term := m.SchoolYear{}

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

func (r *Resource) Post(c *gin.Context) {
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

func (r *Resource) Put(c *gin.Context) {

	id := c.Params.ByName("id")

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
