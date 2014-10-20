package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"

	m "github.com/Lanciv/GoGradeAPI/model"
	s "github.com/Lanciv/GoGradeAPI/store"

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
	Password string
}

// FieldMap ...
func (u *UserNew) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&u.PersonID: "personId",
		&u.Email:    "email",
		&u.Password: "password",
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

	newUser, err := m.NewUserFor(u.Email, u.Password, u.PersonID)
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

	// c.JSON(201, &APIRes{"user": []m.Assignment{*a}})
	return
}
