package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// CreateAssignment ...
func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	a := new(m.Assignment)

	errs := binding.Bind(r, a)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	id, err := store.Assignments.Store(a)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}
	a.ID = id

	writeJSON(w, &APIRes{"assignment": []m.Assignment{*a}})
	return
}

// GetAssignment ...
func GetAssignment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sID, _ := vars["id"]

	a := m.Assignment{}
	err := store.Assignments.FindByID(&a, sID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"assignment": []m.Assignment{a}})
	return
}

// UpdateAssignment ...
func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	a := new(m.Assignment)

	errs := binding.Bind(r, a)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	a.ID = pID
	err := store.Assignments.Update(a, pID)

	if err != nil {
		writeError(w, "Error updating Assignment", 500, err)
		return
	}

	writeJSON(w, &APIRes{"assignment": []m.Assignment{*a}})
	return
}

// GetAllAssignments ...
func GetAllAssignments(w http.ResponseWriter, r *http.Request) {
	assignment := []m.Assignment{}
	err := store.Classes.FindAll(&assignment)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"assignment": assignment})
	return
}
