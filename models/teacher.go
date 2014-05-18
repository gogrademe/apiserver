package models

import (
	"errors"
)

type Teacher struct {
	FirstName  string `db:"firstName"`
	MiddleName string `db:"middleName"`
	LastName   string `db:"lastName"`
	AutoFields
}

func (t *Teacher) Validate() bool {
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
func (t *Teacher) Index() error {
	return nil
}
func CreateTeacher(t *Teacher) (*Teacher, error) {

	if !t.Validate() {
		return nil, errors.New("Teacher not valid.")
	}

	_, err := db.Exec("INSERT INTO teacher(firstName, middleName, lastName, updatedAt, createdAt) VALUES(?,?,?,?,?)", t.FirstName, t.MiddleName, t.LastName, t.UpdatedAt, t.CreatedAt)

	if err != nil {
		return nil, err
	}

	// t.Id, err = res.LastInsertId()

	if err != nil {
		return nil, err
	}
	return t, nil
}
