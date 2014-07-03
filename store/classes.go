package store

type ClassStore struct {
	DefaultStore
}

func NewClassStore() ClassStore {
	return ClassStore{DefaultStore: NewDefaultStore("classes")}
}
