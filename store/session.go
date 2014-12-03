package store

// m "github.com/gogrademe/apiserver/model"
// r "github.com/dancannon/gorethink"

type SessionStore struct {
	DefaultStore
}

func NewSessionStore() SessionStore {
	return SessionStore{DefaultStore: NewDefaultStore("sessions")}
}
