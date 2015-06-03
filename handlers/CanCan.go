package handlers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	m "github.com/gogrademe/apiserver/model"
	"github.com/gogrademe/apiserver/store"
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
		personID, ok := c.Get("personId")
		if !ok && personID == nil {
			c.AbortWithError(401, errors.New("PersonId not found."))
			return
		}

		id := personID.(string)

		person := m.Person{}
		err := store.People.One(&person, id)
		if err != nil {
			c.AbortWithError(401, errors.New("Person not found."))
			return
		}

		for k := range person.Types {
			if RoleIn(person.Types[k], roles) {
				log.Printf("Role %s matched for person: %s", roles, person)
				return
			}
		}

		c.AbortWithError(401, errors.New("Unauthorized"))
		return

	}
}
