package store

type AssignmentStore struct {
	DefaultStore
}

func NewAssignmentStore() AssignmentStore {
	return AssignmentStore{DefaultStore: NewDefaultStore("assignments")}
}
