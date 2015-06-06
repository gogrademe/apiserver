package handlers

import (
	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// GetAllCourses ...
func GetAllCourses(c *gin.Context) {
	courses := []m.Course{}

	query := store.Courses.OrderBy(r.Asc("gradeLevel"), r.Asc("name"))
	err := store.DB.All(&courses, query)
	// err := store.Coursees.FindAll(&courses)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"course": courses})
	return
}

// GetCourse ...
func GetCourse(c *gin.Context) {
	var (
		id     = c.Param("id")
		course = m.Course{}
	)

	err := store.Courses.One(&course, id)
	if err != nil {
		handleDBError(c.Writer, err)
		return
	}
	// if err == store.ErrNotFound {
	// 	writeError(c.Writer, notFoundError, 404, nil)
	// 	return
	// }
	// if err != nil {
	// 	writeError(c.Writer, serverError, 500, nil)
	// 	return
	// }

	c.JSON(200, &APIRes{"course": []m.Course{course}})
	return
}

//CreateCourse ...
func CreateCourse(c *gin.Context) {
	course := new(m.Course)

	errs := binding.Bind(c.Request, course)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	ids, err := store.Courses.Insert(course)

	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	course.ID = ids[0]

	c.JSON(201, &APIRes{"course": []m.Course{*course}})
	return
}

//UpdateCourse ...
func UpdateCourse(c *gin.Context) {

	id := c.Param("id")

	course := new(m.Course)

	errs := binding.Bind(c.Request, course)
	if errs != nil {
		writeError(c.Writer, errs, 400, nil)
		return
	}

	course.ID = id
	err := store.Courses.Update(course, id)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"course": []m.Course{*course}})
	return
}

// DeleteCourse ...
func DeleteCourse(c *gin.Context) {

	id := c.Param("id")

	_, err := store.DB.RunWrite(store.Courses.Get(id).Delete())
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"course": []m.Course{}})
	return
}
