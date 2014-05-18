package models

import (
	"database/sql"
	"errors"
)

type Person struct {
	FirstName        string        `db:"first_name"`
	MiddleName       string        `db:"middle_name"`
	LastName         string        `db:"last_name"`
	StudentProfileId sql.NullInt64 `db:"student_profile_id"`
	AutoFields
}

func (t *Person) Validate() bool {
	if t.FirstName == "" {
		return false
	}
	if t.LastName == "" {
		return false
	}
	t.UpdateAuto()
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

	return t, nil
}
func GetAllPeople() ([]Person, error) {
	people := []Person{}
	err := db.Select(&people, `SELECT * FROM person`)

	if err != nil {
		return nil, err
	}

	return people, nil
}
