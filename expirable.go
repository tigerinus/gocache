package gocache

import "time"

type Expirable[T any] struct {
	data           T
	expirationTime int64
}

func NewExpirable[T any](data T, ttl int64) Expirable[T] {
	return Expirable[T]{
		data:           data,
		expirationTime: time.Now().UnixMilli() + ttl,
	}
}
