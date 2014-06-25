package model

type PersonProfile struct {
	*Person
	StudentProfile *StudentProfile `json:"studentProfile,omitempty"`
	TeacherProfile *TeacherProfile `json:"teacherProfile,omitempty"`
}

func (p *PersonProfile) Validate() bool {
	return p.Person.Validate() && p.StudentProfile.Validate()
}
