package storage

type InMemoryDatabase struct{}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}
