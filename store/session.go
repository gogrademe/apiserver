package store

type SessionStore struct {
	DefaultStore
}

func NewSessionStore() SessionStore {
	return SessionStore{DefaultStore: NewDefaultStore("sessions")}
}
