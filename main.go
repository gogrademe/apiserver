package main

import (
	"flag"
	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/Lanciv/GoGradeAPI/repo"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"log"
)

var (
	apiPort  string
	address  string
	dbName   string
	testData bool
)

func main() {
	flag.StringVar(&apiPort, "apiPort", ":5000", "")
	flag.StringVar(&address, "dbAddress", "localhost:28015", "")
	flag.StringVar(&dbName, "dbName", "dev_go_grade", "")
	flag.BoolVar(&testData, "testData", true, "")
	flag.Parse()

	if err := repo.Connect(address, dbName); err != nil {
		log.Fatalln("Error setting up database: ", err)
	}

	repo.SetupDB(testData)
	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
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
	m.HandleFunc("/session", h.Login).Methods("POST")

	// Users
	m.HandleFunc("/user", h.AuthRequired(h.GetAllUsers)).Methods("GET")

	// Classes
	m.HandleFunc("/class", h.AuthRequired(h.GetAllClasses)).Methods("GET")
	m.HandleFunc("/class", h.AuthRequired(h.CreateClass)).Methods("POST")

	// People
	m.HandleFunc("/person", h.AuthRequired(h.GetAllPeople)).Methods("GET")
	m.HandleFunc("/person", h.AuthRequired(h.CreatePerson)).Methods("POST")
	m.HandleFunc("/person/{id}", h.AuthRequired(h.GetPerson)).Methods("GET")

	// Students
	m.HandleFunc("/student", h.AuthRequired(h.CreateStudent)).Methods("POST")
	return r
}
