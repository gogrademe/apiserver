package rethink

import (
	"errors"

	r "github.com/dancannon/gorethink"
)

var (
	// ErrRecordNotFound for record not being found
	ErrRecordNotFound = errors.New("record not found")
)

// Collection represents a table in rethinkdb.
// Term should be set to r.Table("CollectionName")
// db will be set to the db that Collection was initialized with.
type Collection struct {
	r.Term
	db *DB
}

// Insert will insert from one to many records and will return their IDs.
func (c *Collection) Insert(arg ...interface{}) ([]string, error) {
	writeRes, err := c.db.RunWrite(c.Term.Insert(arg))
	if err != nil {
		return writeRes.GeneratedKeys, err
	}
	return writeRes.GeneratedKeys, nil
}

// Update will update a single record.
func (c *Collection) Update(arg interface{}, id string) error {
	_, err := c.db.RunWrite(c.Term.Get(id).Update(arg))
	if err != nil {
		return err
	}

	return nil
}

// One return single record from the DB.
func (c *Collection) One(i interface{}, id string) error {
	cursor, err := c.db.Run(c.Term.Get(id))
	if err != nil {
		return err
	}
	if cursor.IsNil() {
		return ErrRecordNotFound
	}

	return cursor.One(i)
}

// All will run a query and return the results scanned into an interface.
// func (c *Collection) All(i interface{}, query r.Term) error {
// 	cursor, err := c.db.Run(c.Term)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = cursor.All(i)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
