package handlers

import (
	"bitbucket.org/lanciv/GoGradeAPI/models"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, users)

}
