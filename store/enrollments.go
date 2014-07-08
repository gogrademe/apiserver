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

// r.db('dev_go_grade').table('enrollments')
//   .eqJoin('studentId', r.db('dev_go_grade').table('students')).map(
//     {
//     // I want enrollments to not have a key.
//       "enrollment": r.row("left"),
//     "student": r.row("right")
//   })
//   .eqJoin(r.row('student')('personId'), r.db('dev_go_grade').table('people')).map({
//     // I want this to be added to what is above.
//     "enrollment": r.row("left")("enrollment"),
//      "student": r.row("left")("student"),
//     "person": r.row("right")
//   })

// r.db('dev_go_grade').table('enrollments')
//   .eqJoin('studentId', r.db('dev_go_grade').table('students'))
//   .map(function(result) {
//     return result("left").merge({
//       student: result("right")
//     })
//   })
//   .eqJoin(r.row('student')('personId'), r.db('dev_go_grade').table('people'))
//   .map(function(result) {
//     return result("left").merge({
//       person: result("right")
//     })
//   })

// func (d *DefaultStore) Filter(data interface{}, filter interface{}) error {
// 	res, err := r.Table(d.TableName).Filter(filter).Run(sess)
// 	if err != nil {
// 		return err
// 	}
// 	return res.All(data)
// }

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
