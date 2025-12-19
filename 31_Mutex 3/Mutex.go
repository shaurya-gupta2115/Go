package main

import "sync"

type SafeCache struct {
	data map[string]string
	mu   sync.RWMutex
}

func (c *SafeCache) Get(key string) string {
	c.mu.RLock() // multiple readers allowed
	value := c.data[key]

	defer func() { c.mu.RUnlock() }()
	return value
}

func (c *SafeCache) Set(key, val string) {
	c.mu.Lock() // writer ke time exclusive lock
	defer c.mu.Unlock()
	c.data[key] = val
}

func main() {

}
