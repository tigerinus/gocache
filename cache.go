package gocache

import "time"

type Cache[K comparable, V any] struct {
	cache map[K]Expirable[V]
}

func (c Cache[K, V]) Get(key K) *V {
	expirable, ok := c.cache[key]

	if !ok {
		return nil
	}

	if expirable.expirationTime < time.Now().UnixMilli() {
		delete(c.cache, key)
		return nil
	}

	return &expirable.data
}

func (c *Cache[K, V]) Put(key K, value V, ttl int64) {
	if ttl <= 0 {
		return
	}

	c.cache[key] = NewExpirable(value, ttl)
}

func NewCache[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		cache: make(map[K]Expirable[V], 0),
	}
}
