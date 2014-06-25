package database

import (
	"errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

func CreatePerson(p *m.Person) error {

	if !p.Validate() {
		return errors.New("Person not valid")
	}
	// Create a m.Person and get its ID
	// nstmt, err := db.PrepareNamed(personSaveStmt)
	// if err != nil {
	// 	return err
	// }
	// err = nstmt.QueryRow(p).Scan(&p.ID)
	_, err := r.Table("people").Insert(p).RunWrite(sess)
	if err != nil {
		return err
	}

	return nil
}

// GetAllPeople Return all people without their profiles.
func GetAllPeople() ([]m.Person, error) {
	people := []m.Person{}

	// err := db.Select(&people, peopleGetAllStmt)
	rows, err := r.Table("people").Run(sess)
	if err != nil {
		return nil, err
	}

	err = rows.ScanAll(&people)
	if err != nil {
		return nil, err
	}

	return people, nil
}

// GetPerson get's a single person with it's profile(s)
func GetPerson(id string) (*m.Person, error) {
	var p m.Person

	row, err := r.Table("people").Get(id).RunRow(sess)
	if err != nil {
		return &p, err
	}

	row.Scan(p)
	return &p, nil
}
