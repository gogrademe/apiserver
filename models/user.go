package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	ErrUserOrPasswdIncorrect = errors.New("Username or password incorrect.")
)

type User struct {
	Id           int64
	Email        string `sql:"not null;unique"`
	EmailLower   string `sql:"not null;unique"`
	HashedPasswd []byte `sql:"not null"`
	PlainPasswd  string `sql:"-"`
	Role         string `sql:"not null"`
	Disabled     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetUsersCount() int {
	var count int
	// err := db.QueryRow(`select COUNT(*) from user`).Scan(&count)
	err := db.Model(User{}).Count(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	return count
}

func CreateUser(email string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	//TODO: Make fill in EmailLower!!
	user := User{Email: email, EmailLower: "test", HashedPasswd: hashedPassword, Role: role}
	db.Save(&user)

	if err != nil {
		return nil, err
	}
	// TODO: We are supposed to return a user here!
	return nil, nil
}

// Verifies that the password matches the hashed password.
func VerifyPasswd(email, passwd string) (*User, error) {
	u := &User{}

	db.Where("email = ?", email).First(u)

	if bcrypt.CompareHashAndPassword([]byte(u.HashedPasswd), []byte(passwd)) != nil {

		return nil, ErrUserOrPasswdIncorrect
	}

	return u, nil
}

// Create a token for the user after we verified their password.
// TODO: Store this in a db? This would be helpful if we would like to invalidate a login.
func (a *User) CreateToken() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims["Id"] = a.Id
	token.Claims["Email"] = a.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// TODO: Move this to a config file.
	tokenString, err := token.SignedString([]byte("someRandomSigningKey"))
	if err != nil {
		return "", err
	}
	// Sould we just return a Token instead of a string???
	return tokenString, err
}
func GetAllUsers() []User {
	users := []User{}

	db.Find(&users)

	return users
}
func GetUserById(id int) *User {
	user := &User{}

	db.Find(user, id)
	fmt.Println(id)
	return user
}
