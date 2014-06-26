package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

// CreateSession will create a session for a user.
func SaveSession(s *m.Session) error {
	s.UpdateTime()

	res, err := r.Table("sessions").Insert(s).RunWrite(sess)
	if err != nil {
		return err
	}

	s.ID = res.GeneratedKeys[0]

	return nil
}
