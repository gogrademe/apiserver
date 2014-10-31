package store

// m "github.com/GoGradeMe/APIServer/model"
// r "github.com/dancannon/gorethink"

type SessionStore struct {
	DefaultStore
}

func NewSessionStore() SessionStore {
	return SessionStore{DefaultStore: NewDefaultStore("sessions")}
}
