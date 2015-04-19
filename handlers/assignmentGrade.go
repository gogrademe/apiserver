package handlers

import (
	"errors"

	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignmentGrade ...
func CreateAssignmentGrade(c *gin.Context) {
	a := new(m.AssignmentGrade)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	ids, err := store.AssignmentGrades.Insert(a)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}
	a.ID = ids[0]

	c.JSON(201, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAssignmentGrade ...
func GetAssignmentGrade(c *gin.Context) {

	id := c.Params.ByName("id")

	a := m.AssignmentGrade{}
	err := store.AssignmentGrades.One(&a, id)
	if err == store.ErrNotFound {
		writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}

	c.JSON(200, &APIRes{"grade": []m.AssignmentGrade{a}})
	return
}

// UpdateAssignmentGrade ...
func UpdateAssignmentGrade(c *gin.Context) {
	id := c.Params.ByName("id")

	a := new(m.AssignmentGrade)

	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	a.ID = id
	err := store.AssignmentGrades.Update(a, id)

	if err != nil {
		writeError(c.Writer, "Error updating AssignmentGrade", 500, err)
		return
	}

	c.JSON(200, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// r.db('dev_go_grade').table('grades').eqJoin('assignmentId', r.db('dev_go_grade').table('assignments')).filter({right: {classId: '3dd7e57d-f772-47cf-ade1-a7c6266e631a', termId: 'a80311b9-8c15-44e7-920f-05e5a32d3e8f'}})
//   .map(function(row) {
//     return {
//       assignment: row('right'),
//       assignmentGrade: {
//         gradeAverage: row('left')('grade').div(row('right')('maxScore'))
//       }
//     }
//   })

// GetAllAssignmentGrades ...
func GetAllAssignmentGrades(c *gin.Context) {
	filter := map[string]interface{}{}
	if c.Request.URL.Query().Get("assignmentId") != "" {
		filter["assignmentId"] = c.Request.URL.Query().Get("assignmentId")
	}

	if c.Request.URL.Query().Get("classId") != "" && c.Request.URL.Query().Get("termId") != "" {
		filter["right"] = map[string]string{
			"classId": c.Request.URL.Query().Get("classId"),
			"termId":  c.Request.URL.Query().Get("termId"),
		}
	}

	q := store.AssignmentGrades.EqJoin("assignmentId", r.Table("assignments"))
	q = q.Filter(filter)

	q = q.Map(func(row r.Term) r.Term {
		return row.Field("left").Merge(map[string]interface{}{
			"assignment":   row.Field("right"),
			"gradeAverage": row.Field("left").Field("grade").Div(row.Field("right").Field("maxScore")),
		})
	})

	grades := []m.AssignmentGradeAPIRes{}
	err := store.DB.All(&grades, q)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"grade": grades})
	return
}
