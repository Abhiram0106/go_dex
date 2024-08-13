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
	mux                      *sync.Mutex
	cachedIdToLocations      map[string]cacheEntry
	cachedLocationToPokemons map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux:                      &sync.Mutex{},
		cachedIdToLocations:      make(map[string]cacheEntry),
		cachedLocationToPokemons: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) AddIdToLocations(id string, locations []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cachedIdToLocations[id] = cacheEntry{
		createdAt: time.Now(),
		val:       locations,
	}
}

func (c *Cache) GetIdToLocations(id string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cached, exists := c.cachedIdToLocations[id]
	if !exists {
		return []byte{}, false
	}
	return cached.val, true
}

func (c *Cache) AddLocationToPokemons(location string, pokemons []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cachedLocationToPokemons[location] = cacheEntry{
		createdAt: time.Now(),
		val:       pokemons,
	}
}

func (c *Cache) GetLocationToPokemons(location string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	cached, exists := c.cachedLocationToPokemons[location]
	if !exists {
		return []byte{}, false
	}
	return cached.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		go c.reapIdToLocations(interval)
		go c.reapLocationToPokemons(interval)
	}
}

func (c *Cache) reapIdToLocations(last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, item := range c.cachedIdToLocations {
		if time.Since(item.createdAt) > last {
			delete(c.cachedIdToLocations, k)
			fmt.Printf("Cache ID->Loc REAPED %v\n", k)
			fmt.Printf("go_dex > ")
		}
	}
}

func (c *Cache) reapLocationToPokemons(last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, item := range c.cachedLocationToPokemons {
		if time.Since(item.createdAt) > last {
			delete(c.cachedLocationToPokemons, k)
			fmt.Printf("Cache Loc->Pokemon REAPED %v\n", k)
			fmt.Printf("go_dex > ")
		}
	}
}
