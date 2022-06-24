package shared

import "sync"

type CacheSync interface {
	Load(string) (interface{}, bool)
	Delete(string)
	Store(string, interface{})
}

type cache struct {
	sync.RWMutex
	cache map[string]interface{}
}

func NewCacheSync() *cache {
	return &cache{
		cache: make(map[string]interface{}),
	}
}

func (c *cache) Load(key string) (interface{}, bool) {
	c.RLock()
	result, ok := c.cache[key]
	c.RUnlock()
	return result, ok
}

func (c *cache) Delete(key string) {
	c.Lock()
	delete(c.cache, key)
	c.Unlock()
}

func (c *cache) Store(key string, value interface{}) {
	c.Lock()
	c.cache[key] = value
	c.Unlock()
}
