package store

import (
	m "github.com/gogrademe/apiserver/model"
	// r "github.com/dancannon/gorethink"
)

type AssignmentStore struct {
	DefaultStore
}

func NewAssignmentStore() AssignmentStore {
	return AssignmentStore{DefaultStore: NewDefaultStore("assignments")}
}

// BeforeSave ...
func (a *AssignmentStore) BeforeSave(v *m.Assignment) error {

	aType := m.AssignmentType{}
	err := AssignmentTypes.FindByID(&aType, v.TypeID)
	if err != nil {
		return err
	}

	v.Type = aType
	return nil
}

func (a AssignmentStore) Store(v *m.Assignment) (string, error) {
	err := a.BeforeSave(v)
	if err != nil {
		return "", err
	}

	return a.DefaultStore.Store(v)
}
