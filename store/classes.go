package store

import (
// m "github.com/Lanciv/GoGradeAPI/model"
// r "github.com/dancannon/gorethink"
)

type ClassStore struct {
	DefaultStore
}

func NewClassStore() ClassStore {
	return ClassStore{DefaultStore: NewDefaultStore("classes")}
}

// func (sr *ClassStore) Store(c *m.Class) error {
// 	res, err := r.Table("classes").Insert(c).RunWrite(sess)
// 	if err != nil {
// 		return err
// 	}
//
// 	if c.ID == "" && len(res.GeneratedKeys) == 1 {
// 		c.ID = res.GeneratedKeys[0]
// 	}
//
// 	return nil
// }

// func (sr *ClassStore) FindAll() ([]*m.Class, error) {
// 	var c []*m.Class
//
// 	res, err := r.Table("classes").Run(sess)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = res.All(&c)
// 	return c, nil
// }
