package handlers

import (
	"encoding/json"
	m "github.com/Lanciv/GoGradeAPI/models"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Get username and password then verify it.
	email, password := r.FormValue("email"), r.FormValue("password")
	user, err := m.VerifyPasswd(email, password)

	// Something happened with verifying the password. We either couldn't find the user,
	// password didn't match or something else went wrong.
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create a token for the user.
	token, err := user.CreateToken()
	if err != nil {
		http.Error(w, "Something bad happened!", http.StatusInternalServerError)
		return
	}

	// Send token to the user so they can use it to to authenticate all further requests.
	enc := json.NewEncoder(w)
	err = enc.Encode(map[string]interface{}{"token": token})
	if err != nil {
		http.Error(w, "Something bad happened!", http.StatusInternalServerError)
		return
	}
}

func AuthRequired(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := jwt.ParseFromRequest(r, func(t *jwt.Token) ([]byte, error) {
			return []byte("someRandomSigningKey"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}
