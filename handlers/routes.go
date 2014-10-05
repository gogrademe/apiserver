package handlers

import (
	"github.com/gin-gonic/gin"
)

// SetupHandlers loads all routes
func SetupHandlers(r *gin.Engine) {
	// r.Handle("OPTIONS", "/*cors", []gin.HandlerFunc{CORSMiddleware()})
	r.Use(CORSMiddleware())
	m := r.Group("/api")

	// Auth
	m.POST("/session", Login)

	auth := m.Group("", AuthRequired())

	// Users
	auth.GET("/user", Can("Admin", "Teacher"), GetAllUsers)
	auth.POST("/user", Can("Admin"), CreateUser)

	// Classes
	g := auth.Group("/class")
	g.GET("", Can("Admin", "Teacher"), GetAllClasses)
	g.GET("/:id", Can("Admin", "Teacher"), GetClass)
	g.POST("", Can("Admin"), CreateClass)

	g.PUT("/:id", Can("Admin"), UpdateClass)

	// Enrollments
	g = auth.Group("/enrollment")
	g.GET("", Can("Admin", "Teacher"), GetAllEnrollments)
	g.POST("", Can("Admin"), CreateEnrollment)
	g.GET("/:id", Can("Admin", "Teacher"), GetEnrollment)
	g.PUT("/:id", Can("Admin"), UpdateEnrollment)
	g.DELETE("/:id", Can("Admin"), DeleteEnrollment)

	// Terms
	g = auth.Group("/term")
	g.GET("", Can("Admin", "Teacher"), GetAllTerms)
	g.POST("", Can("Admin"), CreateTerm)
	g.GET("/:id", Can("Admin", "Teacher"), GetTerm)
	g.PUT("/:id", Can("Admin"), UpdateTerm)

	// Assignments
	g = auth.Group("/assignment")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignments)
	g.POST("", Can("Admin", "Teacher"), CreateAssignment)
	// g.POST("", Create(store.Assignments))
	g.GET("/:id", Can("Admin", "Teacher"), GetAssignment)
	g.PUT("/:id", Can("Admin", "Teacher"), UpdateAssignment)
	g.DELETE("/:id", Can("Admin", "Teacher"), DeleteAssignment)

	// AssignmentTypes
	g = auth.Group("/type")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignmentTypes)
	g.POST("", Can("Admin", "Teacher"), CreateAssignmentType)
	g.GET("/:id", Can("Admin", "Teacher"), GetAssignmentType)
	g.PUT("/:id", Can("Admin", "Teacher"), UpdateAssignmentType)

	// AssignmentGrades
	g = auth.Group("/grade")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignmentGrades)
	g.POST("", Can("Admin", "Teacher"), CreateAssignmentGrade)
	g.GET("/:id", Can("Admin", "Teacher"), GetAssignmentGrade)
	g.PUT("/:id", Can("Admin", "Teacher"), UpdateAssignmentGrade)

	// People
	g = auth.Group("/person")
	g.GET("", Can("Admin", "Teacher"), GetAllPeople)
	g.POST("", Can("Admin"), CreatePerson)
	g.GET("/:id", Can("Admin", "Teacher"), GetPerson)
	g.PUT("/:id", Can("Admin"), UpdatePerson)
	g.DELETE("/:id", Can("Admin"), DeletePerson)

	// Students
	g = auth.Group("/student")
	g.GET("", Can("Admin", "Teacher"), GetAllStudents)
	g.POST("", Can("Admin"), CreateStudent)
	g.GET("/:id", Can("Admin", "Teacher"), GetStudent)
	g.PUT("/:id", Can("Admin"), UpdateStudent)

	// Teachers
	g = auth.Group("/teacher")
	g.GET("", Can("Admin", "Teacher"), GetAllTeachers)
	g.POST("", Can("Admin"), CreateTeacher)
	g.GET("/:id", Can("Admin", "Teacher"), GetTeacher)
	g.PUT("/:id", Can("Admin"), UpdateTeacher)

}
