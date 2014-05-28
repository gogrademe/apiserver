package handlers

import (
	m "bitbucket.org/lanciv/GoGradeAPI/models"
	"fmt"
	"net/http"
)

func GetAllClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := m.GetAllClasses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJson(w, classes)
}

// func handleTodoCreate(w http.ResponseWriter, r *http.Request) {
//   var (
//     todo Todo
//     err  error
//   )
//   data := struct {
//     Success bool `json:"success"`
//     Todo    Todo `json:"todo"`
//   }{
//     Success: false,
//   }
//   if readJson(r, &todo) {
//     if err = repo.Create(&todo); err != nil {
//       log.Printf("%v", err)
//     } else {
//       data.Success = true
//       data.Todo = todo
//     }
//   }

//   writeJson(w, data)
// }
func CreateClass(w http.ResponseWriter, r *http.Request) {
	var class m.Class

	if readJson(r, &class) {
		fmt.Println(class)
	}

}
