package store

// m "github.com/Lanciv/GoGradeAPI/model"
// r "github.com/dancannon/gorethink"

type SessionStore struct {
	DefaultStore
}

func NewSessionStore() SessionStore {
	return SessionStore{DefaultStore: NewDefaultStore("sessions")}
}
