package handlers

import (
	"fmt"
	m "github.com/Lanciv/GoGradeAPI/models"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person m.Person
	if readJson(r, &person) {
		fmt.Println(person)
		_, err := m.CreatePerson(&person)

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

	people, err := m.GetAllPeople()
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, people)
}
