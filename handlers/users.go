package handlers

import (
	"encoding/json"
	"github.com/Lanciv/GoGradeAPI/models"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
