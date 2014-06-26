package repo

import (
	// "errors"
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type PersonRepo struct {
}

func NewPersonRepo() PersonRepo {
	return PersonRepo{}
}

func (pr *PersonRepo) Store(p *m.Person) error {
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

func (pr *PersonRepo) Update(p *m.Person) error {
	_, err := r.Table("people").Get(p.ID).Update(p).RunWrite(sess)
	if err != nil {
		return err
	}
	return nil
}

// GetAllPeople Return all people without their profiles.
func (pr *PersonRepo) FindAll() ([]m.Person, error) {
	people := []m.Person{}
	query := r.Table("people")
	// FIXME: Very expensive!
	query = query.Map(func(row r.RqlTerm) interface{} {
		return row.Merge(map[string]interface{}{
			"profiles": map[string]interface{}{
				"studentId": r.Table("students").Filter(func(s r.RqlTerm) r.RqlTerm {
					return s.Field("personId").Eq(row.Field("id"))
				}).CoerceTo("ARRAY").Map(func(s r.RqlTerm) interface{} {
					return s.Field("id")
				}).Nth(0).Default(""),
			},
		})
	})
	rows, err := query.Run(sess)
	if err != nil {
		return nil, err
	}

	err = rows.ScanAll(&people)
	if err != nil {
		return nil, err
	}

	return people, nil
}

// GetPerson get's a single person with it's profile(s)
func (pr *PersonRepo) FindById(id string) (*m.Person, error) {
	var p m.Person

	row, err := r.Table("people").Get(id).RunRow(sess)
	if err != nil {
		return &p, err
	}

	if row.IsNil() {
		return nil, nil
	}

	row.Scan(&p)
	return &p, nil
}
