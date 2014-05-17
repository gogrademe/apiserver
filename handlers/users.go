package handlers

import (
	"github.com/Lanciv/GoGradeAPI/models"
	"github.com/martini-contrib/render"
	"net/http"
)

func GetAllUsers(r render.Render, req *http.Request) {

	users := models.GetAllUsers()

	r.JSON(200, users)
}
