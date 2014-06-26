package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type ClassStore struct {
}

func NewClassStore() ClassStore {
	return ClassStore{}
}

func (sr *ClassStore) Store(c *m.Class) error {
	res, err := r.Table("classes").Insert(c).RunWrite(sess)
	if err != nil {
		return err
	}
	c.ID = res.GeneratedKeys[0]

	return nil
}
func (sr *ClassStore) FindAll() ([]*m.Class, error) {
	var c []*m.Class

	rows, err := r.Table("classes").Run(sess)
	if err != nil {
		return nil, err
	}

	err = rows.ScanAll(&c)
	return c, nil
}

// func CreateClass(c *m.Class) error {
// 	res, err := r.Table("classes").Insert(c).RunWrite(sess)
// 	if err != nil {
// 		return err
// 	}
// 	c.ID = res.GeneratedKeys[0]
//
// 	return nil
// }

// func GetAllClasses() ([]m.Class, error) {
//
// 	classes := []m.Class{}
//
// 	rows, err := r.Table("classes").Run(sess)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = rows.ScanAll(&classes)
// 	return classes, nil
// }
