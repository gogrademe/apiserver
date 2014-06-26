package database

import (
	"errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type StudentRepo struct {
}

func CreateStudent(s *m.Student) error {
	var p *m.Person
	p, err := GetPerson(s.PersonID)
	if err != nil {
		return err
	}
	if p == nil {
		return errors.New("person doesn't exist")
	}
	if p.Profiles["Student"] != "" {
		return errors.New("student for person already exists.")
	}
	res, err := r.Table("students").Insert(s).RunWrite(sess)
	if err != nil {
		return err
	}
	s.ID = res.GeneratedKeys[0]

	p.Profiles["Student"] = s.ID
	err = UpdatePerson(p)

	return nil
}

// CreateStudentProfile Create a profile for a student.
// func CreateStudentProfile(pID int64, s *m.Student) error {
//
// 	s.PersonID = pID
//
// 	if !s.Validate() {
// 		return errors.New("Student not valid")
// 	}
// 	_, err := db.NamedExec(studentProfileSaveStmt, s)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
