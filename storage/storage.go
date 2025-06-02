package storage

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/KiritoCyanPine/smolBasket/storage/basket"
)

type StorageManager struct {
	Baskets map[string]basket.Database
	m       *sync.RWMutex
}

// NewStorageManager creates a new instance of StorageManager.
func NewStorageManager() *StorageManager {
	return &StorageManager{
		Baskets: make(map[string]basket.Database),
		m:       &sync.RWMutex{},
	}
}

// Create creates a new basket with the given name.
func (sm *StorageManager) Create(name string) error {
	sm.m.Lock()

	if _, exists := sm.Baskets[name]; exists {
		return fmt.Errorf("basket %s already exists", name)
	}

	sm.Baskets[name] = basket.NewBasket()

	sm.m.Unlock()
	return nil
}

// Drop removes the basket with the given name.
func (sm *StorageManager) Drop(name string) error {
	sm.m.Lock()
	if _, exists := sm.Baskets[name]; !exists {
		return fmt.Errorf("basket %s does not exist", name)
	}

	delete(sm.Baskets, name)

	sm.m.Unlock()
	return nil
}

// Info returns the status of the basket with the given name.
func (sm *StorageManager) Info(name string) (string, error) {
	sm.m.RLock()

	db, exists := sm.Baskets[name]
	if !exists {
		return "0", fmt.Errorf("basket %s does not exist", name)
	}

	sm.m.RUnlock()

	return strconv.Itoa(db.Count()), nil
}

// List returns the names of all baskets.
func (sm *StorageManager) List() ([]string, error) {
	sm.m.RLock()

	names := make([]string, 0, len(sm.Baskets))
	for name := range sm.Baskets {
		names = append(names, name)
	}

	sm.m.RUnlock()
	return names, nil
}

func (sm *StorageManager) GetBasket(name string) (basket.Database, error) {
	sm.m.RLock()
	defer sm.m.RUnlock()

	db, exists := sm.Baskets[name]
	if !exists {
		return nil, fmt.Errorf("basket %s does not exist", name)
	}

	return db, nil
}
