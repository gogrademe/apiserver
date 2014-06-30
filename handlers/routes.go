package handlers

import (
	"github.com/gorilla/mux"
)

// SetupHandlers loads all routes into gorillaMux.
func SetupHandlers() *mux.Router {
	r := mux.NewRouter()
	// r.StrictSlash(true)
	m := r.PathPrefix("/api").Subrouter()

	// Auth
	m.HandleFunc("/session", Login).Methods("POST")

	// Users
	m.HandleFunc("/user", AuthRequired(GetAllUsers)).Methods("GET")

	// Classes
	m.HandleFunc("/class", AuthRequired(GetAllClasses)).Methods("GET")
	m.HandleFunc("/class", AuthRequired(CreateClass)).Methods("POST")
	m.HandleFunc("/class/{id}", AuthRequired(GetClass)).Methods("GET")
	m.HandleFunc("/class/{id}", AuthRequired(UpdateClass)).Methods("PUT")

	// ClassTerms
	m.HandleFunc("term", AuthRequired(GetAllClassTerms)).Methods("GET")
	m.HandleFunc("term", AuthRequired(CreateClassTerm)).Methods("POST")
	m.HandleFunc("term/{id}", AuthRequired(GetClassTerm)).Methods("GET")
	m.HandleFunc("term/{id}", AuthRequired(UpdateClassTerm)).Methods("PUT")

	// Assignments
	m.HandleFunc("assignment", AuthRequired(GetAllAssignments)).Methods("GET")
	m.HandleFunc("assignment", AuthRequired(CreateAssignment)).Methods("POST")
	m.HandleFunc("assignment/{id}", AuthRequired(GetAssignment)).Methods("GET")
	m.HandleFunc("assignment/{id}", AuthRequired(UpdateAssignment)).Methods("PUT")

	// People
	m.HandleFunc("/person", AuthRequired(GetAllPeople)).Methods("GET")
	m.HandleFunc("/person", AuthRequired(CreatePerson)).Methods("POST")
	m.HandleFunc("/person/{id}", AuthRequired(GetPerson)).Methods("GET")
	m.HandleFunc("/person/{id}", AuthRequired(UpdatePerson)).Methods("PUT")

	// Students
	m.HandleFunc("/student", AuthRequired(GetAllStudents)).Methods("GET")
	m.HandleFunc("/student", AuthRequired(CreateStudent)).Methods("POST")
	m.HandleFunc("/student/{id}", AuthRequired(GetStudent)).Methods("GET")
	m.HandleFunc("/student/{id}", AuthRequired(UpdateStudent)).Methods("PUT")

	// Teachers
	m.HandleFunc("/teacher", AuthRequired(GetAllTeachers)).Methods("GET")
	m.HandleFunc("/teacher", AuthRequired(CreateTeacher)).Methods("POST")
	m.HandleFunc("/teacher/{id}", AuthRequired(GetTeacher)).Methods("GET")
	m.HandleFunc("/teacher/{id}", AuthRequired(UpdateTeacher)).Methods("PUT")
	return r
}
