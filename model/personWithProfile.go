package model

type PersonProfile struct {
	Person         *Person         `json:"person"`
	StudentProfile *StudentProfile `json:"studentProfile,omitempty"`
}

func (p *PersonProfile) Validate() bool {
	return p.Person.Validate() && p.StudentProfile.Validate()
}
