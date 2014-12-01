package handlers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	m "github.com/GoGradeMe/APIServer/model"
	"github.com/GoGradeMe/APIServer/store"
)

// RoleIn ...
func RoleIn(role string, roles []string) bool {
	for _, r := range roles {
		if role == r {
			return true
		}
	}
	return false
}

// Can ...
func Can(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID, err := c.Get("personId")
		if err != nil && personID == nil {
			c.Fail(401, errors.New("PersonId not found."))
			return
		}

		id := personID.(string)

		person := m.Person{}
		err = store.People.One(&person, id)
		if err != nil {
			c.Fail(401, errors.New("Person not found."))
			return
		}

		for k := range person.Types {
			if RoleIn(person.Types[k], roles) {
				log.Printf("Role %s matched for person: %s", roles, person)
				return
			}
		}

		c.Fail(401, errors.New("Unauthorized"))
		return

	}
}
