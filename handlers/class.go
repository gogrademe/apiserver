package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"
	"github.com/mholt/binding"

	"net/http"
)

// GetAllClasses returns all classes, doesn't take in any params
func GetAllClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := store.Classes.FindAll()
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"class": classes})
	return
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	c := new(m.Class)

	errs := binding.Bind(r, c)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	err := store.Classes.Store(c)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"class": c})
	return
}
