package main

type Cache[K comparable, V any] struct {
	capacity int
	data     map[K]V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	var initialSize int
	if capacity > 0 {
		initialSize = capacity
	}

	return &Cache[K, V]{
		capacity: capacity,
		data:     make(map[K]V, initialSize),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	if c.capacity == 0 {
		return
	}
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	if c.capacity == 0 {
		var zero V
		return zero, false
	}
	val, ok := c.data[key]
	return val, ok
}
