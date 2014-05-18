package main

import (
	"fmt"
	h "github.com/Lanciv/GoGradeAPI/handlers"
	// "github.com/Lanciv/GoGradeAPI/models"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Server is starting...")

	// models.SetupDB()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	/* Auth */
	s.HandleFunc("/auth/login", h.Login)

	/* Users */
	s.HandleFunc("/users", h.AuthRequired(h.GetAllUsers))
	http.Handle("/", r)
	// TODO: Add port to Config.

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, nil)
}
