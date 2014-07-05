package store

type ClassPersonStore struct {
	DefaultStore
}

func NewClassPersonStore() ClassPersonStore {
	return ClassPersonStore{DefaultStore: NewDefaultStore("classPeople")}
}
