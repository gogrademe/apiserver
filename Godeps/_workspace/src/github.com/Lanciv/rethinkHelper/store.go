package rethink

import (
	"errors"

	r "github.com/dancannon/gorethink"
)

// DB represents a connection to the DB and allows you to run queriedb.
type DB struct {
	Session     *r.Session
	collections []*Collection
}

// Connect establishes connection with rethinkDB
func Connect(connectOpts r.ConnectOpts) (DB, error) {

	var err error
	db := DB{}
	// db.Session, err = r.Connect(r.ConnectOpts{
	// 	Address:     address,
	// 	Database:    database,
	// 	MaxIdle:     10,
	// 	IdleTimeout: time.Second * 10,
	// })
	db.Session, err = r.Connect(connectOpts)
	if err != nil {
		return db, err
	}

	return db, nil
}

// NewDBFromSession returns a new DB from an existing gorethink session.
func NewDBFromSession(session *r.Session) DB {
	return DB{Session: session}
}

// NewCollection returns a new collection with the db set.
func (db *DB) NewCollection(name string) *Collection {
	c := &Collection{r.Table(name), db}
	db.collections = append(db.collections, c)
	return c
}

// CreateTables will create all tables in rethinkdb
func (db *DB) CreateTables() error {
	return errors.New("not implemented")
}

// RunWrite will run a query for the current session.
func (db *DB) RunWrite(term r.Term) (r.WriteResponse, error) {
	writeRes := r.WriteResponse{}
	writeRes, err := term.RunWrite(db.Session)
	if err != nil {
		return writeRes, err
	}

	return writeRes, nil
}

// Run will run a query and return the cursor.
func (db *DB) Run(term r.Term) (*r.Cursor, error) {
	cursor, err := term.Run(db.Session)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

// All will run a query and return the results scanned into an interface.
func (db *DB) All(i interface{}, term r.Term) error {
	cursor, err := db.Run(term)
	if err != nil {
		return err
	}

	err = cursor.All(i)
	if err != nil {
		return err
	}

	return nil
}
