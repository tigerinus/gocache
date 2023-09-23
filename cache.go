package gocache

import (
	"time"

	"github.com/tigerinus/gpq"
)

type Cache[K comparable, V any] struct {
	cache    map[K]*Expirable[V]
	capacity int
	pq       gpq.PriorityQueue[*Expirable[V]]
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
	e := NewExpirable(value, ttl)
	c.cache[key] = &e
}

func NewCache[K comparable, V any](capacity int) Cache[K, V] {
	return Cache[K, V]{
		cache:    make(map[K]*Expirable[V], 0),
		capacity: capacity,
		pq: gpq.NewPriorityQueue(func(i, j *Expirable[V]) bool {
			return i.expirationTime > j.expirationTime
		}),
	}
}
