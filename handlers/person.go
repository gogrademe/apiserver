package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// CreatePerson ...
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	p := new(m.Person)

	errs := binding.Bind(r, p)
	if errs != nil {
		writeError(w, errs, 500, nil)
		return
	}

	id, err := store.People.Store(p)
	if err != nil {
		writeError(w, "Error creating Person", 500, err)
		return
	}

	p.ID = id

	writeJSON(w, &APIRes{"person": []m.Person{*p}})
	return
}

// UpdatePerson ...
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	p := new(m.Person)

	errs := binding.Bind(r, p)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	p.ID = pID
	err := store.People.Update(p, pID)

	if err != nil {
		writeError(w, "Error updating Person", 500, err)
		return
	}

	writeJSON(w, &APIRes{"person": []m.Person{*p}})
	return
}

// GetPerson ...
func GetPerson(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pID, _ := vars["id"]

	p := m.Person{}
	err := store.People.FindByID(&p, pID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"person": []m.Person{p}})
	return
}

// GetAllPeople ...
func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	people, err := store.People.FindAll()
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"person": people})
	return
}
