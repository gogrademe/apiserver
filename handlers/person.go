package handlers

import (
	"database/sql"
	d "github.com/Lanciv/GoGradeAPI/database"
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// PersonResponse will be used when returning People or a Person.
// type PersonAPI struct {
// 	Person         m.Person         `json:"person"`
// 	StudentProfile m.StudentProfile `json:"studentProfile,omitempty"`
// }

// CreatePerson allows you to create a Person.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p m.PersonProfile
	if readJson(r, &p) {
		// Create Person.

		err := d.CreatePerson(p.Person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if p.StudentProfile != nil {
			err = d.CreateStudentProfile(p.Person.Id, p.StudentProfile)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	} else {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	writeJson(w, p)
}

// GetPerson will return a Person with all of their Profiles.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	var res m.PersonProfile

	vars := mux.Vars(r)

	pID, ok := vars["id"]

	if !ok {
		http.Error(w, "error", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(pID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Person, err = d.GetPerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res.StudentProfile, err = d.StudentProfileForPerson(id)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJson(w, res)
}

// GetAllPeople returns all people without their profiles.
func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	people, err := d.GetAllPeople()
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, people)
}
