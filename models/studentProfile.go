package models

type StudentProfile struct {
	GradeLevel string `db:"gradeLevel"`
	AutoFields
}
