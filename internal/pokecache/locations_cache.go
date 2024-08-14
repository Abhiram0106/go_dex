package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	mux   *sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux:   &sync.Mutex{},
		cache: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(id string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[id] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(id string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cached, exists := c.cache[id]
	if !exists {
		return []byte{}, exists
	}
	return cached.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		go c.reapCache(interval)
	}
}

func (c *Cache) reapCache(last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, item := range c.cache {
		if time.Since(item.createdAt) > last {
			delete(c.cache, k)
			fmt.Printf("Cache ID REAPED %v\n", k)
			fmt.Printf("go_dex > ")
		}
	}
}
