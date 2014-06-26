package store

import (
	"errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
	"log"
	"time"
)

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
)

func InitStores() {
	Students = NewStudentsStore()
	Classes = NewClassStore()
	People = NewPersonStore()
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
		insertTestData()
	}
	return
}

func createDatabase() {
	log.Println("DropDB")
	r.DbDrop(dbName).RunWrite(sess)
	log.Println("CreateDB")
	r.DbCreate(dbName).RunWrite(sess)
	return
}

func createTables() {
	log.Println("CreateTables")
	r.Db(dbName).TableCreate("users", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("classes", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("classTerms", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("assignments", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("people", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("students", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("teachers", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("parents", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("sessions", r.TableCreateOpts{Durability: "soft"}).Run(sess, r.RunOpts{NoReply: true})
	log.Println("CreateTablesDone")
	return
}

func createIndexes() {
	log.Println("CreateIndexes")
	r.Db(dbName).Table("users").IndexCreate("email").Run(sess, r.RunOpts{NoReply: true})
	log.Println("CreateIndexesDone")
	return
}

func insertTestData() {
	log.Println("InsertTestData")
	log.Println("Create User")
	CreateUser("test@test.com", "somePassword", "Admin")
	log.Println("Create People")
	createTestPeople()
	log.Println("TestDataDone")
	return
}

func createTestPeople() {
	log.Println("CreatePeople")
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
