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
