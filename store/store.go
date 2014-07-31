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
	// DB
	sess   *r.Session
	dbName string

	DB rh.DB

	// Tables
	Parents  ParentStore
	Teachers TeacherStore
	Students StudentStore

	// Classes     ClassStore
	Enrollments EnrollmentStore

	Classes          = DB.NewCollection("classes")
	Terms            = DB.NewCollection("terms")
	EnrollmentH      = DB.NewCollection("enrollments")
	People           = DB.NewCollection("people")
	AssignmentH      = DB.NewCollection("assignments")
	AssignmentGrades = DB.NewCollection("grades")

	Assignments     AssignmentStore
	AssignmentTypes AssignmentTypeStore

	Sessions SessionStore
	Users    UserStore

	// Errors
	ErrNotFound   = errors.New("record not found")
	ErrValidation = errors.New("validation error")

	tables = []string{"users", "classes", "enrollments", "terms", "assignments",
		"grades", "assignmentTypes", "people", "students",
		"teachers", "parents", "sessions"}
)

func init() {
	Teachers = NewTeacherStore()
	Parents = NewParentStore()
	Students = NewStudentStore()

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
func SetupDB(testData bool) {
	log.Println("SetupDB: Start")
	createDatabase()
	createTables()
	createIndexes()

	if testData {
		cleanTables()
		insertTestData()
	}

	log.Println("SetupDB: Done")
}
