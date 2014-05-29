package model

type Class struct {
	Id         int64
	Name       string
	TeacherId  int64  `db:"teacher_id"`
	GradeLevel string `db:"grade_level"`
	Subject    string
	TimeStamp
}

// func (c *Class) Validate() bool {
// 	return false
// }

// func GetAllClasses() ([]Class, error) {

// 	// classes := []Class{}
// 	c := Class{}
// 	err := db.Get(&classes, "SELECT * FROM class limit 1")

// 	if err != nil {
// 		return nil, err
// 	}

// 	// classes := []Class{}
// 	// db.Find(&classes)

// 	return classes, nil
// }
