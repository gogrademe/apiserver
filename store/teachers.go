package store

import ()

type TeacherStore struct {
	DefaultStore
}

func NewTeacherStore() TeacherStore {
	return TeacherStore{DefaultStore: NewDefaultStore("teachers")}
}
