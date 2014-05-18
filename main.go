package main

import (
	"fmt"
	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/Lanciv/GoGradeAPI/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Server is starting...")

	models.SetupDB()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	/* Auth */
	s.HandleFunc("/auth/login", h.Login)

	/* Users */
	s.HandleFunc("/users", h.AuthRequired(h.GetAllUsers))

	/* Class */
	s.HandleFunc("/class/create", h.AuthRequired(h.CreateClass))

	/* person */
	s.HandleFunc("/person/create", h.AuthRequired(h.CreatePerson))
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
