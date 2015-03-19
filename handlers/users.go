package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"

	m "github.com/gogrademe/apiserver/model"
	s "github.com/gogrademe/apiserver/store"

	"github.com/mholt/binding"
)

// GetAllUsers http endpoint to return all users.
func GetAllUsers(c *gin.Context) {
	u := []m.User{}
	err := s.Users.FindAll(&u)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(200, &APIRes{"user": u})
	return
}

type UserNew struct {
	PersonID string
	Email    string
}

// FieldMap ...
func (u *UserNew) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&u.PersonID: "personId",
		&u.Email:    "email",
	}
}

// CreateUser ...
func CreateUser(c *gin.Context) {
	u := new(UserNew)

	errs := binding.Bind(c.Request, u)
	if errs != nil {
		c.Error(errors.New("validation"), errs)
		c.JSON(StatusUnprocessable, errs)
		return
	}

	newUser, err := m.NewUserFor(u.Email, u.PersonID)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	err = s.Users.Store(newUser)
	if err != nil {
		writeError(c.Writer, serverError, 500, err)
		return
	}

	c.JSON(201, &APIRes{"user": []m.User{*newUser}})
	return
}

// func ActivateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
//
// 	emailConf := m.EmailConfirmation{}
// 	err := s.EmailConfirmations.FindByID(&emailConf, id)
// 	if err != nil {
// 		writeError(c.Writer, notFoundError, 404, err)
// 	}
//
// 	if emailConf.
//
// }
