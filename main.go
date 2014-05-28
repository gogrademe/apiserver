package main

import (
	// "flag"
	"bitbucket.org/lanciv/GoGradeAPI/config"
	h "bitbucket.org/lanciv/GoGradeAPI/handlers"
	"bitbucket.org/lanciv/GoGradeAPI/models"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"os"
)

var (
	AppConfig *config.Conf
)

func main() {

	log.Println("Server is starting...")

	// configfile := flag.String("c", "config.gcfg", "Configuration file")

	// flag.Parse()

	// var err error
	// AppConfig, err = config.Load(*configfile)

	// if err != nil {
	// 	log.Fatalf("Config Error:", err.Error())
	// }

	models.SetupDB()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	/* Auth */
	s.HandleFunc("/auth/login", h.Login).Methods("POST")

	/* Users */
	s.HandleFunc("/users", h.AuthRequired(h.GetAllUsers)).Methods("GET")

	/* Class */
	s.HandleFunc("/class", h.AuthRequired(h.GetAllClasses)).Methods("GET")
	s.HandleFunc("/class/create", h.AuthRequired(h.CreateClass)).Methods("POST")

	/* Person */
	s.HandleFunc("/person", h.AuthRequired(h.GetAllPeople)).Methods("GET")
	s.HandleFunc("/person/create", h.AuthRequired(h.CreatePerson)).Methods("POST")

	n := negroni.Classic()

	n.UseHandler(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	n.Run(":" + port)

}
