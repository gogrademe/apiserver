package database

import (
	"errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
	"log"
)

var (
	//ErrUserOrPasswdIncorrect err for password incorrect
	ErrUserOrPasswdIncorrect = errors.New("Email or password incorrect.")

	//ErrUserAlreadyExists err for duplicate user
	ErrUserAlreadyExists = errors.New("User with email already exists.")
)

func userExist(email string) bool {
	row, _ := r.Table("users").Filter(r.Row.Field("email").Eq(email)).RunRow(sess)

	return !row.IsNil()
}

// CreateUser will create a user with a email, password and role.
func CreateUser(email string, password string, role string) (*m.User, error) {
	log.Println("1")
	u := m.NewUser(email, role)

	err := u.SetPassword(password)
	if err != nil {
		return nil, err
	}
	log.Println("2")
	u.UpdateTime()

	if userExist(email) {
		return nil, ErrUserAlreadyExists
	}
	log.Println("3")
	res, err := r.Table("users").Insert(u).RunWrite(sess)
	if err != nil {
		return nil, err
	}
	log.Println("4")
	u.ID = res.GeneratedKeys[0]

	return u, nil
}

// GetUserEmail return a single user that matches an email.
func GetUserEmail(email string) (m.User, error) {
	var u m.User

	row, err := r.Table("users").Filter(r.Row.Field("email").Eq(email)).RunRow(sess)
	if err != nil {
		return u, err
	}

	err = row.Scan(&u)
	return u, nil

}

// GetAllUsers return every user in the DB.
func GetAllUsers() ([]m.User, error) {
	users := []m.User{}

	rows, err := r.Table("users").Run(sess)
	if err != nil {
		// Check to make sure this error is okay. (Not a connection error)
		log.Println(err)
		return nil, err
	}

	err = rows.ScanAll(&users)
	if err != nil {
		// Check to make sure this error is okay. (Not a connection error)
		log.Println(err)
		return nil, err
	}
	return users, nil
}

// GetUserByID get a user by a ID.
func GetUserByID(id string) (m.User, error) {
	u := m.User{}

	row, err := r.Table("users").Get(id).Run(sess)
	if err != nil {
		return u, err
	}

	err = row.Scan(&u)
	if err != nil {
		return u, err
	}

	return u, nil
}
