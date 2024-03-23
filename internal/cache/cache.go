package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]CacheEntry
	mu   *sync.RWMutex
	size int
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(s int) *Cache {
	cache := Cache{
		data: make(map[string]CacheEntry),
		mu:   &sync.RWMutex{},
		size: s,
	}

	go cache.ReapLoop(time.Second*60, time.Minute*5)

	return &cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	if len(c.data) > c.size {
		oldest := time.Now()
		var oldestKey string
		for key, entry := range c.data {
			if entry.createdAt.Before(oldest) {
				oldest = entry.createdAt
				oldestKey = key
			}
		}
		delete(c.data, oldestKey)
	}
}

// ReapLoop deletes entries that are older than exp
// It takes 2 inputs, t and exp, which are both of type time.Duration
// t is the time interval between each check
// exp is the expiration time
func (c *Cache) ReapLoop(t time.Duration, exp time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > exp {
				c.mu.Lock()
				delete(c.data, key)
				c.mu.Unlock()
			}
		}
	}
}
