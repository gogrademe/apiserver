package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// CreateAssignmentGrade ...
func CreateAssignmentGrade(w http.ResponseWriter, r *http.Request) {
	a := new(m.AssignmentGrade)

	errs := binding.Bind(r, a)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	id, err := store.AssignmentGrades.Store(a)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}
	a.ID = id

	writeJSON(w, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAssignmentGrade ...
func GetAssignmentGrade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sID, _ := vars["id"]

	a := m.AssignmentGrade{}
	err := store.AssignmentGrades.FindByID(&a, sID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"grade": []m.AssignmentGrade{a}})
	return
}

// UpdateAssignmentGrade ...
func UpdateAssignmentGrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	a := new(m.AssignmentGrade)

	errs := binding.Bind(r, a)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	a.ID = pID
	err := store.AssignmentGrades.Update(a, pID)

	if err != nil {
		writeError(w, "Error updating AssignmentGrade", 500, err)
		return
	}

	writeJSON(w, &APIRes{"grade": []m.AssignmentGrade{*a}})
	return
}

// GetAllAssignmentGrades ...
func GetAllAssignmentGrades(w http.ResponseWriter, r *http.Request) {
	grades := []m.AssignmentGrade{}
	err := store.Classes.FindAll(&grades)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"grade": grades})
	return
}
