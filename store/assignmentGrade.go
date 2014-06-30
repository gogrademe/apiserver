package store

import ()

type AssignmentGradeStore struct {
	DefaultStore
}

func NewAssignmentGradeStore() AssignmentGradeStore {
	return AssignmentGradeStore{DefaultStore: NewDefaultStore("assignments")}
}
