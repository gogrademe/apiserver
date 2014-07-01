package handlers

import (
	"github.com/gin-gonic/gin"
)

// SetupHandlers loads all routes
func SetupHandlers(r *gin.Engine) {
	m := r.Group("/api")

	// Auth
	m.POST("/session", Login)

	auth := m.Group("/", AuthRequired())

	// Users
	auth.GET("/user", GetAllUsers)

	// Classes
	g := auth.Group("/class")
	g.GET("/", GetAllClasses)
	g.POST("/", CreateClass)
	g.GET("/:id", GetClass)
	g.PUT("/:id", UpdateClass)

	// ClassTerms
	g = auth.Group("/term")
	g.GET("/", GetAllClassTerms)
	g.POST("/", CreateClassTerm)
	g.GET("/:id", GetClassTerm)
	g.PUT("/:id", UpdateClassTerm)

	// Assignments
	g = auth.Group("/assignment")
	g.GET("/", GetAllAssignments)
	g.POST("/", CreateAssignment)
	g.GET("/:id", GetAssignment)
	g.PUT("/:id", UpdateAssignment)

	// AssignmentGrades
	g = auth.Group("/grade")
	g.GET("/", GetAllAssignmentGrades)
	g.POST("/", CreateAssignmentGrade)
	g.GET("/:id", GetAssignmentGrade)
	g.PUT("/:id", UpdateAssignmentGrade)

	// People
	g = auth.Group("/person")
	g.GET("/", GetAllPeople)
	g.POST("/", CreatePerson)
	g.GET("/:id", GetPerson)
	g.PUT("/:id", UpdatePerson)

	// Students
	g = auth.Group("/student")
	g.GET("/", GetAllStudents)
	g.POST("/", CreateStudent)
	g.GET("/:id", GetStudent)
	g.PUT("/:id", UpdateStudent)

	// Teachers
	g = auth.Group("/teacher")
	g.GET("/", GetAllTeachers)
	g.POST("/", CreateTeacher)
	g.GET("/:id", GetTeacher)
	g.PUT("/:id", UpdateTeacher)

}
