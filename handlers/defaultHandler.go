package handlers

//
// import (
// 	"errors"
// 	"fmt"
// 	m "github.com/Lanciv/GoGradeAPI/model"
// 	"github.com/Lanciv/GoGradeAPI/store"
//
// 	"github.com/gin-gonic/gin"
// 	"github.com/mholt/binding"
// )
//
// // Create ...
// func Create(s store.Storer) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		a := new(m.Assignment)
//
// 		errs := binding.Bind(c.Request, a)
// 		if errs != nil {
// 			c.Error(errors.New("validation"), errs)
// 			c.JSON(400, errs)
// 			return
// 		}
//
// 		id, err := store.Assignments.Store(a)
// 		if err != nil {
// 			writeError(c.Writer, serverError, 500, err)
// 			return
// 		}
// 		a.ID = id
//
// 		c.JSON(201, &APIRes{"assignment": []m.Assignment{*a}})
// 		return
// 	}
// }
//
// // Get ...
// func Get(c *gin.Context) {
//
// 	id := c.Params.ByName("id")
//
// 	a := m.Assignment{}
// 	err := store.Assignments.FindByID(&a, id)
// 	if err == store.ErrNotFound {
// 		writeError(c.Writer, notFoundError, 404, nil)
// 		return
// 	}
// 	if err != nil {
// 		writeError(c.Writer, serverError, 500, nil)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"assignment": []m.Assignment{a}})
// 	return
// }
//
// // Update ...
// func Update(c *gin.Context) {
// 	id := c.Params.ByName("id")
//
// 	a := new(m.Assignment)
//
// 	errs := binding.Bind(c.Request, a)
// 	if errs != nil {
// 		writeError(c.Writer, errs, 400, nil)
// 		return
// 	}
//
// 	a.ID = id
// 	err := store.Assignments.Update(a, id)
//
// 	if err != nil {
// 		writeError(c.Writer, "error updating Assignment", 500, err)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"assignment": []m.Assignment{*a}})
// 	return
// }
//
// // GetAll ...
// func GetAll(c *gin.Context) {
// 	a := new(m.Assignment)
// 	_ = binding.Bind(c.Request, a)
// 	fmt.Println(a)
//
// 	assignment := []m.Assignment{}
// 	// err := store.Assignments.FindAll(&assignment)
// 	err := store.Assignments.Filter(&assignment, map[string]string{"classId": a.ClassID, "termId": a.TermID})
// 	if err != nil {
// 		writeError(c.Writer, serverError, 500, err)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"assignment": assignment})
// 	return
// }
