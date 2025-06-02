package basket

type Database interface {
	// Set adds or updates an item in the basket.
	Set(key, value string)
	// Get retrieves an item from the basket.
	Get(key string) (string, bool)
	// Delete removes an item from the basket.
	Delete(key string)
	// Exists checks if an item exists in the basket.
	Exists(key string) bool
	// Clear empties the basket.
	Clear()
	// Count returns the number of items in the basket.
	Count() int
	// Keys returns all keys in the basket matching the given regular expression pattern.
	Keys(pattern string) ([]string, error)
}
