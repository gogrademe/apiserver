package model

type Student struct {
	PersonID   string `gorethink:"personID"json:"personId"`
	GradeLevel string `gorethink:"gradeLevel"json:"gradeLevel"`
	TimeStamp
}

func (sP *Student) Validate() bool {

	// TODO: Fix
	if sP.PersonID == "" {
		return false
	}
	if sP.GradeLevel == "" {
		return false
	}
	sP.UpdateTime()
	return true
}

// func CreateStudentProfile(s *Student) (*Student, error) {
// 	s.UpdateTime()
// 	_, err := db.Exec(`INSERT INTO student_profile(person_id,grade_level, updated_at, created_at)
//     VALUES($1,$2,$3,$4)`, s.PersonId, s.GradeLevel, s.UpdatedAt, s.CreatedAt)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return s, nil
// }
