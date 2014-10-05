package store

import (
	"time"

	m "github.com/Lanciv/GoGradeAPI/model"
	r "github.com/dancannon/gorethink"

	//"github.com/Pallinder/go-randomdata"
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

	student1 m.Student
	student2 m.Student
	student3 m.Student
	student4 m.Student
	student5 m.Student
	student6 m.Student
	student7 m.Student
	student8 m.Student

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
	r.Db(dbName).Table("grades").IndexCreate("studentId").Run(sess)

	r.Db(dbName).Table("enrollments").IndexCreate("studentId").Run(sess)
	r.Db(dbName).Table("enrollments").IndexCreate("classId").Run(sess)
	r.Db(dbName).Table("enrollments").IndexCreate("termId").Run(sess)

	r.Db(dbName).Table("person").IndexCreate("firstName").Run(sess)
	r.Db(dbName).Table("person").IndexCreate("middleName").Run(sess)
	r.Db(dbName).Table("person").IndexCreate("lastName").Run(sess)

	r.Db(dbName).Table("students").IndexCreate("personId").Run(sess)

	r.Db(dbName).Table("teachers").IndexCreate("personId").Run(sess)

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

	u, _ := m.NewUserFor("test@test.com", "somePassword", "Admin", person7.ID)
	u2, _ := m.NewUserFor("Susan.Feathers@test.com", "somePassword", "Teacher", person9.ID)
	Users.Store(u)
	Users.Store(u2)
}

func insertTestTerms() {
	term1 = m.Term{
		Name:       "Term 1",
		SchoolYear: "2014-2015",
		StartDate:  time.Date(2014, time.September, 10, 8, 0, 0, 0, time.UTC),
	}
	term2 = m.Term{
		Name:       "Term 2",
		SchoolYear: "2014-2015",
		StartDate:  time.Date(2015, time.January, 10, 8, 0, 0, 0, time.UTC),
	}
	term3 = m.Term{
		Name:       "Term 3",
		SchoolYear: "2014-2015",
		StartDate:  time.Date(2015, time.May, 10, 8, 0, 0, 0, time.UTC),
	}
	term4 = m.Term{
		Name:       "Term 1",
		SchoolYear: "2012-2013",
		StartDate:  time.Date(2012, time.September, 10, 8, 0, 0, 0, time.UTC),
	}
	term5 = m.Term{
		Name:       "Term 2",
		SchoolYear: "2012-2013",
		StartDate:  time.Date(2013, time.January, 10, 8, 0, 0, 0, time.UTC),
	}
	term6 = m.Term{
		Name:       "Term 3",
		SchoolYear: "2012-2013",
		StartDate:  time.Date(2013, time.May, 10, 8, 0, 0, 0, time.UTC),
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
	assignment1 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Notebook Check",
	}
	assignment2 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Formula Quiz",
	}
	assignment3 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Solar Model",
	}
	assignment4 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type1.ID,
		Name:    "Cell Model",
	}
	assignment5 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type2.ID,
		Name:    "Test 1",
	}
	assignment6 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type3.ID,
		Name:    "Formula Quiz",
	}
	assignment7 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type4.ID,
		Name:    "Solar Model",
	}
	assignment8 = m.Assignment{
		ClassID: class1.ID,
		TermID:  term1.ID,
		TypeID:  type5.ID,
		Name:    "Cell Model",
	}
	assignment9 = m.Assignment{
		ClassID: class2.ID,
		TermID:  term1.ID,
		TypeID:  type5.ID,
		Name:    "Cell Model",
	}
	assignment10 = m.Assignment{
		ClassID: class2.ID,
		TermID:  term1.ID,
		TypeID:  type5.ID,
		Name:    "Cell Model",
	}
	assignment11 = m.Assignment{
		ClassID: class2.ID,
		TermID:  term2.ID,
		TypeID:  type5.ID,
		Name:    "Cell Model",
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
	}
	person2 = m.Person{
		FirstName: "Angel",
		LastName:  "Heredia",
	}
	person3 = m.Person{
		FirstName: "Nicole",
		LastName:  "Aitchison",
	}
	person4 = m.Person{
		FirstName: "Frankie",
		LastName:  "Bagnardi",
	}
	person5 = m.Person{
		FirstName: "Adam",
		LastName:  "Price",
	}
	person6 = m.Person{
		FirstName:  "Jake",
		MiddleName: "Matthew",
		LastName:   "Price",
	}
	person7 = m.Person{
		FirstName: "Matthew",
		LastName:  "Aitchison",
	}
	person8 = m.Person{
		FirstName: "Natalie",
		LastName:  "Aitchison",
	}
	person9 = m.Person{

		FirstName: "Susan",
		LastName:  "Feathers",
	}
	person10 = m.Person{

		FirstName: "Karen",
		LastName:  "Portman",
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

	student1 = m.Student{
		PersonID:   person1.ID,
		GradeLevel: "Year 1",
	}
	student2 = m.Student{
		PersonID:   person2.ID,
		GradeLevel: "Year 9",
	}
	student3 = m.Student{
		PersonID:   person3.ID,
		GradeLevel: "Year 12",
	}
	student4 = m.Student{
		PersonID:   person4.ID,
		GradeLevel: "Year 9",
	}
	student5 = m.Student{
		PersonID:   person5.ID,
		GradeLevel: "Year 12",
	}
	student6 = m.Student{
		PersonID:   person6.ID,
		GradeLevel: "Year 12",
	}
	student7 = m.Student{
		PersonID:   person7.ID,
		GradeLevel: "Year 12",
	}
	student8 = m.Student{
		PersonID:   person8.ID,
		GradeLevel: "Year 12",
	}

	student1.ID, _ = Students.Store(&student1)
	student2.ID, _ = Students.Store(&student2)
	student3.ID, _ = Students.Store(&student3)
	student4.ID, _ = Students.Store(&student4)
	student5.ID, _ = Students.Store(&student5)
	student6.ID, _ = Students.Store(&student6)
	student7.ID, _ = Students.Store(&student7)
	student8.ID, _ = Students.Store(&student8)

	t1 := m.Teacher{
		PersonID: person9.ID,
		Email:    "Susan.Feathers@test.com",
	}
	t2 := m.Teacher{
		PersonID: person10.ID,
		Email:    "Karen.Portman@test.com",
	}

	Teachers.Store(&t1)
	Teachers.Store(&t2)

}

