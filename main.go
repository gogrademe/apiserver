package main

import (
	"fmt"
	h "github.com/Lanciv/GoGradeAPI/handlers"
	"github.com/Lanciv/GoGradeAPI/models"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	fmt.Println("Server is starting...")

	models.SetupDB()

	r := mux.NewRouter()
	r.HandleFunc("/auth", h.Login)

	http.Handle("/api", r)
	// TODO: Add port to Config.
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server is running...")
}
