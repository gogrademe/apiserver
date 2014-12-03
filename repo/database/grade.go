package database

import r "github.com/dancannon/gorethink"

type Gradestore struct {
	r.Session
}

func NewGradestore(s r.Session) *Gradestore {
	return &Gradestore{s}
}

// func (db *Gradestore) PostGrade(grade *model.AssignmentGrade) error {
//
//
// 	return
// }