func insertTestEnrollments() {
	p1 := m.Enrollment{
		StudentID: student1.ID,
		ClassID:   class1.ID,
		TermID:    term1.ID,
	}
	p2 := m.Enrollment{
		StudentID: student2.ID,
		ClassID:   class1.ID,
		TermID:    term1.ID,
	}
	p3 := m.Enrollment{
		StudentID: student3.ID,
		ClassID:   class1.ID,
		TermID:    term1.ID,
	}
	p4 := m.Enrollment{
		StudentID: student4.ID,
		ClassID:   class1.ID,
		TermID:    term1.ID,
	}
	p5 := m.Enrollment{
		StudentID: student4.ID,
		ClassID:   class2.ID,
		TermID:    term1.ID,
	}
	p6 := m.Enrollment{
		StudentID: student4.ID,
		ClassID:   class3.ID,
		TermID:    term1.ID,
	}
	p7 := m.Enrollment{
		StudentID: student4.ID,
		ClassID:   class5.ID,
		TermID:    term1.ID,
	}
	p8 := m.Enrollment{
		StudentID: student4.ID,
		ClassID:   class6.ID,
		TermID:    term1.ID,
	}
	p9 := m.Enrollment{
		StudentID: student4.ID,
		ClassID:   class7.ID,
		TermID:    term1.ID,
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
		StudentID:    student1.ID,
		Grade:        "50",
	}
	grade2 := m.AssignmentGrade{
		AssignmentID: assignment1.ID,
		StudentID:    student2.ID,
		Grade:        "50",
	}

	AssignmentGrades.Insert(&grade1, &grade2)
}
