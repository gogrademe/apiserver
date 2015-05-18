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

	// Courses
	g := auth.Group("/course")
	g.GET("", Can("Admin", "Teacher"), GetAllCourses)
	g.GET("/:id", Can("Admin", "Teacher"), GetCourse)
	g.POST("", Can("Admin"), CreateCourse)
	g.DELETE("/:id", Can("Admin"), DeleteCourse)

	g.PUT("/:id", Can("Admin"), UpdateCourse)

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
	g.GET("", Can("Admin", "Teacher"), getAllAssignments)
	g.POST("", Can("Admin", "Teacher"), createAssignment)
	g.GET("/:id", Can("Admin", "Teacher"), getAssignment)
	g.PUT("/:id", Can("Admin", "Teacher"), updateAssignment)
	g.DELETE("/:id", Can("Admin", "Teacher"), deleteAssignment)

	// AssignmentGroups
	g = auth.Group("/assignmentGroup")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignmentGroups)
	g.POST("", Can("Admin"), CreateAssignmentGroup)
	g.GET("/:id", Can("Admin", "Teacher"), GetAssignmentGroup)
	g.PUT("/:id", Can("Admin"), UpdateAssignmentGroup)
	g.DELETE("/:id", Can("Admin"), DeleteAssignmentGroup)

	// AssignmentGrades
	g = auth.Group("/grade")
	g.GET("", Can("Admin", "Teacher"), GetAllAssignmentGrades)
	g.POST("", Can("Admin", "Teacher"), CreateAssignmentGrade)
	// g.GET("/:id", Can("Admin", "Teacher"), GetAssignmentGrade)
	// g.PUT("/:id", Can("Admin", "Teacher"), UpdateAssignmentGrade)

	// People
	g = auth.Group("/person")
	g.GET("", Can("Admin", "Teacher"), GetAllPeople)
	g.POST("", Can("Admin"), CreatePerson)
	g.GET("/:id", Can("Admin", "Teacher"), GetPerson)
	g.PUT("/:id", Can("Admin"), UpdatePerson)
	g.DELETE("/:id", Can("Admin"), DeletePerson)

}
