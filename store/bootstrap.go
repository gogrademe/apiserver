package store

import (
	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"
)

var (
	term1 m.Term
	term2 m.Term
	term3 m.Term

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

}

func insertTestData() {

	insertTestUsers()

	insertTestTerms()
	insertTestTypes()

	insertTestPeople()
	insertTestClasses()

	insertTestAssignments()

}
func insertTestUsers() {

	u, _ := m.NewUser("test@test.com", "somePassword", "Admin")
	Users.Store(u)
}
func insertTestTerms() {
	term1 = m.Term{
		Name: "Term 1",
	}
	term2 = m.Term{
		Name: "Term 2",
	}
	term3 = m.Term{
		Name: "Term 3",
	}

	term1.ID, _ = Terms.Store(&term1)
	term2.ID, _ = Terms.Store(&term2)
	term3.ID, _ = Terms.Store(&term3)
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

	class1.ID, _ = Classes.Store(&class1)
	class2.ID, _ = Classes.Store(&class2)
	class3.ID, _ = Classes.Store(&class3)
	class4.ID, _ = Classes.Store(&class4)
	class5.ID, _ = Classes.Store(&class5)
	class6.ID, _ = Classes.Store(&class6)
	class7.ID, _ = Classes.Store(&class7)

}
func insertTestTypes() {
	type1 = m.AssignmentType{
		Name:     "Class Test",
		Weight:   .20,
		MaxScore: 100,
	}
	type2 = m.AssignmentType{
		Name:     "Written Work",
		Weight:   .40,
		MaxScore: 100,
	}
	type3 = m.AssignmentType{
		Name:     "Quiz",
		Weight:   .20,
		MaxScore: 100,
	}
	type4 = m.AssignmentType{
		Name:     "Project",
		Weight:   .20,
		MaxScore: 100,
	}
	type5 = m.AssignmentType{
		Name:     "Mid Term",
		Weight:   .60,
		MaxScore: 100,
	}

	type1.ID, _ = AssignmentTypes.Store(&type1)
	type2.ID, _ = AssignmentTypes.Store(&type2)
	type3.ID, _ = AssignmentTypes.Store(&type3)
	type4.ID, _ = AssignmentTypes.Store(&type4)
	type5.ID, _ = AssignmentTypes.Store(&type5)

}

func insertTestAssignments() {
	a1 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Notebook Check",
	}
	a2 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Formula Quiz",
	}
	a3 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Solar Model",
	}
	a4 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Cell Model",
	}
	a5 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type2.ID,
		Name:    "Test 1",
	}
	a6 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type3.ID,
		Name:    "Formula Quiz",
	}
	a7 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type4.ID,
		Name:    "Solar Model",
	}
	a8 := m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type5.ID,
		Name:    "Cell Model",
	}

	Assignments.Store(&a1)
	Assignments.Store(&a2)
	Assignments.Store(&a3)
	Assignments.Store(&a4)
	Assignments.Store(&a5)
	Assignments.Store(&a6)
	Assignments.Store(&a7)
	Assignments.Store(&a8)
}
func insertTestPeople() {

	p1 := m.Person{
		FirstName:  "Jon",
		MiddleName: "David",
		LastName:   "Bush",
	}
	p2 := m.Person{
		FirstName: "Angel",
		LastName:  "Heredia",
	}
	p3 := m.Person{
		FirstName: "Nicole",
		LastName:  "Aitchison",
	}
	p4 := m.Person{
		FirstName: "Frankie",
		LastName:  "Bagnardi",
	}
	p5 := m.Person{
		FirstName: "Adam",
		LastName:  "Price",
	}
	p6 := m.Person{
		FirstName:  "Jake",
		MiddleName: "Matthew",
		LastName:   "Price",
	}
	p7 := m.Person{
		FirstName: "Matthew",
		LastName:  "Aitchison",
	}
	p8 := m.Person{
		FirstName: "Natalie",
		LastName:  "Aitchison",
	}
	p9 := m.Person{

		FirstName: "Susan",
		LastName:  "Feathers",
	}
	p10 := m.Person{

		FirstName: "Karen",
		LastName:  "Portman",
	}

	p1.ID, _ = People.Store(&p1)
	p2.ID, _ = People.Store(&p2)
	p3.ID, _ = People.Store(&p3)
	p4.ID, _ = People.Store(&p4)
	p5.ID, _ = People.Store(&p5)
	p6.ID, _ = People.Store(&p6)
	p7.ID, _ = People.Store(&p7)
	p8.ID, _ = People.Store(&p8)
	p9.ID, _ = People.Store(&p9)
	p10.ID, _ = People.Store(&p10)

	s1 := m.Student{
		PersonID:   p1.ID,
		GradeLevel: "Year 1",
	}
	s2 := m.Student{
		PersonID:   p2.ID,
		GradeLevel: "Year 9",
	}
	s3 := m.Student{
		PersonID:   p3.ID,
		GradeLevel: "Year 12",
	}
	s4 := m.Student{
		PersonID:   p4.ID,
		GradeLevel: "Year 9",
	}
	s5 := m.Student{
		PersonID:   p5.ID,
		GradeLevel: "Year 12",
	}
	s6 := m.Student{
		PersonID:   p6.ID,
		GradeLevel: "Year 12",
	}
	s7 := m.Student{
		PersonID:   p7.ID,
		GradeLevel: "Year 12",
	}
	s8 := m.Student{
		PersonID:   p8.ID,
		GradeLevel: "Year 12",
	}

	Students.Store(&s1)
	Students.Store(&s2)
	Students.Store(&s3)
	Students.Store(&s4)
	Students.Store(&s5)
	Students.Store(&s6)
	Students.Store(&s7)
	Students.Store(&s8)

	t1 := m.Teacher{
		PersonID: p9.ID,
		Email:    "Susan.Feathers@test.com",
	}
	t2 := m.Teacher{
		PersonID: p10.ID,
		Email:    "Karen.Portman@test.com",
	}

	Teachers.Store(&t1)
	Teachers.Store(&t2)

}
