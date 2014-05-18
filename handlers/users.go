package handlers

import (
	"encoding/json"
	"github.com/Lanciv/GoGradeAPI/models"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users := models.GetAllUsers()

	enc := json.NewEncoder(w)
	err := enc.Encode(users)
	if err != nil {
		panic(err)
	}
}
