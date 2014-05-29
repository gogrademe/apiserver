package handlers

import (
	d "bitbucket.org/lanciv/GoGradeAPI/database"
	// "bitbucket.org/lanciv/GoGradeAPI/model"

	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := d.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, users)

}
