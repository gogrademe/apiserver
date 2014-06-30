package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// GetAllClasses ...
func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes := []m.Class{}
	err := store.Classes.FindAll(&classes)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"class": classes})
	return
}

// GetClass ...
func GetClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pID, _ := vars["id"]
	c := m.Class{}
	err := store.Classes.FindByID(&c, pID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"class": []m.Class{c}})
	return
}

//CreateClass ...
func CreateClass(w http.ResponseWriter, r *http.Request) {
	c := new(m.Class)

	errs := binding.Bind(r, c)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	id, err := store.Classes.Store(c)

	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}
	c.ID = id

	writeJSON(w, &APIRes{"class": []m.Class{*c}})
	return
}

//UpdateClass ...
func UpdateClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	c := new(m.Class)

	errs := binding.Bind(r, c)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	c.ID = pID
	err := store.Classes.Update(c, pID)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"class": []m.Class{*c}})
	return
}
