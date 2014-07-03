package store

type TermStore struct {
	DefaultStore
}

func NewTermStore() TermStore {
	return TermStore{DefaultStore: NewDefaultStore("terms")}
}
