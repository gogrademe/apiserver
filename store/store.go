package store

import (
	"errors"
	"time"

	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

const testAddress = "localhost:28015"
const testDBName = "test_goGrade"

var (
	// DB
	sess   *r.Session
	dbName string

	// Tables
	Parents  ParentStore
	Teachers TeacherStore
	Students StudentStore

	Classes    ClassStore
	ClassTerms ClassTermStore
	People     PersonStore

	Assignments AssignmentStore

	Sessions SessionStore
	Users    UserStore

	// Errors
	ErrNotFound   = errors.New("record not found")
	ErrValidation = errors.New("validation error")
)

func init() {
	// People
	People = NewPersonStore()
	Teachers = NewTeacherStore()
	Parents = NewParentStore()
	Students = NewStudentStore()

	Assignments = NewAssignmentStore()

	// Classes
	Classes = NewClassStore()
	ClassTerms = NewClassTermStore()

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

	return nil
}

// SetupDB will be used to bootstrap the DB
func SetupDB(testData bool) {

	createDatabase()
	createTables()
	createIndexes()

	if testData {
		cleanTables()
		insertTestData()
	}
}

func createDatabase() {

	r.DbCreate(dbName).RunWrite(sess)
}

func createTables() {
	r.Db(dbName).TableCreate("users").Run(sess)
	r.Db(dbName).TableCreate("classes").Run(sess)
	r.Db(dbName).TableCreate("classTerms").Run(sess)
	r.Db(dbName).TableCreate("assignments").Run(sess)
	r.Db(dbName).TableCreate("people").Run(sess)
	r.Db(dbName).TableCreate("students").Run(sess)
	r.Db(dbName).TableCreate("teachers").Run(sess)
	r.Db(dbName).TableCreate("parents").Run(sess)
	r.Db(dbName).TableCreate("sessions").Run(sess)

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

	r.Db(dbName).Table("users").IndexCreate("email").Run(sess)

}

func insertTestData() {

	insertTestUsers()
	peopleIds := insertTestPeople()
	insertTestStudents(peopleIds)
	insertTestTeachers(peopleIds)
	insertTestClasses()

}
func insertTestUsers() {

	u, _ := m.NewUser("test@test.com", "somePassword", "Admin")
	Users.Store(u)
}

func insertTestClasses() {
	for _, c := range []m.Class{
		m.Class{
			Name:       "Algebra II",
			Subject:    "Math",
			GradeLevel: "12th",
		},
		m.Class{
			Name:       "Geography",
			Subject:    "Social Studies",
			GradeLevel: "12th",
		},
		m.Class{
			Name:       "Economics",
			Subject:    "Social Studies",
			GradeLevel: "12th",
		},
		m.Class{
			Name:       "Writing",
			Subject:    "Language Arts",
			GradeLevel: "8th",
		},
		m.Class{
			Name:       "Reading",
			Subject:    "Language Arts",
			GradeLevel: "9th",
		},
	} {
		Classes.Store(&c)
	}
}

func insertTestStudents(peopleIds []string) {
	for _, id := range peopleIds {
		Students.Store(&m.Student{
			PersonID:   id,
			GradeLevel: "12th",
		})
	}
}
func insertTestTeachers(peopleIds []string) {
	for _, id := range peopleIds {
		Teachers.Store(&m.Teacher{
			PersonID: id,
			Email:    "test@test.com",
		})
	}
}
func insertTestPeople() []string {

	ids, _ := People.StoreMany([]m.Person{
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

	return ids
}
