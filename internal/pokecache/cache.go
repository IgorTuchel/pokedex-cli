package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	interval time.Duration
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache(i time.Duration) *Cache {
	c := &Cache{
		interval: i,
		entries: map[string]cacheEntry{},
	}
	go c.reapLoop()
	return c
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()
    for range ticker.C {
        cutoff := time.Now().Add(-c.interval)
        c.mu.Lock()
        for k, e := range c.entries {
            if e.createdAt.Before(cutoff) {
                delete(c.entries, k)
            }
        }
        c.mu.Unlock()
    }
}