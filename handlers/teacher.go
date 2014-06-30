package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// CreateTeacher allows you to create a Teacher.
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	t := new(m.Teacher)

	errs := binding.Bind(r, t)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	id, err := store.Teachers.Store(t)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}
	t.ID = id

	writeJSON(w, &APIRes{"teacher": []m.Teacher{*t}})
	return
}

// GetTeacher will return a Teacher with all of their Profiles.
func GetTeacher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sID, _ := vars["id"]

	t := m.Teacher{}
	err := store.Teachers.FindByID(&t, sID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"teacher": []m.Teacher{t}})
	return
}

// UpdateTeacher allows you to create a Teacher.
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	t := new(m.Teacher)

	errs := binding.Bind(r, t)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	t.ID = pID
	err := store.Teachers.Update(t, pID)

	if err != nil {
		writeError(w, "Error updating Teacher", 500, err)
		return
	}

	writeJSON(w, &APIRes{"teacher": []m.Teacher{*t}})
	return
}

// GetAllTeachers returns all people without their profiles.
func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	teachers := []m.Teacher{}
	err := store.Classes.FindAll(&teachers)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"teacher": teachers})
	return
}
