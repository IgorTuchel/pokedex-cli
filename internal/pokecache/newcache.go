package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, value []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	newCache := cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
	c.entries[key] = newCache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	obj, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return obj.val, true
}