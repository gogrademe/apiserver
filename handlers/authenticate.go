package handlers

import (
	"errors"
	d "github.com/Lanciv/GoGradeAPI/database"
	m "github.com/Lanciv/GoGradeAPI/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mholt/binding"
	"net/http"
)

var ErrLoginFailed = errors.New("Login Failed! Email and/or password incorrect.")

type LoginForm struct {
	Email    string
	Password string
}

// Then provide a field mapping (pointer receiver is vital)
func (lf *LoginForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&lf.Email:    binding.Field{Form: "email", Required: true},
		&lf.Password: binding.Field{Form: "password", Required: true},
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Get username and password
	lf := new(LoginForm)

	errs := binding.Bind(r, lf)
	if errs != nil {
		writeError(w, errs, 400)
		return
	}

	user, err := d.GetUserEmail(lf.Email)
	if err != nil {
		writeError(w, ErrLoginFailed, http.StatusUnauthorized)
		return
	}

	if err := user.ComparePassword(lf.Password); err != nil {
		writeError(w, ErrLoginFailed, http.StatusUnauthorized)
		return
	}

	// Create a session for the user.
	session, err := m.NewSession(user)
	if err != nil {
		writeError(w, serverError, 500)
		return
	}

	d.SaveSession(&session)
	// Send token to the user so they can use it to to authenticate all further requests.
	writeJSON(w, &APIRes{"session": []m.Session{session}})
	return
}

func AuthRequired(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := jwt.ParseFromRequest(r, func(t *jwt.Token) ([]byte, error) {
			return []byte("someRandomSigningKey"), nil
		})
		if err != nil {
			writeError(w, "Access denied.", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}
