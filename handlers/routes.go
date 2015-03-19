package handlers

import (
	"github.com/gin-gonic/gin"
)

// SetupHandlers loads all routes
func SetupHandlers(r *gin.Engine) {
	r.Use(CORSMiddleware())

	r.OPTIONS("/*path", CORSMiddleware())

	// Auth
	r.POST("/session", Login)

	auth := r.Group("", AuthRequired())

	// Users
	auth.GET("/user", Can("Admin"), GetAllUsers)
	auth.POST("/user", Can("Admin"), CreateUser)

	// Classes
	g := auth.Group("/class")
	g.GET("", Can("Admin", "Teacher"), GetAllClasses)
	g.GET("/:id", Can("Admin", "Teacher"), GetClass)
	g.POST("", Can("Admin"), CreateClass)
	g.DELETE("/:id", Can("Admin"), DeleteClass)

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

	// SchoolYears
	g = auth.Group("/schoolYear")
	g.GET("", Can("Admin", "Teacher"), GetAllSchoolYears)
	g.POST("", Can("Admin"), CreateSchoolYear)
	g.GET("/:id", Can("Admin", "Teacher"), GetSchoolYear)
	g.PUT("/:id", Can("Admin"), UpdateSchoolYear)

	// Assignments
	g = auth.Group("/assignment")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignments)
	g.POST("", Can("Admin", "Teacher"), CreateAssignment)
	g.GET("/:id", Can("Admin", "Teacher"), GetAssignment)
	g.PUT("/:id", Can("Admin", "Teacher"), UpdateAssignment)
	g.DELETE("/:id", Can("Admin", "Teacher"), DeleteAssignment)

	// AssignmentTypes
	g = auth.Group("/type")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignmentTypes)
	g.POST("", Can("Admin"), CreateAssignmentType)
	g.GET("/:id", Can("Admin", "Teacher"), GetAssignmentType)
	g.PUT("/:id", Can("Admin"), UpdateAssignmentType)
	g.DELETE("/:id", Can("Admin"), DeleteAssignmentType)

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

}
