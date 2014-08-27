package rethink

import (
	"os"
	"testing"

	r "github.com/dancannon/gorethink"
)

var (
	db      DB
	url     string
	authKey string
)

func init() {
	// Needed for wercker. By default url is "localhost:28015"
	url = os.Getenv("RETHINKDB_URL")
	if url == "" {
		url = "localhost:28015"
	}

	// Needed for running tests for RethinkDB with a non-empty authkey
	authKey = os.Getenv("RETHINKDB_AUTHKEY")
}

func TestConnect(t *testing.T) {
	var err error

	db, err = Connect(r.ConnectOpts{
		Address:  url,
		AuthKey:  authKey,
		Database: "test",
	})

	if err != nil {
		t.Errorf("expected nil error but got %s", err)
	}

	if len(db.collections) != 0 {
		t.Error("new DB should not have any collections.")
	}
}

func TestNewCollection(t *testing.T) {
	TestCollection := db.NewCollection("test")

	if TestCollection.db != &db {
		t.Error("expected TestCollection.db to equal global DB")
	}

	if len(db.collections) != 1 {
		t.Error("new DB should have one collection.")
	}

	if db.collections[0] != TestCollection {
		t.Error("collections should have TestCollection")
	}
}
