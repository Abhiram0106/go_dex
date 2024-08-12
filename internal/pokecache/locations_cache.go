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
	mux                 *sync.Mutex
	cachedIdToLocations map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux:                 &sync.Mutex{},
		cachedIdToLocations: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

// func (c *Cache) reapLoop(interval time.Duration) {
//
// 	ticker := time.NewTicker(interval)
// 	defer c.mux.Unlock()
//
// 	go func() {
// 		for range ticker.C {
// 			c.mux.Lock()
// 			for x, item := range c.cachedIdToLocations {
// 				if time.Since(item.createdAt) > interval {
// 					delete(c.cachedIdToLocations, x)
// 					fmt.Printf("Cache REAPED %v\n", x)
// 				}
// 			}
// 			c.mux.Unlock()
// 		}
// 	}()
// }

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cachedIdToLocations[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cached, exists := c.cachedIdToLocations[key]
	if !exists {
		return []byte{}, false
	}
	return cached.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, item := range c.cachedIdToLocations {
		if time.Since(item.createdAt) > last {
			delete(c.cachedIdToLocations, k)
			fmt.Printf("Cache REAPED %v\n", k)
			fmt.Printf("go_dex > ")
		}
	}
}
