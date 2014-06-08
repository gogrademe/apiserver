package database

import (
	"errors"
	m "github.com/Lanciv/GoGradeAPI/model"
)

const studentGetForPerson = `
SELECT person_id, grade_level, created_at, updated_at
FROM student_profile where person_id = $1;
`

const studentProfileSaveStmt = `
INSERT INTO student_profile(person_id,grade_level, updated_at, created_at)
VALUES(:person_id,:grade_level,:updated_at,:created_at) RETURNING person_id
`

// StudentProfileForPerson Get's a student profile from a Person.ID
func StudentProfileForPerson(id int) (*m.StudentProfile, error) {
	var sP m.StudentProfile

	err := db.Get(&sP, studentGetForPerson, id)

	if err != nil {
		return &sP, err
	}
	return &sP, nil
}

// CreateStudentProfile Create a profile for a student.
// func CreateStudentProfile(s *m.StudentProfile) (*m.StudentProfile, error) {
// 	_, err := db.Exec(studentProfileSaveStmt, s.PersonID, s.GradeLevel, s.UpdatedAt, s.CreatedAt)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return s, nil
// }

// CreateStudentProfile Create a profile for a student.
func CreateStudentProfile(pID int64, sP *m.StudentProfile) error {

	sP.PersonID = pID

	if !sP.Validate() {
		return errors.New("Student not valid")
	}
	_, err := db.NamedExec(studentProfileSaveStmt, sP)
	if err != nil {
		return err
	}

	return nil
}
