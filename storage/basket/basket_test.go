package basket

import "testing"

func TestBasketExists(t *testing.T) {
	b := NewBasket()
	b.Set("key1", "value1")

	if !b.Exists("key1") {
		t.Errorf("Expected key1 to exist")
	}

	if b.Exists("key2") {
		t.Errorf("Expected key2 to not exist")
	}
}

func TestBasketClear(t *testing.T) {
	b := NewBasket()
	b.Set("key1", "value1")
	b.Set("key2", "value2")

	b.Clear()

	if b.Count() != 0 {
		t.Errorf("Expected basket to be empty after Clear")
	}
}

func TestBasketCount(t *testing.T) {
	b := NewBasket()
	if b.Count() != 0 {
		t.Errorf("Expected count to be 0, got %d", b.Count())
	}

	b.Set("key1", "value1")
	b.Set("key2", "value2")

	if b.Count() != 2 {
		t.Errorf("Expected count to be 2, got %d", b.Count())
	}
}

func TestBasketKeysInvalidPattern(t *testing.T) {
	b := NewBasket()
	b.Set("key1", "value1")
	b.Set("key2", "value2")

	_, err := b.Keys("[invalid")
	if err == nil {
		t.Errorf("Expected error for invalid regex pattern")
	}
}

func TestBasketKeysEmptyPattern(t *testing.T) {
	b := NewBasket()
	b.Set("key1", "value1")
	b.Set("key2", "value2")

	keys, err := b.Keys("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(keys) != 0 {
		t.Errorf("Expected no keys for empty pattern, got %v", keys)
	}
}
