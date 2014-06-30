package store

type ClassStore struct {
	DefaultStore
}

func NewClassStore() ClassStore {
	return ClassStore{DefaultStore: NewDefaultStore("classes")}
}

type ClassTermStore struct {
	DefaultStore
}

func NewClassTermStore() ClassTermStore {
	return ClassTermStore{DefaultStore: NewDefaultStore("classTerms")}
}
