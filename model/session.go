package model

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type Session struct {
	ID        string    `gorethink:"id,omitempty"json:"id"`
	Token     string    `gorethink:"token"json:"token"`
	UserID    string    `gorethink:"userId"json:"userId"`
	ExpiresAt time.Time `gorethink:"expiresAt"json:"expiresAt"`
	TimeStamp
}

// Create a token for the user after we verified their password.
func NewSession(u User) (Session, error) {
	var s Session
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	tokenExpires := time.Now().UTC().Add(time.Hour * 72)
	token.Claims["id"] = u.ID
	token.Claims["email"] = u.Email
	token.Claims["exp"] = tokenExpires.Unix()

	// TODO: Move this to a config file.
	tokenString, err := token.SignedString([]byte("someRandomSigningKey"))
	if err != nil {
		return s, err
	}

	s = Session{
		Token:     tokenString,
		UserID:    u.ID,
		ExpiresAt: tokenExpires,
	}

	return s, err
}
