package handlers

import (
	"net/http"

	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"

	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

// GetAllClassTerms returns all terms, doesn't take in any params
func GetAllClassTerms(w http.ResponseWriter, r *http.Request) {
	terms := []m.ClassTerm{}
	err := store.ClassTerms.FindAll(&terms)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"classTerm": terms})
	return
}

func GetClassTerm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pID, _ := vars["id"]
	c := m.ClassTerm{}
	err := store.ClassTerms.FindByID(&c, pID)
	if err == store.ErrNotFound {
		writeError(w, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(w, serverError, 500, nil)
		return
	}

	writeJSON(w, &APIRes{"classTerm": []m.ClassTerm{c}})
	return
}

func CreateClassTerm(w http.ResponseWriter, r *http.Request) {
	c := new(m.ClassTerm)

	errs := binding.Bind(r, c)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	id, err := store.ClassTerms.Store(c)

	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}
	c.ID = id

	writeJSON(w, &APIRes{"classTerm": []m.ClassTerm{*c}})
	return
}

func UpdateClassTerm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pID, _ := vars["id"]

	c := new(m.ClassTerm)

	errs := binding.Bind(r, c)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	c.ID = pID
	err := store.ClassTerms.Update(c, pID)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"classTerm": []m.ClassTerm{*c}})
	return
}
