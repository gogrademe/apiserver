package model

type Assignment struct {
	ID        int64
	Name      string
	Type      string
	DueDate   int64 `db:"due_date"`
	ClassID   int64 `db:"class_id"`
	ClassTerm int64 `db:"class_term_id"`
	TimeStamp
}
