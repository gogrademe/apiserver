package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type EnrollmentStore struct {
	DefaultStore
}

func NewEnrollmentStore() EnrollmentStore {
	return EnrollmentStore{DefaultStore: NewDefaultStore("enrollments")}
}

func (pr EnrollmentStore) Filter(enrollments *[]m.EnrollmentAPIRes, filter interface{}) error {
	q := r.Table("enrollments").Filter(filter)
	q = q.EqJoin("studentId", r.Table("students"))
	// Join Students.
	q = q.Map(func(row r.Term) r.Term {
		return row.Field("left").Merge(map[string]interface{}{
			"student": row.Field("right"),
		})
	})
	// Join Person.
	q = q.EqJoin(r.Row.Field("student").Field("personId"), r.Table("people"))
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
