package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	s "github.com/Lanciv/GoGradeAPI/store"
)

// GetAllUsers http endpoint to return all users.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	u := []m.User{}
	err := s.Users.FindAll(&u)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"user": u})
	return
}
