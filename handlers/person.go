package handlers

import (
	d "github.com/Lanciv/GoGradeAPI/repo"
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"

	"net/http"
)

// CreatePerson allows you to create a Person.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	p := new(m.Person)

	errs := binding.Bind(r, p)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}
	err := d.CreatePerson(p)

	if err != nil {
		writeError(w, "Error creating Person", 500, err)
		return
	}

	writeJSON(w, &APIRes{"person": p})
	return
}

// GetPerson will return a Person with all of their Profiles.
func GetPerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pID, _ := vars["id"]

	p, err := d.GetPerson(pID)
	if err != nil {
		writeError(w, serverError, 400, nil)
		return
	}
	if p == nil {
		writeError(w, notFoundError, 404, nil)
		return
	}

	writeJSON(w, &APIRes{"person": p})
	return
}

// GetAllPeople returns all people without their profiles.
func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	people, err := d.GetAllPeople()
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"person": people})
	return
}
