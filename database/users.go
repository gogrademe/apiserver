package database

import (
	. "github.com/Lanciv/GoGradeAPI/model"
	"errors"
	"log"
)

var (
	ErrUserOrPasswdIncorrect = errors.New("Username or password incorrect.")
)

const userFindEmailStmt = `
SELECT id, email, hashed_password, role, created_at, updated_at
FROM user_account WHERE email_lower = $1 and disabled = false
`
const userCreateStmt = `
INSERT INTO user_account (email, email_lower, hashed_password, role, created_at, updated_at)
VALUES($1,$2,$3,$4,$5,$6) RETURNING id
`

const userGetAllStmt = `
SELECT id, email, hashed_password, role, created_at, updated_at
FROM user_account WHERE disabled = false
`

func CreateUser(email string, password string, role string) (*User, error) {
	u := NewUser(email, role)

	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}

	u.UpdateTime()

	err = db.QueryRow(userCreateStmt, u.Email, u.EmailLower, u.HashedPassword,
		u.Role, u.CreatedAt, u.UpdatedAt).Scan(&u.Id)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// func GetUserEmail(email string) (*User, error) {
// 	u := &User{}
//
// 	err := db.Get(u, userFindEmailStmt, email)
//
// 	if err != nil {
// 		return nil, err
// 	}
// 	return u, nil
//
// }
func GetUserEmail(email string) (User, error) {
	var u User

	err := db.Get(&u, userFindEmailStmt, email)

	if err != nil {
		return u, err
	}
	return u, nil

}
func GetAllUsers() ([]*User, error) {
	users := []*User{}

	err := db.Select(&users, userGetAllStmt)
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
