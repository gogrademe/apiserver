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
	address     string
	database string
)

func main() {
	flag.StringVar(&address, "address", "localhost:28015", "")
	flag.StringVar(&database, "database", "dev_go_grade", "")
	flag.Parse()

	if err := database.Init(driver, datasource); err != nil {
		log.Fatalln("Error setting up database: ", err)
	}

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(h.CORSMiddleware))
	n.UseHandler(setupHandlers())

	n.Run(port)

}

func setupHandlers() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
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

	return r
	// https://groups.google.com/forum/#!searchin/gorilla-web/options/gorilla-web/Xv4vMOlACyc/g5k7FoazMyoJ
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// 	if r.Method == "OPTIONS" {
	// 		fmt.Fprint(w)
	// 		return
	// 	}
	// 	m.ServeHTTP(w, r)
	//
	// })
}
