package store

import (
	"time"

	r "github.com/dancannon/gorethink"
	m "github.com/gogrademe/apiserver/model"
)

var (
	term1 m.Term
	term2 m.Term
	term3 m.Term
	term4 m.Term
	term5 m.Term
	term6 m.Term

	type1 m.AssignmentType
	type2 m.AssignmentType
	type3 m.AssignmentType
	type4 m.AssignmentType
	type5 m.AssignmentType

	class1 m.Class
	class2 m.Class
	class3 m.Class
	class4 m.Class
	class5 m.Class
	class6 m.Class
	class7 m.Class

	person1  m.Person
	person2  m.Person
	person3  m.Person
	person4  m.Person
	person5  m.Person
	person6  m.Person
	person7  m.Person
	person8  m.Person
	person9  m.Person
	person10 m.Person

	assignment1  m.Assignment
	assignment2  m.Assignment
	assignment3  m.Assignment
	assignment4  m.Assignment
	assignment5  m.Assignment
	assignment6  m.Assignment
	assignment7  m.Assignment
	assignment8  m.Assignment
	assignment9  m.Assignment
	assignment10 m.Assignment
	assignment11 m.Assignment
)

func createDatabase() {

	r.DbCreate(dbName).RunWrite(sess)
}

func createTables() {
	for _, name := range tables {
		r.Db(dbName).TableCreate(name).Run(sess)
	}
}

func cleanTables() {
	for _, name := range tables {
		r.Table(name).Delete().Run(sess)
	}
}

func createIndexes() {

	r.Db(dbName).Table("users").IndexCreate("email").Run(sess)
	r.Db(dbName).Table("users").IndexCreate("emailLower").Run(sess)
	r.Db(dbName).Table("users").IndexCreate("personId").Run(sess)

	r.Db(dbName).Table("assignments").IndexCreate("classId").Run(sess)
	r.Db(dbName).Table("assignments").IndexCreate("termId").Run(sess)
	r.Db(dbName).Table("assignments").IndexCreate("typeId").Run(sess)

	r.Db(dbName).Table("grades").IndexCreate("assignmentId").Run(sess)
	r.Db(dbName).Table("grades").IndexCreate("personId").Run(sess)

	r.Db(dbName).Table("enrollments").IndexCreate("personId").Run(sess)
	r.Db(dbName).Table("enrollments").IndexCreate("classId").Run(sess)
	r.Db(dbName).Table("enrollments").IndexCreate("termId").Run(sess)

	r.Db(dbName).Table("person").IndexCreate("firstName").Run(sess)
	r.Db(dbName).Table("person").IndexCreate("middleName").Run(sess)
	r.Db(dbName).Table("person").IndexCreate("lastName").Run(sess)

	r.Db(dbName).Table("emailConfirmations").IndexCreate("userId").Run(sess)

}

func insertTestData() {

	insertTestTerms()
	insertTestTypes()

	insertTestPeople()

	insertTestUsers()
	insertTestClasses()

	insertTestAssignments()

	insertTestEnrollments()

	insertTestGrades()

}

func insertTestUsers() {

	u, _ := m.NewUserForWithPassword("test@test.com", "somePassword", person7.ID)
	u.Disabled = false
	Users.Store(u)
	// Users.Store(u2)
}

func insertTestTerms() {
	term1 = m.Term{
		Name: "Term 1",
		SchoolYear: m.SchoolYear{
			Start: 2014,
			End:   2015,
		},
		StartDate: time.Date(2014, time.September, 10, 8, 0, 0, 0, time.UTC),
	}
	term2 = m.Term{
		Name: "Term 2",
		SchoolYear: m.SchoolYear{
			Start: 2014,
			End:   2015,
		},
		StartDate: time.Date(2015, time.January, 10, 8, 0, 0, 0, time.UTC),
	}
	term3 = m.Term{
		Name: "Term 3",
		SchoolYear: m.SchoolYear{
			Start: 2014,
			End:   2015,
		},
		StartDate: time.Date(2015, time.May, 10, 8, 0, 0, 0, time.UTC),
	}
	term4 = m.Term{
		Name: "Term 1",
		SchoolYear: m.SchoolYear{
			Start: 2012,
			End:   2013,
		},
		StartDate: time.Date(2012, time.September, 10, 8, 0, 0, 0, time.UTC),
	}
	term5 = m.Term{
		Name: "Term 2",
		SchoolYear: m.SchoolYear{
			Start: 2012,
			End:   2013,
		},
		StartDate: time.Date(2013, time.January, 10, 8, 0, 0, 0, time.UTC),
	}
	term6 = m.Term{
		Name: "Term 3",
		SchoolYear: m.SchoolYear{
			Start: 2012,
			End:   2013,
		},
		StartDate: time.Date(2013, time.May, 10, 8, 0, 0, 0, time.UTC),
	}

	keys, _ := Terms.Insert(&term1, &term2, &term3, &term4, &term5, &term6)
	term1.ID = keys[0]
	term2.ID = keys[1]
	term3.ID = keys[2]
	term4.ID = keys[3]
	term5.ID = keys[4]
	term6.ID = keys[5]
}

