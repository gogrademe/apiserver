package model

type StudentProfile struct {
	PersonID   int64  `db:"person_id"`
	GradeLevel string `db:"grade_level"`
	TimeStamp
}

func (sP *StudentProfile) Validate() bool {

	// TODO: Fix
	if sP.PersonID == 0 {
		return false
	}
	if sP.GradeLevel == "" {
		return false
	}
	sP.UpdateTime()
	return true
}

// func CreateStudentProfile(s *StudentProfile) (*StudentProfile, error) {
// 	s.UpdateTime()
// 	_, err := db.Exec(`INSERT INTO student_profile(person_id,grade_level, updated_at, created_at)
//     VALUES($1,$2,$3,$4)`, s.PersonId, s.GradeLevel, s.UpdatedAt, s.CreatedAt)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return s, nil
// }
