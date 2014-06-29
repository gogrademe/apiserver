package store

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type StoreSuite struct{}

var _ = Suite(&StoreSuite{})

func (s *StoreSuite) SetUpSuite(c *C) {
	if err := Connect(testAddress, testDBName); err != nil {
		c.Fatalf("Failed to connect to db with: %s", err)
	}
	SetupDB(false)
}

func (s *StoreSuite) SetUpTest(c *C) {
	cleanTables()
}
