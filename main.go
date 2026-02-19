package main

import (
	"fmt"
	"sync"
)

// Cache is a thread-safe in-memory store using Generics.
// K must be 'comparable' (string, int, etc.), V can be 'any' type.
type Cache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

// NewCache initializes and returns a new Cache instance.
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]V),
	}
}

// Set adds or updates an item in the cache.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()         // Lock for writing
	defer c.mu.Unlock() // Ensure unlock happens after the function ends
	c.items[key] = value
}

// Get retrieves an item from the cache.
// It returns the value and a boolean indicating if the key exists.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock() // RLock allows multiple simultaneous readers
	defer c.mu.RUnlock()
	val, found := c.items[key]
	return val, found
}

func main() {
	// 1. Initialize a cache that stores string keys and int values
	userAgeCache := NewCache[string, int]()

	// 2. Set a value
	fmt.Println("Setting user 'Alice' age to 25...")
	userAgeCache.Set("Alice", 25)

	// 3. Get a value
	age, found := userAgeCache.Get("Alice")
	if found {
		fmt.Printf("Found Alice: %d years old\n", age)
	}

	// 4. Handle missing keys
	_, foundBob := userAgeCache.Get("Bob")
	if !foundBob {
		fmt.Println("User 'Bob' not found in cache.")
	}
}
