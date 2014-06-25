package database

import (
	// m "github.com/Lanciv/GoGradeAPI/model"
	// "log"
	"testing"
)

// TestUserDatabase verifies that a User can be saved and loaded from the database
func TestUserDatabase(t *testing.T) {
	// var userTests = []struct {
	// 	email      string
	// 	password   string
	// 	role       string
	// 	shouldFail bool
	// }{}

	_, err := CreateUser("test@test.com", "somePassword", "Admin")
	if err != nil {
		t.Errorf("Create valid user expected nil got %s", err.Error())
	}

	_, err = CreateUser("test@test.com", "somePassword", "Admin")
	t.Log(err)
	if err == nil {
		t.Error("Create invalid user expected Err got nil")
	}

	// if user.ID == "" {
	// 	t.Fatalf("exected user id to be defined")
	// }

}
