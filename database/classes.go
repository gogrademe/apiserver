package database

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

func GetAllClasses() ([]m.Class, error) {

	classes := []m.Class{}

	rows, err := r.Table("classes").Run(sess)

	if err != nil {
		return nil, err
	}

	err = rows.ScanAll(&classes)
	return classes, nil
}
