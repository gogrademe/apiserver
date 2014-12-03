package database

import (
	"errors"

	r "github.com/dancannon/gorethink"
	"github.com/gogrademe/apiserver/model"
)

type Gradestore struct {
	Table r.Term
	r.Session
}

func NewGradestore(s r.Session) *Gradestore {
	return &Gradestore{r.Table("grades"), s}
}

func (db *Gradestore) PostGrade(grade *model.AssignmentGrade) error {
	res, err := db.Table.Insert(grade).RunWrite(&db.Session)
	if err != nil {
		return err
	}

	if len(res.GeneratedKeys) != 1 {
		return errors.New("No keys")
	}

	grade.ID = res.GeneratedKeys[0]

	return nil
}
