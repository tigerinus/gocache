package gocache

import (
	"sync"
	"time"

	"github.com/tigerinus/gpq"
)

type Cache[K comparable, V any] struct {
	cache map[K]*Expirable[V]
	pq    gpq.PriorityQueue[*Expirable[V]]

	mutex *sync.Mutex
}

func (c Cache[K, V]) Get(key K) *V {
	expirable, ok := c.cache[key]

	if !ok {
		return nil
	}

	if expirable.expirationTime > time.Now().UnixMilli() {
		return &expirable.data
	}

	delete(c.cache, key)

	go c.Purge(expirable.expirationTime) // purge anything expires up to expirable.expirationTime

	return nil
}

func (c *Cache[K, V]) Put(key K, value V, ttl int64) {
	if ttl <= 0 {
		return
	}
	e := NewExpirable(value, ttl)
	c.cache[key] = &e
}

func (c *Cache[K, V]) Purge(upto int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	top := c.pq.Peek()

	for top != nil && top.expirationTime <= upto {
		c.pq.Pop()
	}
}

func NewCache[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		cache: make(map[K]*Expirable[V], 0),
		pq: gpq.NewPriorityQueue(func(i, j *Expirable[V]) bool {
			return i.expirationTime > j.expirationTime
		}),
		mutex: &sync.Mutex{},
	}
}
