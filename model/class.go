package model

type Class struct {
	Id         int64
	Name       string
	TeacherId  int64  `db:"teacher_id"`
	GradeLevel string `db:"grade_level"`
	Subject    string
	TimeStamp
}
