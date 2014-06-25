package database

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
	"log"
	"time"
)

var (
	sess   *r.Session
	dbName string
)

const testAddress = "localhost:28015"
const testDBName = "test_goGrade"

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
	r.Db(dbName).TableCreate("users").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("classes").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("classTerms").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("assignments").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("people").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("students").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("teachers").Run(sess, r.RunOpts{NoReply: true})
	r.Db(dbName).TableCreate("parents").Run(sess, r.RunOpts{NoReply: true})
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
	CreateUser("test@test.com", "somePassword", "Admin")

	CreateUser("test@test.com", "somePassword", "Admin")
	createTestPeople()
	log.Println("TestDataDone")
	return
}

func createTestPeople() {
	log.Println("CreatePeople")
	log.Println("1")
	CreatePerson(&m.Person{
		FirstName:  "Jon",
		MiddleName: "David",
		LastName:   "Bush",
	})
	log.Println("2")
	CreatePerson(&m.Person{
		FirstName: "Frankie",
		LastName:  "Bagnardi",
	})
	log.Println("3")
	CreatePerson(&m.Person{
		FirstName: "Adam",
		LastName:  "Price",
	})
	log.Println("4")
	CreatePerson(&m.Person{
		FirstName:  "Jake",
		MiddleName: "Matthew",
		LastName:   "Price",
	})
	log.Println("CreatePeopleDone")
	return
}
