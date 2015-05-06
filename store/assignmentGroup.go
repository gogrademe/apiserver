package store

type AssignmentGroupStore struct {
	DefaultStore
}

func NewAssignmentGroupStore() AssignmentGroupStore {
	return AssignmentGroupStore{DefaultStore: NewDefaultStore("assignmentGroups")}
}
