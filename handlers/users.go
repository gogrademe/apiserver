package handlers

import (
	s "github.com/Lanciv/GoGradeAPI/store"
	// "github.com/Lanciv/GoGradeAPI/model"

	"net/http"
)

// GetAllUsers http endpoint to return all users.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := s.Users.FindAll()
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"user": users})
	return
}
