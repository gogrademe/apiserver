package handlers

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	"github.com/Lanciv/GoGradeAPI/store"
	"github.com/mholt/binding"
	"net/http"
)

// CreateStudent allows you to create a Student.
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	s := new(m.Student)

	errs := binding.Bind(r, s)
	if errs != nil {
		writeError(w, errs, 400, nil)
		return
	}

	err := store.Students.Store(s)
	if err != nil {
		writeError(w, serverError, 500, err)
		return
	}

	writeJSON(w, &APIRes{"student": s})
	return
}

// GetStudent will return a Student with all of their Profiles.
// func GetStudent(w http.ResponseWriter, r *http.Request) {
//
// 	vars := mux.Vars(r)
// 	pID, _ := vars["id"]
// 	// if !ok {
// 	// 	writeError(w, "Invalid Student ID", 400)
// 	// 	return
// 	// }
// 	log.Println(pID)
//
// 	s, err := d.GetStudent(pID)
// 	if err != nil {
// 		writeError(w, serverError, 400)
// 		return
// 	}
//
// 	if s == nil {
// 		writeError(w, notFoundError, 404)
// 		return
// 	}
//
// 	writeJSON(w, &APIRes{"person": s})
// 	return
// }

// GetAllPeople returns all students without their profiles.
// func GetAllPeople(w http.ResponseWriter, r *http.Request) {
//
// 	students, err := d.GetAllStudents()
// 	if err != nil {
// 		writeError(w, serverError, 500)
// 		return
// 	}
//
// 	writeJSON(w, &APIRes{"student": students})
// 	return
// }
