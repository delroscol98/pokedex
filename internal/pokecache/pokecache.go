package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]CacheEntry
	mu   sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data: map[string]CacheEntry{},
	}

	go cache.reapLoop(interval)

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	c.data[key] = entry

	for key := range c.data {
		fmt.Println(key)
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.data[key]
	c.mu.Unlock()
	if exists {
		return entry.val, exists
	}

	return nil, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(interval)
		c.mu.Lock()
		for key, cacheEntry := range c.data {
			if cacheEntry.createdAt.Before(cutoff) {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