func insertTestClasses() {

	class1 = m.Class{
		Name:       "Math",
		GradeLevel: "Year 7",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}
	class2 = m.Class{
		Name:       "Math",
		GradeLevel: "Year 8",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}
	class3 = m.Class{
		Name:       "Science",
		GradeLevel: "Year 10",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}
	class4 = m.Class{
		Name:       "Science",
		GradeLevel: "Year 7",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}
	class5 = m.Class{
		Name:       "Art",
		GradeLevel: "Year 7",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}
	class6 = m.Class{
		Name:       "Art",
		GradeLevel: "Year 8",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}
	class7 = m.Class{
		Name:       "Art",
		GradeLevel: "Year 10",
		Terms: []string{
			term1.ID,
			term2.ID,
			term3.ID,
		},
	}

	keys, _ := Classes.Insert(&class1, &class2, &class3, &class4, &class5, &class6, &class7)
	class1.ID = keys[0]
	class2.ID = keys[1]
	class3.ID = keys[2]
	class4.ID = keys[3]
	class5.ID = keys[4]
	class6.ID = keys[5]
	class7.ID = keys[6]

}
func insertTestTypes() {
	type1 = m.AssignmentType{
		Name:   "Class Test",
		Weight: .20,
	}
	type2 = m.AssignmentType{
		Name:   "Written Work",
		Weight: .40,
	}
	type3 = m.AssignmentType{
		Name:   "Quiz",
		Weight: .20,
	}
	type4 = m.AssignmentType{
		Name:   "Project",
		Weight: .20,
	}
	type5 = m.AssignmentType{
		Name:   "Mid Term",
		Weight: .60,
	}

	type1.ID, _ = AssignmentTypes.Store(&type1)
	type2.ID, _ = AssignmentTypes.Store(&type2)
	type3.ID, _ = AssignmentTypes.Store(&type3)
	type4.ID, _ = AssignmentTypes.Store(&type4)
	type5.ID, _ = AssignmentTypes.Store(&type5)

}

