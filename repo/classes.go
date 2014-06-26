package repo

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type ClassRepo struct {
}

func NewClassRepo() ClassRepo {
	return ClassRepo{}
}

func (sr *ClassRepo) Store(c *m.Class) error {
	res, err := r.Table("classes").Insert(c).RunWrite(sess)
	if err != nil {
		return err
	}
	c.ID = res.GeneratedKeys[0]

	return nil
}
func (sr *ClassRepo) FindAll() ([]*m.Class, error) {
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
