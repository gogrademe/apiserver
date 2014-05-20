package models

type StudentProfile struct {
	PersonId   int64  `db:"person_id"`
	GradeLevel string `db:"grade_level"`
	TimeStamp
}

func CreateStudentProfile(s *StudentProfile) (*StudentProfile, error) {
	s.UpdateTime()
	_, err := db.Exec(`INSERT INTO student_profile(person_id,grade_level, updated_at, created_at)
    VALUES($1,$2,$3,$4)`, s.PersonId, s.GradeLevel, s.UpdatedAt, s.CreatedAt)

	if err != nil {
		return nil, err
	}

	return s, nil
}
