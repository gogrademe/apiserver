package database

import (
	. "github.com/Lanciv/GoGradeAPI/model"
	"errors"
	"log"
)

const peopleGetAllStmt = `
SELECT id, first_name, middle_name, last_name, created_at, updated_at
FROM person
`

func CreatePerson(t *Person) (*Person, error) {

	if !t.Validate() {
		return nil, errors.New("Person not valid.")
	}

	err := db.QueryRow(`INSERT INTO person(first_name, middle_name, last_name, updated_at, created_at)
    VALUES($1,$2,$3,$4,$5) RETURNING id`, t.FirstName, t.MiddleName, t.LastName, t.UpdatedAt, t.CreatedAt).Scan(&t.Id)

	if err != nil {
		return nil, err
	}

	if t.StudentProfile != nil {
		t.StudentProfile.PersonId = t.Id
		_, err = CreateStudentProfile(t.StudentProfile)
	}

	if err != nil {
		log.Println(err)
		t.StudentProfile = nil
		return t, errors.New("Failed to create student profile")
	}

	return t, nil
}
func GetAllPeople() ([]Person, error) {
	people := []Person{}

	err := db.Select(&people, peopleGetAllStmt)

	if err != nil {
		return nil, err
	}

	return people, nil
}
