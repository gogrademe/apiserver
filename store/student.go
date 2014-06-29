package store

import (
	// "errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type StudentStore struct {
}

func NewStudentStore() StudentStore {
	return StudentStore{}
}

func (sr *StudentStore) Store(s *m.Student) error {
	res, err := r.Table("students").Insert(s).RunWrite(sess)
	if err != nil {
		return err
	}
	if s.ID == "" && len(res.GeneratedKeys) == 1 {
		s.ID = res.GeneratedKeys[0]
	}
	return nil
}

// FindByID get's a single student with it's profile(s)
func (pr *StudentStore) FindByID(id string) (*m.Student, error) {
	var p m.Student

	res, err := r.Table("student").Get(id).Run(sess)
	if err != nil {
		return &p, err
	}

	if res.IsNil() {
		return nil, nil
	}

	res.One(&p)
	return &p, nil
}
