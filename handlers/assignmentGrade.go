package handlers

import (
	"errors"

	"github.com/Sirupsen/logrus"
	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"

	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"
)

// CreateAssignmentGrade ...
func CreateAssignmentGrade(c *gin.Context) {
	a := new(m.AttemptResource)
	errs := binding.Bind(c.Request, a)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	logrus.Info(a)

	attemptsUpdate := m.AssignmentAttempts{
		AssignmentID:  a.AssignmentID,
		PersonID:      a.PersonID,
		LatestAttempt: m.Attempt{Score: a.Score},
		AttemptHistory: []*m.Attempt{
			&m.Attempt{Score: a.Score},
		},
	}

	res, err := store.DB.RunWrite(store.Attempts.Term.Insert(attemptsUpdate, r.InsertOpts{Conflict: "update"}))
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	logrus.Info(res)
	// a.ID = ids[0]

	// c.JSON(201, &APIRes{"grade": []m.GradebookResource{*a}})
	return
}

// GetAssignmentGrade ...
// func GetAssignmentGrade(c *gin.Context) {
//
// 	id := c.Params.ByName("id")
//
// 	a := m.AssignmentGrade{}
// 	err := store.AssignmentGrades.One(&a, id)
// 	if err == store.ErrNotFound {
// 		writeError(c.Writer, notFoundError, 404, nil)
// 		return
// 	}
// 	if err != nil {
// 		writeError(c.Writer, serverError, 500, nil)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"grade": []m.AssignmentGrade{a}})
// 	return
// }

// UpdateAssignmentGrade ...
// func UpdateAssignmentGrade(c *gin.Context) {
// 	id := c.Params.ByName("id")
//
// 	a := new(m.AssignmentGrade)
//
// 	errs := binding.Bind(c.Request, a)
// 	if errs != nil {
// 		c.Error(errors.New("validation"), errs)
// 		c.JSON(StatusUnprocessable, errs)
// 		return
// 	}
//
// 	a.ID = id
// 	err := store.AssignmentGrades.Update(a, id)
//
// 	if err != nil {
// 		writeError(c.Writer, "Error updating AssignmentGrade", 500, err)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"grade": []m.AssignmentGrade{*a}})
// 	return
// }

// GetAllAssignmentGrades ...
func GetAllAssignmentGrades(c *gin.Context) {
	filter := map[string]interface{}{}
	if c.Request.URL.Query().Get("assignmentId") != "" {
		filter["assignmentId"] = c.Request.URL.Query().Get("assignmentId")
	}

	if c.Request.URL.Query().Get("classId") != "" && c.Request.URL.Query().Get("termId") != "" {
		filter["classId"] = c.Request.URL.Query().Get("classId")
		filter["termId"] = c.Request.URL.Query().Get("termId")
	}

	q := store.EnrollmentH.Filter(filter)
	q = q.Merge(func(row r.Term) map[string]interface{} {
		return map[string]interface{}{
			"assignmentAttempts": store.Attempts.GetAllByIndex("personId", row.Field("personId")).CoerceTo("array"),
		}
	})

	grades := []m.GradebookResource{}
	err := store.DB.All(&grades, q)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"grade": grades})
	return
}
