package handlers

import (
	"errors"
	d "github.com/Lanciv/GoGradeAPI/database"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var ErrLoginFailed = errors.New("Login Failed! Email and/or password incorrect.")

func Login(w http.ResponseWriter, r *http.Request) {
	// Get username and password
	email, password := r.FormValue("email"), r.FormValue("password")
	user, err := d.GetUserEmail(email)

	if err != nil {
		http.Error(w, ErrLoginFailed.Error(), http.StatusUnauthorized)
		return
	}

	if err := user.ComparePassword(password); err != nil {
		http.Error(w, ErrLoginFailed.Error(), http.StatusUnauthorized)
		return
	}

	// Create a token for the user.
	token, err := user.CreateToken()
	if err != nil {
		http.Error(w, "Something bad happened!", http.StatusInternalServerError)
		return
	}

	// Send token to the user so they can use it to to authenticate all further requests.
	writeJSON(w, &APIRes{"session": map[string]interface{}{"token": token}})
	return
}

func AuthRequired(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := jwt.ParseFromRequest(r, func(t *jwt.Token) ([]byte, error) {
			return []byte("someRandomSigningKey"), nil
		})
		if err != nil {
			log.Println("error", err)
			// w.WriteHeader(http.StatusUnauthorized)
			// writeJSON(w, map[string]interface{}{"error": "Access denied."})
			writeError(w, "Access denied.", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}
