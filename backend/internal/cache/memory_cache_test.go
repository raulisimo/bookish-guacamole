package cache

import (
	"testing"
	"time"
)

func TestInMemoryCache_SetAndGet(t *testing.T) {
	cache := NewInMemoryCache()

	// Test setting and getting a value
	key := "key1"
	value := "value1"
	ttl := 5 * time.Second

	cache.Set(key, value, ttl)

	retrievedValue, found := cache.Get(key)
	if !found {
		t.Errorf("Expected to find key %q in cache", key)
	}
	if retrievedValue != value {
		t.Errorf("Expected value %q, got %q", value, retrievedValue)
	}
}

func TestInMemoryCache_Expiration(t *testing.T) {
	cache := NewInMemoryCache()

	// Set a value with a short TTL
	key := "key2"
	value := "value2"
	ttl := 1 * time.Second

	cache.Set(key, value, ttl)

	// Immediately check that the value is retrievable
	retrievedValue, found := cache.Get(key)
	if !found {
		t.Errorf("Expected to find key %q in cache", key)
	}
	if retrievedValue != value {
		t.Errorf("Expected value %q, got %q", value, retrievedValue)
	}

	// Wait for the item to expire
	time.Sleep(2 * time.Second)

	retrievedValue, found = cache.Get(key)
	if found {
		t.Errorf("Expected key %q to expire, but it was found", key)
	}
	if retrievedValue != nil {
		t.Errorf("Expected nil value for expired key %q, but got %v", key, retrievedValue)
	}
}

func TestInMemoryCache_Overwrite(t *testing.T) {
	cache := NewInMemoryCache()

	// Set a value
	key := "key3"
	value1 := "value3a"
	value2 := "value3b"
	ttl := 5 * time.Second

	cache.Set(key, value1, ttl)

	// Overwrite the value
	cache.Set(key, value2, ttl)

	// Check that the new value is stored
	retrievedValue, found := cache.Get(key)
	if !found {
		t.Errorf("Expected to find key %q in cache", key)
	}
	if retrievedValue != value2 {
		t.Errorf("Expected value %q, got %q", value2, retrievedValue)
	}
}

func TestInMemoryCache_NoSuchKey(t *testing.T) {
	cache := NewInMemoryCache()

	// Try to get a key that was never set
	retrievedValue, found := cache.Get("nonexistent")
	if found {
		t.Errorf("Expected key %q not to be found, but it was", "nonexistent")
	}
	if retrievedValue != nil {
		t.Errorf("Expected nil value for key %q, but got %v", "nonexistent", retrievedValue)
	}
}
