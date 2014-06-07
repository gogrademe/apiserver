package handlers

import (
	d "github.com/Lanciv/GoGradeAPI/database"
	m "github.com/Lanciv/GoGradeAPI/model"
	"fmt"
	"net/http"
)

func GetAllClasses(w http.ResponseWriter, r *http.Request) {

	classes, err := d.GetAllClasses()
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
