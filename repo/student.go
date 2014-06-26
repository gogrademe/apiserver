package repo

import (
	// "errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type StudentsRepo struct {
}

func NewStudentsRepo() StudentsRepo {
	return StudentsRepo{}
}

func (sr *StudentsRepo) Store(s *m.Student) error {
	res, err := r.Table("students").Insert(s).RunWrite(sess)
	if err != nil {
		return err
	}
	s.ID = res.GeneratedKeys[0]
	return nil
}

// func CreateStudent(s *m.Student) error {
// 	var p *m.Person
// 	p, err := GetPerson(s.PersonID)
// 	if err != nil {
// 		return err
// 	}
// 	if p == nil {
// 		return errors.New("person doesn't exist")
// 	}
// 	if p.Profiles["Student"] != "" {
// 		return errors.New("student for person already exists.")
// 	}
// 	res, err := r.Table("students").Insert(s).RunWrite(sess)
// 	if err != nil {
// 		return err
// 	}
// 	s.ID = res.GeneratedKeys[0]
//
// 	p.Profiles["Student"] = s.ID
// 	err = UpdatePerson(p)
//
// 	return nil
// }
