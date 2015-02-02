package store

// m "github.com/gogrademe/apiserver/model"
// r "github.com/dancannon/gorethink"

type EmailConfirmationStore struct {
	DefaultStore
}

func NewEmailConfirmationStore() EmailConfirmationStore {
	return EmailConfirmationStore{DefaultStore: NewDefaultStore("emailConfimations")}
}
