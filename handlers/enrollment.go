package handlers

import (
	"errors"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateEnrollment ...
func CreateEnrollment(c *gin.Context) {
	a := new(m.Enrollment)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(400, errs)
		return
	}

	id, err := store.Enrollments.Store(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = id

	c.JSON(201, &APIRes{"enrollment": []m.Enrollment{*a}})
	return
}

// GetEnrollment ...
func GetEnrollment(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.Enrollment{}
	err := store.Enrollments.FindByID(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"enrollment": []m.Enrollment{a}})
	return
}

// DeleteEnrollment ...
func DeleteEnrollment(c *gin.Context) {

	id := c.Params.ByName("id")

	err := store.Enrollments.Delete(id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"enrollment": []m.Enrollment{}})
	return
}

// UpdateEnrollment ...
func UpdateEnrollment(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.Enrollment)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	a.ID = id
	err := store.Enrollments.Update(a, id)

	if err != nil {
		writeError(c.Writer, "error updating Enrollment", 500, err)
		return
	}

	c.JSON(200, &APIRes{"enrollment": []m.Enrollment{*a}})
	return
}

// GetAllEnrollments ...
func GetAllEnrollments(c *gin.Context) {
	filter := map[string]string{}
	if c.Request.URL.Query().Get("classId") != "" {
		filter["classId"] = c.Request.URL.Query().Get("classId")
	}
	if c.Request.URL.Query().Get("studentId") != "" {
		filter["studentId"] = c.Request.URL.Query().Get("studentId")
	}
	if c.Request.URL.Query().Get("termId") != "" {
		filter["termId"] = c.Request.URL.Query().Get("termId")
	}

	enrollments := []m.EnrollmentAPIRes{}
	err := store.Enrollments.Filter(&enrollments, filter)
	//query := store.EnrollmentH.OrderBy("firstName", "middleName", "lastName").Filter(filter)
	//err := store.DB.All(&enrollments, query)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"enrollment": enrollments})
	return
}
