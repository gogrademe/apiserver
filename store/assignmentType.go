package store

import ()

type AssignmentTypeStore struct {
	DefaultStore
}

func NewAssignmentTypeStore() AssignmentTypeStore {
	return AssignmentTypeStore{DefaultStore: NewDefaultStore("assignmentTypes")}
}
