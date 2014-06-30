package store

import ()

type ParentStore struct {
	DefaultStore
}

func NewParentStore() ParentStore {
	return ParentStore{DefaultStore: NewDefaultStore("parents")}
}
