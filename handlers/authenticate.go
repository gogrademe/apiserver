package handlers

import (
	"encoding/json"
	"fmt"
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
func VerifyToken(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.ParseFromRequest(r, func(t *jwt.Token) ([]byte, error) {
		return []byte("someRandomSigningKey"), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(token.Claims["Id"])
	userId := token.Claims["Id"].(int)
	user := &m.User{}
	m.GetUserById(userId)

	fmt.Println(user.Email)
	// c.Map(user)

}
