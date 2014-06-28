package model

import (
	"errors"
	"strings"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("Username or password incorrect.")
)

type User struct {
	ID             string `gorethink:"id,omitempty"json:"id"`
	Email          string `gorethink:"email"json:"email"`
	EmailLower     string `gorethink:"emailLower"json:"emailLower"`
	HashedPassword string `gorethink:"hashedPassword"json:"-"`
	Role           string `gorethink:"role"json:"role"`
	Disabled       bool   `gorethink:"disabled"json:"disabled"`
	TimeStamp
}

func NewUser(email, password, role string) (*User, error) {

	emailLower := strings.ToLower(email)

	user := User{
		Email:      email,
		EmailLower: emailLower,
		Role:       role,
	}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	user.UpdateTime()
	return &user, nil
}

func (u *User) SetPassword(password string) error {
	// Password validation.
	switch {
	case len(password) < 6:
		return ErrInvalidPassword
	case len(password) > 265:
		return ErrInvalidPassword
	}
	// Hash password
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(b)

	return nil
}
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}
