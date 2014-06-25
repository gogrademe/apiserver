package main

import (
	"flag"
	"github.com/Lanciv/GoGradeAPI/database"
	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
)

var (
	apiPort string
	address string
	dbName  string
)

func main() {
	flag.StringVar(&apiPort, "apiPort", ":5000", "")
	flag.StringVar(&address, "address", "localhost:28015", "")
	flag.StringVar(&dbName, "dbName", "dev_go_grade", "")
	flag.Parse()

	if err := database.Connect(address, dbName); err != nil {
		log.Fatalln("Error setting up database: ", err)
	}

	database.SetupDB(true)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(h.CORSMiddleware))
	n.UseHandler(setupHandlers())

	n.Run(apiPort)

}

// setupHandlers loads all routes into gorillaMux.
func setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	m := r.PathPrefix("/api").Subrouter()

	// Auth
	m.HandleFunc("/auth/login", h.Login).Methods("POST")
	m.HandleFunc("/session", h.Login).Methods("POST")

	// Users
	m.HandleFunc("/user", h.AuthRequired(h.GetAllUsers)).Methods("GET")

	// Classes
	m.HandleFunc("/class", h.AuthRequired(h.GetAllClasses)).Methods("GET")
	m.HandleFunc("/class/create", h.AuthRequired(h.CreateClass)).Methods("POST")

	// People
	m.HandleFunc("/people", h.AuthRequired(h.GetAllPeople)).Methods("GET")
	m.HandleFunc("/people/{id}", h.AuthRequired(h.GetPerson)).Methods("GET")
	m.HandleFunc("/people/create", h.AuthRequired(h.CreatePerson)).Methods("POST")

	return r
}
