package models

import (
	"errors"
	"log"
)

type Person struct {
	Id             int64
	FirstName      string          `db:"first_name"`
	MiddleName     string          `db:"middle_name"`
	LastName       string          `db:"last_name"`
	StudentProfile *StudentProfile `json:",omitempty"`
	TimeStamp
}

func (t *Person) Validate() bool {
	if t.FirstName == "" {
		return false
	}
	if t.LastName == "" {
		return false
	}
	t.UpdateTime()
	return true
}

// Index in elastic
func (t *Person) Index() error {
	return nil
}
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
	err := db.Select(&people, `SELECT * FROM person LEFT OUTER JOIN student_profile ON (person.Id = student_profile.person_id)`)

	if err != nil {
		return nil, err
	}

	return people, nil
}
