package handlers

import (
	"github.com/gin-gonic/gin"

	m "github.com/Lanciv/GoGradeAPI/model"
	s "github.com/Lanciv/GoGradeAPI/store"
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
