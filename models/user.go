package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	// "fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

var (
	ErrUserOrPasswdIncorrect = errors.New("Username or password incorrect.")
)

type User struct {
	Id             int64
	Email          string
	EmailLower     string `db:"emailLower"`
	HashedPassword []byte `db:"hashedPassword"`
	// PlainPasswd    string
	Role      string
	Disabled  bool
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
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

	res, err := db.Exec("INSERT INTO user(email, emailLower, hashedPassword, role) VALUES(?,?,?,?)", user.Email, user.EmailLower, user.HashedPassword, user.Role)

	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()

	user.Id = lastId
	return user, nil
}
func GetUserByEmail(email string) (*User, error) {
	u := &User{}

	err := db.Get(u, "SELECT * FROM user where emailLower = ? LIMIT 1", email)

	if err != nil {
		return nil, err
	}
	return u, nil

}

// Verifies that the password matches the hashed password.
func VerifyPasswd(email, passwd string) (*User, error) {
	u, err := GetUserByEmail(email)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v",u)

	// db.Where("email = ?", email).First(u)

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

func GetAllUsers() []User {
	users := []User{}

	err := db.Select(&users, "SELECT * FROM user WHERE disabled = 0")
	if err != nil {
		panic(err)
	}
	// db.Find(&users)

	return users
}

func GetUserById(id int) *User {
	// user := &User{}

	// db.Find(user, id)
	// fmt.Println(id)
	return nil
}
