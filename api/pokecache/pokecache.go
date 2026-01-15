// Package pokecache cache implementation for
// LocationAreas api call
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	val, ok := c.cache[key]
	c.mu.Unlock()

	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	for range t.C {
		cutoffTime := time.Now().UTC().Add(-interval)
		c.mu.Lock()
		for key, cache := range c.cache {
			if cache.createdAt.Before(cutoffTime) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
