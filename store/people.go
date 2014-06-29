package store

import (
	"fmt"

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

// Store a single person
// func (pr *PersonStore) Store(p *m.Person) error {
// 	res, err := r.Table("people").Insert(p).RunWrite(sess)
// 	if err != nil {
// 		return err
// 	}
// 	if p.ID == "" && len(res.GeneratedKeys) == 1 {
// 		p.ID = res.GeneratedKeys[0]
// 	}
// 	return nil
// }

//StoreMany store a slice of people
func (pr *PersonStore) StoreMany(p []m.Person) error {
	_, err := r.Table("people").Insert(p).RunWrite(sess)
	if err != nil {
		return err
	}
	return nil
}

//Update update a person
func (pr *PersonStore) Update(p *m.Person) error {
	res, err := r.Table("people").Get(p.ID).Update(p).RunWrite(sess)
	if err != nil {
		return err
	}
	fmt.Println(res)
	if res.Replaced == 0 {
		return ErrNotFound
	}
	return nil
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

// FindByID get's a single person with it's profile(s)
func (pr *PersonStore) FindByID(id string) (*m.Person, error) {
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
