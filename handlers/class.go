package handlers

import (
	"github.com/Lanciv/GoGradeAPI/models"
	"github.com/martini-contrib/render"
	"net/http"
)

func GetAllClasses(r render.Render, req *http.Request) {

	classes, err := models.GetAllClasses()
	if err != nil {
		panic(err)
	}
	r.JSON(200, classes)
}
