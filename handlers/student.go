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

	id, err := store.Students.Store(s)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}
	s.ID = id

	writeJSON(w, &APIRes{"student": []m.Student{*s}})
	return
}

// GetStudent will return a Student with all of their Profiles.
func GetStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sID, _ := vars["id"]

	s := m.Student{}
	err := store.Students.FindByID(&s, sID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"student": []m.Student{s}})
	return
}

// UpdateStudent allows you to create a Student.
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	s := new(m.Student)

	errs := binding.Bind(r, s)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	s.ID = pID
	err := store.Students.Update(s, pID)

	if err != nil {
		writeError(w, "Error updating Student", 500, err)
		return
	}

	writeJSON(w, &APIRes{"student": []m.Student{*s}})
	return
}

// GetAllStudents returns all people without their profiles.
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := []m.Student{}
	err := store.Classes.FindAll(&students)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"student": students})
	return
}
