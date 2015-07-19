package model

import (
	"net/http"
	"strings"

	"github.com/mholt/binding"
)

type Attempt struct {
	Score        string  `gorethink:"score,omitempty" json:"score"`
	GradeAverage float32 `gorethink:"gradeAverage,omitempty" json:"gradeAverage"`
	TimeStamp
}

// AttemptResource is used for in the API to add new attempts.
type AttemptResource struct {
	AssignmentID string `gorethink:"assignmentId,omitempty" json:"assignmentId"`
	PersonID     string `gorethink:"personId,omitempty" json:"personId"`
	Attempt
}

type AssignmentAttempts struct {
	ID             string     `gorethink:"id,omitempty" json:"id"`
	AssignmentID   string     `gorethink:"assignmentId,omitempty" json:"assignmentId"`
	PersonID       string     `gorethink:"personId,omitempty" json:"personId"`
	LatestAttempt  Attempt    `gorethink:"latestAttempt,omitempty" json:"latestAttempt"`
	AttemptHistory []*Attempt `gorethink:"attemptHistory,omitempty" json:"attemptHistory"`
}

type GradebookResource struct {
	Enrollment
	AssignmentAttempts []AssignmentAttempts `gorethink:"assignmentAttempts,omitempty" json:"assignmentAttempts"`
}

// FieldMap ...
func (a *AttemptResource) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{

		&a.AssignmentID: "assignmentId",
		&a.PersonID:     "personId",
		&a.Score:        "score",
	}
}

func (a AttemptResource) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if strings.TrimSpace(a.AssignmentID) == "" {
		errs = append(errs, RequiredErr("assignmentId"))
	}
	if strings.TrimSpace(a.PersonID) == "" {
		errs = append(errs, RequiredErr("personId"))
	}
	// if strings.TrimSpace(a.Grade) == "" {
	// 	errs = append(errs, RequiredErr("grade"))
	// }
	return errs
}
