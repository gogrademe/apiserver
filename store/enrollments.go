package store

import (
	r "github.com/dancannon/gorethink"
	m "github.com/gogrademe/apiserver/model"
)

type EnrollmentStore struct {
	DefaultStore
}

func NewEnrollmentStore() EnrollmentStore {
	return EnrollmentStore{DefaultStore: NewDefaultStore("enrollments")}
}

func (pr EnrollmentStore) Filter(enrollments *[]m.EnrollmentAPIRes, filter interface{}) error {
	q := r.Table("enrollments").Filter(filter)
	q = q.EqJoin("personId", r.Table("people"))
	// Join Students.
	q = q.Map(func(row r.Term) r.Term {
		return row.Field("left").Merge(map[string]interface{}{
			"person": row.Field("right"),
		})
	})

	first := r.Asc(r.Row.Field("person").Field("firstName"))
	middle := r.Asc(r.Row.Field("person").Field("middleName"))
	last := r.Asc(r.Row.Field("person").Field("lastName"))
	q = q.OrderBy(first, middle, last)

	res, err := q.Run(sess)
	if err != nil {
		return err
	}

	err = res.All(enrollments)
	if err != nil {
		return err
	}

	return nil
}
