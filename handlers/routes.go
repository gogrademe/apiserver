package handlers

import (
	"github.com/gin-gonic/gin"
)

// SetupHandlers loads all routes into gorillaMux.
func SetupHandlers(r *gin.Engine) {

	// r.StrictSlash(true)
	m := r.Group("/api")

	// Auth
	m.POST("/session", Login)

	// Users
	m.GET("/user", GetAllUsers)

	// Classes
	m.GET("/class", GetAllClasses)
	m.POST("/class", CreateClass)
	m.GET("/class/:id", GetClass)
	m.PUT("/class/:id", UpdateClass)

	// ClassTerms
	m.GET("term", GetAllClassTerms)
	m.POST("term", CreateClassTerm)
	m.GET("term/:id", GetClassTerm)
	m.PUT("term/:id", UpdateClassTerm)

	// Assignments
	m.GET("assignment", GetAllAssignments)
	m.POST("assignment", CreateAssignment)
	m.GET("assignment/:id", GetAssignment)
	m.PUT("assignment/:id", UpdateAssignment)

	// AssignmentGrades
	m.GET("grade", GetAllAssignmentGrades)
	m.POST("grade", CreateAssignmentGrade)
	m.GET("grade/:id", GetAssignmentGrade)
	m.PUT("grade/:id", UpdateAssignmentGrade)

	// People
	m.GET("/person", GetAllPeople)
	m.POST("/person", CreatePerson)
	m.GET("/person/:id", GetPerson)
	m.PUT("/person/:id", UpdatePerson)

	// Students
	m.GET("/student", GetAllStudents)
	m.POST("/student", CreateStudent)
	m.GET("/student/:id", GetStudent)
	m.PUT("/student/:id", UpdateStudent)

	// Teachers
	m.GET("/teacher", GetAllTeachers)
	m.POST("/teacher", CreateTeacher)
	m.GET("/teacher/:id", GetTeacher)
	m.PUT("/teacher/:id", UpdateTeacher)
}
