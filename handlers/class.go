package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// GetAllClasses returns all classes, doesn't take in any params
func GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes := []*m.Class{}
	err := store.Classes.FindAll(&classes)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"class": classes})
	return
}

func GetClass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pID, _ := vars["id"]
	c := m.Class{}
	err := store.Classes.FindByID(&c, pID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"class": []m.Class{c}})
	return
}

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

	writeJSON(w, &APIRes{"class": c})
	return
}
