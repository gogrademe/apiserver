package store

import (
	"errors"
	"log"
	"time"

	rh "github.com/Lanciv/rethinkHelper"
	r "github.com/dancannon/gorethink"
)

const testAddress = "localhost:28015"
const testDBName = "test_goGrade"

var (
	sess   *r.Session
	dbName string

	// DB Global DB object
	DB rh.DB

	Classes          = DB.NewCollection("classes")
	Terms            = DB.NewCollection("terms")
	SchoolYears      = DB.NewCollection("schoolYears")
	EnrollmentH      = DB.NewCollection("enrollments")
	People           = DB.NewCollection("people")
	UserH            = DB.NewCollection("users")
	AssignmentH      = DB.NewCollection("assignments")
	AssignmentGrades = DB.NewCollection("grades")

	Assignments     AssignmentStore
	AssignmentTypes AssignmentTypeStore

	Sessions    SessionStore
	Users       UserStore
	Enrollments EnrollmentStore

	// Errors
	ErrNotFound   = errors.New("record not found")
	ErrValidation = errors.New("validation error")

	tables = []string{"users", "classes", "enrollments", "terms", "assignments",
		"grades", "assignmentTypes", "people", "sessions"}
)

func init() {

	Assignments = NewAssignmentStore()
	AssignmentTypes = NewAssignmentTypeStore()

	// Classes
	Enrollments = NewEnrollmentStore()

	// Users/Auth
	Sessions = NewSessionStore()
	Users = NewUserStore()
}

// Connect establishes connection with rethinkDB
func Connect(address, database string) error {

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

	DB = rh.NewDBFromSession(sess)

	return nil
}

// SetupDB will be used to bootstrap the DB
func SetupDB(bootstrap, testData bool) {

	if bootstrap {
		log.Println("SetupDB: Bootstrapping...")
		createDatabase()
		createTables()
		createIndexes()
		log.Println("SetupDB: Bootstrap Done")
	}

	if testData {
		log.Println("SetupDB: Cleaning...")
		cleanTables()
		log.Println("SetupDB: Inserting Data...")
		insertTestData()
	}

	log.Println("SetupDB: Done")
}
