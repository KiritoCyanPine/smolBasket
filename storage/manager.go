package storage

import "github.com/KiritoCyanPine/smolBasket/storage/basket"

type Manager interface {
	Create(name string) error
	Drop(name string) error
	Info(name string) (string, error)
	List() ([]string, error)
	GetBasket(name string) (basket.Database, error)
}
