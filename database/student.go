package database

import (
	// "errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

//StudentProfileForPerson Get's a student profile from a Person.ID
// func StudentProfileForPerson(id int) (*m.Student, error) {
// 	var s m.Student
//
// 	err := db.Get(&s, studentGetForPerson, id)
//
// 	if err != nil {
// 		return &s, err
// 	}
// 	return &s, nil
// }

func CreateStudent(s *m.Student) error {
	res, err := r.Table("students").Insert(s).RunWrite(sess)
	if err != nil {
		return err
	}
	s.ID = res.GeneratedKeys[0]

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
