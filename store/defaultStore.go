package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
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

func (d *DefaultStore) Store(v m.Model) (string, error) {
	v.UpdateTime()
	res, err := r.Table(d.TableName).Insert(v).RunWrite(sess)
	if err != nil {
		return "", err
	}
	if len(res.GeneratedKeys) == 1 {
		return res.GeneratedKeys[0], nil
	}
	return "", nil
}

func (d *DefaultStore) Update(v m.Model, id string) error {
	v.UpdateTime()
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

func (d *DefaultStore) FindByID(data m.Model, id string) error {

	res, err := r.Table(d.TableName).Get(id).Run(sess)
	if err != nil {
		return err
	}

	if res.IsNil() {
		return ErrNotFound
	}

	return res.One(data)
}
