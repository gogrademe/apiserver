package handlers

import (
	"fmt"
	d "github.com/Lanciv/GoGradeAPI/database"
	m "github.com/Lanciv/GoGradeAPI/model"
	"net/http"
)

// GetAllClasses returns all classes, doesn't take in any params
func GetAllClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := d.GetAllClasses()
	if err != nil {
		writeError(w, serverError, 500)
		return
	}

	writeJSON(w, &APIRes{"class": classes})
	return
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	var class m.Class

	if readJSON(r, &class) {
		fmt.Println(class)
	}

}
