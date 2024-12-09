package cache

import (
	"sync"
	"time"
)

// Cache defines the interface for caching data
type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
}
type cacheItem struct {
	value      interface{}
	expiration int64
}

type inMemoryCache struct {
	data map[string]cacheItem
	mu   sync.Mutex
}

func NewInMemoryCache() Cache {
	return &inMemoryCache{
		data: make(map[string]cacheItem),
	}
}

func (c *inMemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	expiration := time.Now().Add(ttl).Unix()

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheItem{value: value, expiration: expiration}
}

func (c *inMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.data[key]
	if !found {
		return nil, false
	}

	// Check if the item has expired
	if time.Now().Unix() > item.expiration {
		delete(c.data, key)
		return nil, false
	}

	return item.value, true
}