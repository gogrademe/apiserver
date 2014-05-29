package database

import (
	. "bitbucket.org/lanciv/GoGradeAPI/model"
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	// jwt "github.com/dgrijalva/jwt-go"
	"log"
	"strings"
)

var (
	ErrUserOrPasswdIncorrect = errors.New("Username or password incorrect.")
)

const userFindEmailStmt = `
SELECT * FROM user_account where email_lower = $1 and disabled = false
`

func CreateUser(email string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	//TODO: Make fill in EmailLower!!
	emailLower := strings.ToLower(email)
	user := &User{Email: email, EmailLower: emailLower, HashedPassword: hashedPassword, Role: role}
	//TODO: Move this to a Validate() func.
	user.UpdateTime()

	err = db.QueryRow(`INSERT INTO user_account (email, email_lower, hashed_password, role, created_at, updated_at)
		VALUES($1,$2,$3,$4,$5,$6) RETURNING id`, user.Email, user.EmailLower, user.HashedPassword, user.Role, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)

	if err != nil {
		return nil, err
	}

	// lastId, err := res.LastInsertId()

	// user.Id = lastId
	return user, nil
}
func GetUserEmail(email string) (*User, error) {
	u := &User{}

	err := db.Get(u, userFindEmailStmt, email)

	if err != nil {
		return nil, err
	}
	return u, nil

}

// Verifies that the password matches the hashed password.
func VerifyPasswd(email, passwd string) (*User, error) {
	u, err := GetUserEmail(email)
	if err != nil {
		log.Println(err)
		return nil, ErrUserOrPasswdIncorrect
	}

	if bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(passwd)) != nil {

		return nil, ErrUserOrPasswdIncorrect
	}

	return u, nil
}

func GetAllUsers() ([]User, error) {
	users := []User{}

	err := db.Select(&users, "SELECT * FROM user_account WHERE disabled = false")
	if err != nil {
		// Check to make sure this error is okay. (Not a connection error)
		log.Println(err)
		return nil, errors.New("Couldn't find any users.")
	}

	return users, nil
}

func GetUserById(id int) *User {
	// user := &User{}

	// db.Find(user, id)
	// fmt.Println(id)
	return nil
}
