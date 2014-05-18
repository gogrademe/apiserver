package models

import (
	"errors"
)

type Person struct {
	FirstName        string `db:"firstName"`
	MiddleName       string `db:"middleName"`
	LastName         string `db:"lastName"`
	StudentProfileId int    `db:"studentProfileId"`
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

	_, err := db.Exec("INSERT INTO person(first_name, middle_name, last_name, updated_at, created_at) VALUES($1,$2,$3,$4,$5)", t.FirstName, t.MiddleName, t.LastName, t.UpdatedAt, t.CreatedAt)

	if err != nil {
		return nil, err
	}

	// t.Id, err = res.LastInsertId()

	if err != nil {
		return nil, err
	}
	return t, nil
}
