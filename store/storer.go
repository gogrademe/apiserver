package store

// Storer interface
type Storer interface {
	Store()
	Update()
	Find()
}
