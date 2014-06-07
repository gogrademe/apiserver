package database

import (
	. "github.com/Lanciv/GoGradeAPI/model"
)

func GetAllClasses() ([]Class, error) {

	classes := []Class{}

	err := db.Get(&classes, "SELECT * FROM class limit 1")

	if err != nil {
		return nil, err
	}

	// classes := []Class{}
	// db.Find(&classes)

	return classes, nil
}
