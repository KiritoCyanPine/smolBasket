package storage

type Manager interface {
	Create(name string) error
	Drop(name string) error
	Info(name string) (Status, error)
	List() ([]string, error)
}

// Status represents the status of a storage.
type Status struct {
	Name      string
	ItemCount int
}
