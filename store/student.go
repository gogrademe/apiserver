package store

import ()

type StudentStore struct {
	DefaultStore
}

func NewStudentStore() StudentStore {
	return StudentStore{DefaultStore: NewDefaultStore("students")}
}
