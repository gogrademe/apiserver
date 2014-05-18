package models

import (
	"time"
)

type Class struct {
	Id         int64
	Name       string
	TeacherId  int64  `db:"teacher_id"`
	GradeLevel string `db:"grade_level"`
	Subject    string
	AutoFields
}

func (c *Class) Validate() bool {
	return false
}

func GetAllClasses() ([]Class, error) {

	// classes := []Class{}
	// db.Find(&classes)

	return nil, nil
}
