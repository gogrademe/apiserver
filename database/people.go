package database

import (
	"errors"
	m "github.com/Lanciv/GoGradeAPI/model"
)

const peopleGetAllStmt = `
SELECT id, first_name, middle_name, last_name, created_at, updated_at
FROM person
`
const peopleFindIDStmt = `
SELECT id, first_name, middle_name, last_name, created_at, updated_at
from person where id = $1
`
const personSaveStmt = `
INSERT INTO person(first_name, middle_name, last_name, updated_at, created_at)
VALUES(:first_name,:middle_name,:last_name,:updated_at,:created_at) RETURNING id
`

func CreatePerson(p *m.Person) error {

	if !p.Validate() {
		return errors.New("Person not valid")
	}
	// Create a m.Person and get its ID
	nstmt, err := db.PrepareNamed(personSaveStmt)
	if err != nil {
		return err
	}
	err = nstmt.QueryRow(p).Scan(&p.Id)
	if err != nil {
		return err
	}

	return nil
}

// func InsertPersonProfile(p *m.PersonProfile) error {
// 	if !p.Person.Validate() {
// 		return errors.New("Person not valid")
// 	}
// 	tx, err := db.Beginx()
//
// 	rows, err = tx.NamedQuery(personSaveStmt, p.Person)
//
// 	if p.StudentProfile != nil {
// 		err = tx.QueryRowx(studentProfileSaveStmt, p.Person.Id, p.StudentProfile.GradeLevel, p.StudentProfile.UpdatedAt, p.StudentProfile.CreatedAt).Scan(&p.StudentProfile.PersonID)
// 	}
// 	err = tx.Commit()
//
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// CreatePerson Creates a Person in the DB. Can also create Profiles for Students.
// func CreatePerson(p *m.Person) (*m.Person, error) {
//
// 	if !p.Validate() {
// 		return nil, errors.New("Person not valid")
// 	}
// 	// Create a m.Person and get its ID
// 	err := db.QueryRow(`INSERT INTO person(first_name, middle_name, last_name, updated_at, created_at)
//     VALUES($1,$2,$3,$4,$5) RETURNING id`, p.FirstName, p.MiddleName, p.LastName, p.UpdatedAt, p.CreatedAt).Scan(&p.Id)
//
// 	result, err := db.NamedQuery(personSaveStmt, p)
// 	log.Println(result)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Create StudentProfile
// 	// if p.StudentProfile != nil {
// 	// 	p.StudentProfile.PersonId = p.Id
// 	// 	_, err = CreateStudentProfile(p.StudentProfile)
// 	// }
//
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	p.StudentProfile = nil
// 	// 	return p, errors.New("Failed to create student profile")
// 	// }
// 	return p, nil
// }

// GetAllPeople Return all people without their profiles.
func GetAllPeople() ([]m.Person, error) {
	people := []m.Person{}

	err := db.Select(&people, peopleGetAllStmt)

	if err != nil {
		return nil, err
	}

	return people, nil
}

// GetPerson get's a single person with it's profile(s)
func GetPerson(id int) (*m.Person, error) {
	var p m.Person

	err := db.Get(&p, peopleFindIDStmt, id)

	if err != nil {
		return &p, err
	}

	return &p, nil
}
