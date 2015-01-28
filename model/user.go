package model

import (
	"errors"
	"strings"

	"code.google.com/p/go.crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("password must be between 6 and 256 characters")
)

type User struct {
	ID             string `gorethink:"id,omitempty"json:"id"`
	Email          string `gorethink:"email,omitempty"json:"email"`
	EmailLower     string `gorethink:"emailLower,omitempty"json:"emailLower"`
	HashedPassword string `gorethink:"hashedPassword,omitempty"json:"-"`
	PersonID       string `gorethink:"personId,omitempty"json:"personId"`
	Role           string `gorethink:"role,omitempty"json:"role"`
	Disabled       bool   `gorethink:"disabled,omitempty"json:"disabled"`
	TimeStamp
}

func NewUserFor(email, password, personID string) (*User, error) {
	user := User{
		Email:      email,
		EmailLower: strings.ToLower(email),
		PersonID:   personID,
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
