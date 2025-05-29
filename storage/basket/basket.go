package basket

import (
	"regexp"
	"sync"
)

type Basket struct {
	// Storage is the interface for storage operations.
	store map[string]string
	m     *sync.RWMutex
}

// NewBasket creates a new instance of Basket.
func NewBasket() *Basket {
	return &Basket{
		store: make(map[string]string),
		m:     &sync.RWMutex{},
	}
}

// Set adds or updates an item in the basket.
func (b *Basket) Set(key, value string) {
	b.m.Lock()
	b.store[key] = value
	b.m.Unlock()
}

// Get retrieves an item from the basket.
func (b *Basket) Get(key string) (string, bool) {
	b.m.RLock()
	value, exists := b.store[key]
	b.m.RUnlock()
	return value, exists
}

// Delete removes an item from the basket.
func (b *Basket) Delete(key string) {
	b.m.Lock()
	delete(b.store, key)
	b.m.Unlock()
}

// Exists checks if an item exists in the basket.
func (b *Basket) Exists(key string) bool {
	b.m.RLock()
	_, exists := b.store[key]
	b.m.RUnlock()
	return exists
}

// Clear empties the basket.
func (b *Basket) Clear() {
	b.m.Lock()
	b.store = make(map[string]string)
	b.m.Unlock()
}

// Keys returns all keys in the basket matching the given regular expression pattern.
func (b *Basket) Keys(pattern string) ([]string, error) {
	if pattern == "" {
		return nil, nil
	}

	var keys []string
	var re *regexp.Regexp
	var err error

	// Compile the regular expression if a pattern is provided
	if pattern != "" {
		re, err = regexp.Compile(pattern)
		if err != nil {
			// If the pattern is invalid, return an empty list
			return keys, err
		}
	}

	b.m.RLock()
	for key := range b.store {
		// Match the key against the regular expression
		if re == nil || re.MatchString(key) {
			keys = append(keys, key)
		}
	}
	b.m.RUnlock()

	return keys, nil
}

// Count returns the number of items in the basket.
func (b *Basket) Count() int {
	b.m.RLock()
	count := len(b.store)
	b.m.RUnlock()
	return count
}
