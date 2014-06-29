package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// CreateStudent allows you to create a Student.
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	s := new(m.Student)

	errs := binding.Bind(r, s)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	err := store.Students.Store(s)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"student": s})
	return
}

// GetStudent will return a Student with all of their Profiles.
func GetStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sID, _ := vars["id"]

	p, err := store.Students.FindByID(sID)
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}
	if p == nil {
		writeError(w, notFoundError, 404, nil)
		return
	}

	writeJSON(w, &APIRes{"student": p})
	return
}
