package models

type StudentProfile struct {
	GradeLevel string `db:"grade_level"`
	AutoFields
}
