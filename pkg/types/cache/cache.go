package cache

import "sync"

type Cache[T any] struct {
	storage map[string]T
	mux     sync.RWMutex
}

func New[T any]() *Cache[T] {
	return &Cache[T]{
		storage: make(map[string]T),
	}
}

func (c *Cache[T]) Set(key string, value T) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.storage[key] = value
}

func (c *Cache[T]) Drain() map[string]T {
	c.mux.Lock()
	defer c.mux.Unlock()

	storage := c.storage

	c.storage = make(map[string]T)

	return storage
}
