package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

type SessionStore struct {
}

func NewSessionStore() SessionStore {
	return SessionStore{}
}

// CreateSession will create a session for a user.
func (ss *SessionStore) Store(s *m.Session) error {
	s.UpdateTime()

	res, err := r.Table("sessions").Insert(s).RunWrite(sess)
	if err != nil {
		return err
	}

	s.ID = res.GeneratedKeys[0]

	return nil
}
