package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

// PersonStore used to interact with person table.
type PersonStore struct {
	DefaultStore
}

func NewPersonStore() PersonStore {
	return PersonStore{DefaultStore: NewDefaultStore("people")}
}

//StoreMany store a slice of people
func (pr *PersonStore) StoreMany(p []m.Person) ([]string, error) {
	res, err := r.Table("people").Insert(p).RunWrite(sess)
	if err != nil {
		return nil, err
	}
	return res.GeneratedKeys, nil
}

// FindAll Return all people without their profiles.
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
				"teacherId": r.Table("teachers").Filter(func(s r.Term) r.Term {
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
