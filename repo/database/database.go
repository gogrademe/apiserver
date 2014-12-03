package database

import "github.com/dancannon/gorethink"

type DBResource interface {
	Session gorethink.Session
	Table gorethink.Term
}

// func Save() {
// 	res, err := db.Table.Insert(grade).RunWrite(&db.Session)
// 	if err != nil {
// 		return err
// 	}
//
// 	if len(res.GeneratedKeys) != 1 {
// 		return errors.New("No keys")
// 	}
//
// 	grade.ID = res.GeneratedKeys[0]
//
// 	return nil
// }
