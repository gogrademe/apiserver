package model

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("password must be between 6 and 256 characters")
)

type User struct {
	ID             string `gorethink:"id,omitempty" json:"id"`
	Email          string `gorethink:"email,omitempty" json:"email"`
	EmailLower     string `gorethink:"emailLower,omitempty" json:"emailLower"`
	HashedPassword string `gorethink:"hashedPassword,omitempty" json:"-"`
	PersonID       string `gorethink:"personId,omitempty" json:"personId"`
	Role           string `gorethink:"role,omitempty" json:"role"`
	// base64 url encoded random hash.
	ActivationToken string `gorethink:"activationToken,omitempty" json:"-"`
	Disabled        bool   `gorethink:"disabled,omitempty" json:"disabled"`
	TimeStamp
}

func (u *User) GenActivationToken() error {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return err
	}
	u.ActivationToken = base64.URLEncoding.EncodeToString(rb)
	return nil
}

func NewUserFor(email, personID string) (*User, error) {
	user := &User{
		Email:      email,
		EmailLower: strings.ToLower(email),
		PersonID:   personID,
		Disabled:   true,
	}
	if err := user.GenActivationToken(); err != nil {
		return nil, err
	}
	user.UpdateTime()
	return user, nil
}

func NewUserForWithPassword(email, password, personID string) (*User, error) {
	user := &User{
		Email:      email,
		EmailLower: strings.ToLower(email),
		PersonID:   personID,
		Disabled:   false,
	}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	user.UpdateTime()
	return user, nil
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
