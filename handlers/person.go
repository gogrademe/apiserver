package handlers

import (
	"database/sql"
	d "github.com/Lanciv/GoGradeAPI/database"
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreatePerson allows you to create a Person.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p m.PersonProfile

	if readJSON(r, &p) {
		// Person should exist before trying to do anything.
		if p.Person == nil {
			writeError(w, "Person required", 400)
			return
		}

		err := d.CreatePerson(p.Person)
		if err != nil {
			writeError(w, "Error creating Person", 500)
			return
		}
		// If there is a StudentProfile then create it.
		if p.StudentProfile != nil {
			err = d.CreateStudentProfile(p.Person.ID, p.StudentProfile)
			if err != nil {
				writeError(w, "Error creating Student Profile", 500)
				return
			}
		}

	} else {
		writeError(w, "Error parsing JSON", 400)
		return
	}

	writeJSON(w, p)
	return
}

// GetPerson will return a Person with all of their Profiles.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	var res m.PersonProfile

	vars := mux.Vars(r)

	pID, ok := vars["id"]
	if !ok {
		writeError(w, "Invalid Person ID", 400)
		return
	}

	id, err := strconv.Atoi(pID)
	if err != nil {
		writeError(w, "Invalid Person ID", 400)
		return
	}

	res.Person, err = d.GetPerson(id)
	if err == sql.ErrNoRows {
		writeError(w, "Person not found", 404)
		return
	}
	if err != nil {
		writeError(w, serverError, 400)
		return
	}

	res.StudentProfile, err = d.StudentProfileForPerson(id)
	if err != nil && err != sql.ErrNoRows {
		writeError(w, serverError, 500)
		return
	}

	writeJSON(w, res)
	return
}

// GetAllPeople returns all people without their profiles.
func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	people, err := d.GetAllPeople()
	if err != nil {
		writeError(w, serverError, 500)
		return
	}
	writeJSON(w, people)
	return
}
