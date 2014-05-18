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
			panic(err)
		}
	}

	writeJson(w, person)
}
