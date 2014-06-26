package database

import (
	"testing"
)

func TestDBCon(t *testing.T) {

	if err := Connect(testAddress, testDBName); err != nil {
		t.Fatal("Failed to connect to db with: %s", err)
	}
	SetupDB(false)
}
