package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Cache: make(map[string]CacheEntry),
	}
	go func() {
		for {
			time.Sleep(interval)
			c.reapLoop(interval)
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, ok := c.Cache[key]; ok {
		return value.Val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range c.Cache {
		if time.Since(value.CreatedAt) > interval {
			delete(c.Cache, key)
		}
	}
}
