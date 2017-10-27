package main

import (
	"sync"
)

// Cache represents a cache of strings to ints that is
// safe for concurrent use.
type Cache struct {
	entries map[string]int // Map to store the entries.
	// We wanted to use a map to cache frequently accessed data,
	// and we wanted to share that map across multiple goroutines.
	// Since maps are not safe for concurrent access,
	// we need to protect it.
	// But since we expect many more read operations than write operations,
	// we should protect it using an RWMutex instead of the traditional mutex.
	mx sync.RWMutex // RWMutex to protect the map.
}

// NewCache constructs a new Cache.
func NewCache() *Cache {
	return &Cache{
		entries: map[string]int{},
	}
}

// Set adds a new key and value to the map.
func (c *Cache) Set(key string, value int) {
	// Obtain an exclusive lock.
	c.mx.Lock()
	// Set the key/value in the map.
	c.entries[key] = value
	// Release the exclusive lock.
	c.mx.Unlock()
}

// Get retrieves the value for a given key.
func (c *Cache) Get(key string) int {
	// Obtain a read lock.
	// The non-exclusive read lock in the .Get() method can be obtained by multiple goroutines at the same time,
	// so multiple readers can lookup the value for a key at the same time while there are no writers.
	c.mx.RLock()
	// Use defer to release the read lock
	// as we exit the function.
	defer c.mx.RUnlock()
	// Return the value from the map.
	return c.entries[key]
}
