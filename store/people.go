package store

import (
	// "errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type PersonStore struct {
}

func NewPersonStore() PersonStore {
	return PersonStore{}
}

func (pr *PersonStore) Store(p *m.Person) error {
	res, err := r.Table("people").Insert(p).RunWrite(sess)
	if err != nil {
		return err
	}
	p.ID = res.GeneratedKeys[0]
	return nil
}
func CreatePeople(p []m.Person) error {
	_, err := r.Table("people").Insert(p).RunWrite(sess)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PersonStore) Update(p *m.Person) error {
	_, err := r.Table("people").Get(p.ID).Update(p).RunWrite(sess)
	if err != nil {
		return err
	}
	return nil
}

// GetAllPeople Return all people without their profiles.
func (pr *PersonStore) FindAll() ([]m.Person, error) {
	people := []m.Person{}
	query := r.Table("people")
	// FIXME: Very expensive!
	query = query.Map(func(row r.Term) interface{} {
		return row.Merge(map[string]interface{}{
			"profiles": map[string]interface{}{
				"studentId": r.Table("students").Filter(func(s r.Term) r.Term {
					return s.Field("personId").Eq(row.Field("id"))
				}).CoerceTo("ARRAY").Map(func(s r.Term) interface{} {
					return s.Field("id")
				}).Nth(0).Default(""),
			},
		})
	})
	res, err := query.Run(sess)
	if err != nil {
		return nil, err
	}

	err = res.All(&people)
	if err != nil {
		return nil, err
	}

	return people, nil
}

// GetPerson get's a single person with it's profile(s)
func (pr *PersonStore) FindById(id string) (*m.Person, error) {
	var p m.Person

	res, err := r.Table("people").Get(id).Run(sess)
	if err != nil {
		return &p, err
	}

	if res.IsNil() {
		return nil, nil
	}

	res.One(&p)
	return &p, nil
}
