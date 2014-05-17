package models

import (
	"time"
)

type Class struct {
	Id         int64
	Name       string
	TeacherId  int64
	GradeLevel string
	Subject    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt time.Time
	Archived   bool
}

// func CreateClass(class *Class) error {

// 	db.Save(&class)

// 	return nil
// }

func GetAllClasses() ([]Class, error) {

	// classes := []Class{}
	// db.Find(&classes)

	return nil, nil
}
