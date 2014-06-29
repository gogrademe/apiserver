package store

import (
	r "github.com/dancannon/gorethink"
)

// Storer interface
// type Storer interface {
// 	Store(interface{}) (string, error)
// 	Update(string, interface{}) error
// 	FindAll(interface{}) error
// 	FindByID(string) (interface{}, error)
// }

type DefaultStore struct {
	TableName string
}

func NewDefaultStore(tableName string) DefaultStore {
	return DefaultStore{TableName: tableName}
}

func (d *DefaultStore) Store(v interface{}) (string, error) {
	res, err := r.Table(d.TableName).Insert(v).RunWrite(sess)
	if err != nil {
		return "", err
	}
	if len(res.GeneratedKeys) == 1 {
		return res.GeneratedKeys[0], nil
	}
	return "", nil
}

func (d *DefaultStore) Update(id string, v interface{}) error {
	res, err := r.Table(d.TableName).Get(id).Update(v).RunWrite(sess)
	if err != nil {
		return err
	}

	if res.Replaced == 0 {
		return ErrNotFound
	}
	return nil
}

func (d *DefaultStore) FindAll(data interface{}) error {

	res, err := r.Table(d.TableName).Run(sess)
	if err != nil {
		return err
	}
	return res.All(data)
}

func (d *DefaultStore) FindByID(data interface{}, id string) error {

	res, err := r.Table(d.TableName).Get(id).Run(sess)
	if err != nil {
		return err
	}

	if res.IsNil() {
		return ErrNotFound
	}

	return res.One(data)
}

//
// //Update update a person
// func (pr *PersonStore) Update(p *m.Person) error {
// 	res, err := r.Table("people").Get(p.ID).Update(p).RunWrite(sess)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(res)
// 	if res.Replaced == 0 {
// 		return ErrNotFound
// 	}
// 	return nil
// }
//
// // FindAll Return all people without their profiles.
// func (pr *PersonStore) FindAll() ([]m.Person, error) {
// 	people := []m.Person{}
// 	query := r.Table("people")
// 	// FIXME: Very expensive!
// 	query = query.Map(func(row r.Term) interface{} {
// 		return row.Merge(map[string]interface{}{
// 			"profiles": map[string]interface{}{
// 				"studentId": r.Table("students").Filter(func(s r.Term) r.Term {
// 					return s.Field("personId").Eq(row.Field("id"))
// 				}).CoerceTo("ARRAY").Map(func(s r.Term) interface{} {
// 					return s.Field("id")
// 				}).Nth(0).Default(""),
// 			},
// 		})
// 	})
// 	res, err := query.Run(sess)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = res.All(&people)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return people, nil
// }
//
// // FindByID get's a single person with it's profile(s)
// func (pr *PersonStore) FindByID(id string) (*m.Person, error) {
// 	var p m.Person
//
// 	res, err := r.Table("people").Get(id).Run(sess)
// 	if err != nil {
// 		return &p, err
// 	}
//
// 	if res.IsNil() {
// 		return nil, nil
// 	}
//
// 	res.One(&p)
// 	return &p, nil
// }
