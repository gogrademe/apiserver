package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	// "fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"time"
)

var (
	ErrUserOrPasswdIncorrect = errors.New("Username or password incorrect.")
)

type User struct {
	Email          string
	EmailLower     string `db:"email_lower"`
	HashedPassword []byte `db:"hashed_password"`
	Role           string
	Disabled       bool
	AutoFields
}

// func GetUsersCount() int {
// 	var count int
// 	// err := db.QueryRow(`select COUNT(*) from user`).Scan(&count)
// 	err := db.Model(User{}).Count(&count)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(count)
// 	return count
// }

func CreateUser(email string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	//TODO: Make fill in EmailLower!!
	emailLower := strings.ToLower(email)
	user := &User{Email: email, EmailLower: emailLower, HashedPassword: hashedPassword, Role: role}
	//TODO: Move this to a Validate() func.
	user.UpdateAuto()

	err = db.QueryRow(`INSERT INTO user_account (email, email_lower, hashed_password, role, created_at, updated_at)
		VALUES($1,$2,$3,$4,$5,$6) RETURNING id`, user.Email, user.EmailLower, user.HashedPassword, user.Role, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)

	if err != nil {
		return nil, err
	}

	// lastId, err := res.LastInsertId()

	// user.Id = lastId
	return user, nil
}
func GetUserByEmail(email string) (*User, error) {
	u := &User{}

	err := db.Get(u, "SELECT * FROM user_account where email_lower = $1 LIMIT 1", email)

	if err != nil {
		return nil, err
	}
	return u, nil

}

// Verifies that the password matches the hashed password.
func VerifyPasswd(email, passwd string) (*User, error) {
	u, err := GetUserByEmail(email)
	if err != nil {
		return nil, ErrUserOrPasswdIncorrect
	}

	if bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(passwd)) != nil {

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

func GetAllUsers() ([]User, error) {
	users := []User{}

	err := db.Select(&users, "SELECT * FROM user_account WHERE disabled = 0")
	if err != nil {
		// Check to make sure this error is okay. (Not a connection error)
		log.Println(err)
		return nil, errors.New("Couldn't find any users.")
	}
	// db.Find(&users)

	return users, nil
}

func GetUserById(id int) *User {
	// user := &User{}

	// db.Find(user, id)
	// fmt.Println(id)
	return nil
}
