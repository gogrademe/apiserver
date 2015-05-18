package store

type EmailConfirmationStore struct {
	DefaultStore
}

func NewEmailConfirmationStore() EmailConfirmationStore {
	return EmailConfirmationStore{DefaultStore: NewDefaultStore("emailConfimations")}
}
