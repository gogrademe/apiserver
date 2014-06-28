package store

import (
	"errors"
	"time"

	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
	"log"
)

// var log = logrus.New()
var (
	sess        *r.Session
	dbName      string
	errNotFound = errors.New("record not found")
)

const testAddress = "localhost:28015"
const testDBName = "test_goGrade"

var (
	Students StudentsStore
	Classes  ClassStore
	People   PersonStore
	Sessions SessionStore
	Users    UserStore
)

func InitStores() {
	Students = NewStudentsStore()
	Classes = NewClassStore()
	People = NewPersonStore()
	Sessions = NewSessionStore()
	Users = NewUserStore()
}

// Connect establishes connection with rethinkDB
func Connect(address, database string) error {
	log.Println("Connecting")
	dbName = database
	var err error
	sess, err = r.Connect(r.ConnectOpts{
		Address:     address,
		Database:    dbName,
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})
	if err != nil {
		return err
	}
	// createTables()
	log.Println("Connected")
	return nil
}

// SetupDB will be used to bootstrap the DB
func SetupDB(testData bool) {
	log.Println("SetupDB")
	createDatabase()
	createTables()
	createIndexes()

	if testData {
		cleanTables()
		insertTestData()
	}
	return
}

func createDatabase() {
	log.Println("DropDB")

	// r.DbDrop(dbName).RunWrite(sess)

	// log.Println(r.DbList().Run(sess))
	log.Println("CreateDB")
	r.DbCreate(dbName).RunWrite(sess)
}

func createTables() {
	log.Println("CreateTables")
	r.Db(dbName).TableCreate("users").Run(sess)
	r.Db(dbName).TableCreate("classes").Run(sess)
	r.Db(dbName).TableCreate("classTerms").Run(sess)
	r.Db(dbName).TableCreate("assignments").Run(sess)
	r.Db(dbName).TableCreate("people").Run(sess)
	r.Db(dbName).TableCreate("students").Run(sess)
	r.Db(dbName).TableCreate("teachers").Run(sess)
	r.Db(dbName).TableCreate("parents").Run(sess)
	r.Db(dbName).TableCreate("sessions").Run(sess)
	log.Println("CreateTablesDone")
}

func cleanTables() {
	r.Table("users").Delete().Run(sess)
	r.Table("classes").Delete().Run(sess)
	r.Table("classTerms").Delete().Run(sess)
	r.Table("assignments").Delete().Run(sess)
	r.Table("people").Delete().Run(sess)
	r.Table("students").Delete().Run(sess)
	r.Table("teachers").Delete().Run(sess)
	r.Table("parents").Delete().Run(sess)
	r.Table("sessions").Delete().Run(sess)
}

func createIndexes() {
	log.Println("CreateIndexes")
	r.Db(dbName).Table("users").IndexCreate("email").Run(sess)
	log.Println("CreateIndexesDone")
}

func insertTestData() {
	log.Println("InsertTestData")
	insertTestUsers()
	insertTestPeople()
	log.Println("TestDataDone")
}
func insertTestUsers() {
	log.Println("insertTestUsers")
	u, _ := m.NewUser("test@test.com", "somePassword", "Admin")
	Users.Store(u)
}

func insertTestPeople() {
	log.Println("insertTestPeople")
	CreatePeople([]m.Person{
		m.Person{
			FirstName:  "Jon",
			MiddleName: "David",
			LastName:   "Bush",
		},
		m.Person{
			FirstName: "Frankie",
			LastName:  "Bagnardi",
		},
		m.Person{
			FirstName: "Adam",
			LastName:  "Price",
		},
		m.Person{
			FirstName:  "Jake",
			MiddleName: "Matthew",
			LastName:   "Price",
		},
		m.Person{
			FirstName: "Matthew",
			LastName:  "Aitchison",
		},
		m.Person{
			FirstName: "Natalie",
			LastName:  "Aitchison",
		},
		m.Person{
			FirstName: "Nicole",
			LastName:  "Aitchison",
		},
		m.Person{
			FirstName: "Angel",
			LastName:  "Heredia",
		},
	})
	log.Println("CreatePeopleDone")
	return
}
