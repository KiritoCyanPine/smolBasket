package storage

import (
	"testing"
)

func TestStorageManagerCreateDuplicate(t *testing.T) {
	sm := NewStorageManager()

	err := sm.Create("basket1")
	if err != nil {
		t.Fatalf("Failed to create basket: %v", err)
	}

	err = sm.Create("basket1")
	if err == nil {
		t.Fatalf("Expected error when creating duplicate basket, got nil")
	}
}

func TestStorageManagerDropNonExistent(t *testing.T) {
	sm := NewStorageManager()

	err := sm.Create("basket1")
	if err != nil {
		t.Fatalf("Failed to create basket: %v", err)
	}

	err = sm.Drop("basket1")
	if err != nil {
		t.Fatalf("Failed to drop basket: %v", err)
	}

	err = sm.Drop("nonexistent")
	if err == nil {
		t.Fatalf("Expected error when dropping non-existent basket, got nil")
	}
}

func TestStorageManagerInfo(t *testing.T) {
	sm := NewStorageManager()
	sm.Create("basket1")

	info, err := sm.Info("basket1")
	if err != nil {
		t.Fatalf("Failed to get basket info: %v", err)
	}

	if info != "0" {
		t.Errorf("Expected info to be '0', got %q", info)
	}

	_, err = sm.Info("nonexistent")
	if err == nil {
		t.Fatalf("Expected error when getting info for non-existent basket, got nil")
	}
}

func TestStorageManagerGetBasket(t *testing.T) {
	sm := NewStorageManager()
	sm.Create("basket1")

	basket, err := sm.GetBasket("basket1")
	if err != nil {
		t.Fatalf("Failed to get basket: %v", err)
	}

	if basket == nil {
		t.Errorf("Expected basket to be non-nil")
	}

	_, err = sm.GetBasket("nonexistent")
	if err == nil {
		t.Fatalf("Expected error when getting non-existent basket, got nil")
	}
}
