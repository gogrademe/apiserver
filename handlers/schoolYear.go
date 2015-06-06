package handlers

import (
	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllSchoolYears ...
func GetAllSchoolYears(c *gin.Context) {
	schoolYears := []m.SchoolYear{}

	query := store.SchoolYears.OrderBy(r.Desc("startYear"), r.Desc("endYear"))
	err := store.DB.All(&schoolYears, query)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"schoolYear": schoolYears})
	return
}

//GetSchoolYear ...
func GetSchoolYear(c *gin.Context) {

	id := c.Param("id")
	schoolYear := m.SchoolYear{}

	//err := store.SchoolYears.FindByID(&schoolYear, id)
	err := store.SchoolYears.One(&schoolYear, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"schoolYear": []m.SchoolYear{schoolYear}})
	return
}

//CreateSchoolYear ...
func CreateSchoolYear(c *gin.Context) {
	schoolYear := new(m.SchoolYear)

	errs := binding.Bind(c.Request, schoolYear)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	id, err := store.SchoolYears.Insert(schoolYear)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	schoolYear.ID = id[0]

	c.JSON(201, &APIRes{"schoolYear": []m.SchoolYear{*schoolYear}})
	return
}

//UpdateSchoolYear ...
func UpdateSchoolYear(c *gin.Context) {

	id := c.Param("id")

	schoolYear := new(m.SchoolYear)

	errs := binding.Bind(c.Request, schoolYear)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	schoolYear.ID = id
	err := store.SchoolYears.Update(schoolYear, id)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"schoolYear": []m.SchoolYear{*schoolYear}})
	return
}
