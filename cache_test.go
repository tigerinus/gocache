package gocache_test

import (
	"testing"
	"time"

	"github.com/tigerinus/gocache"
)

type DummyType struct {
	data string
}

func TestPositive1(t *testing.T) {
	cache := gocache.NewCache[string, DummyType]()

	expected1 := DummyType{data: "test1"}

	cache.Put("test1", expected1, 0)
	actual1 := cache.Get("test1")
	if actual1 != nil {
		t.Fail()
	}

	cache.Put("test1", expected1, 1000)
	actual1 = cache.Get("test1")
	if actual1 == nil || actual1.data != expected1.data {
		t.Fail()
	}

	time.Sleep(1 * time.Second)
	actual1 = cache.Get("test1")
	if actual1 != nil {
		t.Fail()
	}
}