func insertTestAssignments() {
	assignment1 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type1.ID,
		MaxScore: 100,
		Name:     "Notebook Check",
	}
	assignment2 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type1.ID,
		MaxScore: 100,
		Name:     "Formula Quiz",
	}
	assignment3 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type1.ID,
		MaxScore: 100,
		Name:     "Solar Model",
	}
	assignment4 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type1.ID,
		MaxScore: 100,
		Name:     "Cell Model",
	}
	assignment5 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type2.ID,
		MaxScore: 100,
		Name:     "Test 1",
	}
	assignment6 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type3.ID,
		MaxScore: 100,
		Name:     "Formula Quiz",
	}
	assignment7 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type4.ID,
		MaxScore: 100,
		Name:     "Solar Model",
	}
	assignment8 = m.Assignment{
		ClassID:  class1.ID,
		TermID:   term1.ID,
		TypeID:   type5.ID,
		MaxScore: 100,
		Name:     "Cell Model",
	}
	assignment9 = m.Assignment{
		ClassID:  class2.ID,
		TermID:   term1.ID,
		MaxScore: 100,
		TypeID:   type5.ID,
		Name:     "Cell Model",
	}
	assignment10 = m.Assignment{
		ClassID:  class2.ID,
		TermID:   term1.ID,
		TypeID:   type5.ID,
		MaxScore: 100,
		Name:     "Cell Model",
	}
	assignment11 = m.Assignment{
		ClassID:  class2.ID,
		TermID:   term2.ID,
		TypeID:   type5.ID,
		MaxScore: 100,
		Name:     "Cell Model",
	}

	assignment1.ID, _ = Assignments.Store(&assignment1)
	Assignments.Store(&assignment2)
	Assignments.Store(&assignment3)
	Assignments.Store(&assignment4)
	Assignments.Store(&assignment5)
	Assignments.Store(&assignment6)
	Assignments.Store(&assignment7)
	Assignments.Store(&assignment8)
	Assignments.Store(&assignment9)
	Assignments.Store(&assignment10)
	Assignments.Store(&assignment11)
}
func insertTestPeople() {

	person1 = m.Person{
		FirstName:  "Jon",
		MiddleName: "David",
		LastName:   "Bush",
		GradeLevel: "Year 1",
		Types:      []string{"Student"},
	}
	person2 = m.Person{
		FirstName:  "Angel",
		LastName:   "Heredia",
		GradeLevel: "Year 9",
		Types:      []string{"Student"},
	}
	person3 = m.Person{
		FirstName:  "Nicole",
		LastName:   "Aitchison",
		GradeLevel: "Year 7",
		Types:      []string{"Student"},
	}
	person4 = m.Person{
		FirstName:  "Frankie",
		LastName:   "Bagnardi",
		GradeLevel: "Year 9",
		Types:      []string{"Student"},
	}
	person5 = m.Person{
		FirstName:  "Adam",
		LastName:   "Price",
		GradeLevel: "Year 4",
		Types:      []string{"Student"},
	}
	person6 = m.Person{
		FirstName:  "Jake",
		MiddleName: "Matthew",
		LastName:   "Price",
		GradeLevel: "Year 9",
		Types:      []string{"Student"},
	}
	person7 = m.Person{
		FirstName:  "Matthew",
		LastName:   "Aitchison",
		GradeLevel: "Year 12",
		Types:      []string{"Admin"},
	}
	person8 = m.Person{
		FirstName:  "Natalie",
		LastName:   "Aitchison",
		GradeLevel: "Year 3",
		Types:      []string{"Student"},
	}
	person9 = m.Person{
		FirstName: "Susan",
		LastName:  "Feathers",
		Email:     "Susan.Feathers@test.com",
		Types:     []string{"Teacher"},
	}
	person10 = m.Person{
		FirstName: "Karen",
		LastName:  "Portman",
		Email:     "Karen.Portman@test.com",
		Types:     []string{"Teacher"},
	}

	keys, _ := People.Insert(&person1, &person2, &person3, &person4, &person5,
		&person6, &person7, &person8, &person9, &person10)

	person1.ID = keys[0]
	person2.ID = keys[1]
	person3.ID = keys[2]
	person4.ID = keys[3]
	person5.ID = keys[4]
	person6.ID = keys[5]
	person7.ID = keys[6]
	person8.ID = keys[7]
	person9.ID = keys[8]
	person10.ID = keys[9]

}

func insertTestEnrollments() {
	p1 := m.Enrollment{
		PersonID: person1.ID,
		ClassID:  class1.ID,
		TermID:   term1.ID,
	}
	p2 := m.Enrollment{
		PersonID: person2.ID,
		ClassID:  class1.ID,
		TermID:   term1.ID,
	}
	p3 := m.Enrollment{
		PersonID: person3.ID,
		ClassID:  class1.ID,
		TermID:   term1.ID,
	}
	p4 := m.Enrollment{
		PersonID: person4.ID,
		ClassID:  class1.ID,
		TermID:   term1.ID,
	}
	p5 := m.Enrollment{
		PersonID: person4.ID,
		ClassID:  class2.ID,
		TermID:   term1.ID,
	}
	p6 := m.Enrollment{
		PersonID: person4.ID,
		ClassID:  class3.ID,
		TermID:   term1.ID,
	}
	p7 := m.Enrollment{
		PersonID: person4.ID,
		ClassID:  class5.ID,
		TermID:   term1.ID,
	}
	p8 := m.Enrollment{
		PersonID: person4.ID,
		ClassID:  class6.ID,
		TermID:   term1.ID,
	}
	p9 := m.Enrollment{
		PersonID: person4.ID,
		ClassID:  class7.ID,
		TermID:   term1.ID,
	}
	Enrollments.Store(&p1)
	Enrollments.Store(&p2)
	Enrollments.Store(&p3)
	Enrollments.Store(&p4)
	Enrollments.Store(&p5)
	Enrollments.Store(&p6)
	Enrollments.Store(&p7)
	Enrollments.Store(&p8)
	Enrollments.Store(&p9)
}

func insertTestGrades() {
	grade1 := m.AssignmentGrade{
		AssignmentID: assignment1.ID,
		PersonID:     person1.ID,
		Grade:        "50",
	}
	grade2 := m.AssignmentGrade{
		AssignmentID: assignment1.ID,
		PersonID:     person2.ID,
		Grade:        "50",
	}

	AssignmentGrades.Insert(&grade1, &grade2)
}
