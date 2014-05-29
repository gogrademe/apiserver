package handlers

import (
	d "bitbucket.org/lanciv/GoGradeAPI/database"
	m "bitbucket.org/lanciv/GoGradeAPI/model"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person m.Person
	if readJson(r, &person) {
		_, err := d.CreatePerson(&person)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	writeJson(w, person)
}
func GetAllPeople(w http.ResponseWriter, r *http.Request) {

	people, err := d.GetAllPeople()
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, people)
}
