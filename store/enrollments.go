package store

type EnrollmentStore struct {
	DefaultStore
}

func NewEnrollmentStore() EnrollmentStore {
	return EnrollmentStore{DefaultStore: NewDefaultStore("enrollments")}
}
