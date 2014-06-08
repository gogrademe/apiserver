package main

import (
	"flag"
	"fmt"
	"github.com/Lanciv/GoGradeAPI/database"
	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port       string
	driver     string
	datasource string
)

func main() {

	flag.StringVar(&port, "port", ":3000", "")
	flag.StringVar(&driver, "driver", "postgres", "")
	flag.StringVar(&datasource, "datasource", "user=Matt dbname=dev_go_grade sslmode=disable", "")
	flag.Parse()

	if err := database.Init(driver, datasource); err != nil {
		log.Fatalf("Error setting up database: ", err)
	}

	setupHandlers()

	panic(http.ListenAndServe(port, nil))

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
	m.HandleFunc("/people", h.AuthRequired(h.GetAllPeople)).Methods("GET")
	m.HandleFunc("/people/{id}", h.AuthRequired(h.GetPerson)).Methods("GET")
	m.HandleFunc("/people/create", h.AuthRequired(h.CreatePerson)).Methods("POST")

	// https://groups.google.com/forum/#!searchin/gorilla-web/options/gorilla-web/Xv4vMOlACyc/g5k7FoazMyoJ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			fmt.Fprint(w)
			return
		}
		m.ServeHTTP(w, r)
	})
}
