package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	// "log"
	"testing"
)

// TestUserDatabase verifies that a User can be saved and loaded from the database
func TestUserDatabase(t *testing.T) {
	var userTests = []struct {
		email      string
		password   string
		role       string
		shouldFail bool
		checkErr   error
	}{
		{
			email:    "test@test.com",
			password: "somePassword",
			role:     "Admin",
		},
		{
			password: "somePassword",
			role:     "Admin",
		},
		{
			role: "Admin",
		},
		{
			email:    "test@test.com",
			password: "somePassword",
			role:     "Admin",
			checkErr: ErrUserAlreadyExists,
		},
	}

	for i, user := range userTests {
		u, err := m.NewUser(user.email, user.password, user.role)
		err = Users.Store(u)
		if user.checkErr == nil {
			if err != nil {
				t.Errorf("Expected: nil error for user %v. Got: %s", i, err)
			}
		} else {
			if err != user.checkErr {
				t.Errorf("Expected: %s for user %v. Got: %s", user.checkErr, i, err)
			}
		}
	}
}
