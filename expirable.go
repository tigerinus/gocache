package gocache

import "time"

type Expirable[K, V any] struct {
	key            K
	value          *V
	expirationTime int64
}

func NewExpirable[K, V any](key K, value V, ttl int64) Expirable[K, V] {
	return Expirable[K, V]{
		key:            key,
		value:          &value,
		expirationTime: time.Now().UnixMilli() + ttl,
	}
}
