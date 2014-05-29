package main

import (
	// "flag"
	"bitbucket.org/lanciv/GoGradeAPI/config"
	"bitbucket.org/lanciv/GoGradeAPI/database"
	h "bitbucket.org/lanciv/GoGradeAPI/handlers"
	// "github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	database.SetupDB()

	setupHandlers()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3007"
	}

	panic(http.ListenAndServe(":"+port, nil))

}

func setupHandlers() {
	r := mux.NewRouter()
	m := r.PathPrefix("/api").Subrouter()

	// Auth
	m.HandleFunc("/auth/login", h.Login).Methods("POST")

	// Users
	m.HandleFunc("/user", h.AuthRequired(h.GetAllUsers)).Methods("GET")

	// Classes
	m.HandleFunc("/class", h.AuthRequired(h.GetAllClasses)).Methods("GET")
	m.HandleFunc("/class/create", h.AuthRequired(h.CreateClass)).Methods("POST")

	// People
	m.HandleFunc("/person", h.AuthRequired(h.GetAllPeople)).Methods("GET")
	m.HandleFunc("/person/create", h.AuthRequired(h.CreatePerson)).Methods("POST")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.ServeHTTP(w, r)
	})
}
