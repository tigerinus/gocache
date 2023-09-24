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

func TestPositive2(t *testing.T) {
	cache := gocache.NewCache[string, DummyType]()

	expected1 := DummyType{data: "test1"}
	expected2 := DummyType{data: "test2"}
	expected3 := DummyType{data: "test3"}

	cache.Put(expected1.data, expected1, 1000)
	cache.Put(expected2.data, expected2, 3000)
	cache.Put(expected3.data, expected3, 2000)

	actual1 := cache.Get(expected1.data)
	actual2 := cache.Get(expected2.data)
	actual3 := cache.Get(expected3.data)

	if actual1 == nil || actual2 == nil || actual3 == nil {
		t.Fail()
	}

	if actual1.data != expected1.data || actual2.data != expected2.data || actual3.data != expected3.data {
		t.Fail()
	}

	time.Sleep(1 * time.Second)

	actual1 = cache.Get(expected1.data)
	actual2 = cache.Get(expected2.data)
	actual3 = cache.Get(expected3.data)

	if actual1 != nil || actual2 == nil || actual3 == nil {
		t.Fail()
	}

	if actual2.data != expected2.data || actual3.data != expected3.data {
		t.Fail()
	}

	time.Sleep(1 * time.Second)

	actual1 = cache.Get(expected1.data)
	actual2 = cache.Get(expected2.data)
	actual3 = cache.Get(expected3.data)

	if actual1 != nil || actual2 == nil || actual3 != nil {
		t.Fail()
	}

	if actual2.data != expected2.data {
		t.Fail()
	}

	time.Sleep(1 * time.Second)

	actual1 = cache.Get(expected1.data)
	actual2 = cache.Get(expected2.data)
	actual3 = cache.Get(expected3.data)

	if actual1 != nil || actual2 != nil || actual3 != nil {
		t.Fail()
	}
}

func TestPositive3(t *testing.T) {
	cache := gocache.NewCache[string, DummyType]()

	expected1 := DummyType{data: "test1"}
	expected2 := DummyType{data: "test2"}
	expected3 := DummyType{data: "test3"}

	cache.Put(expected1.data, expected1, 1000)
	cache.Put(expected2.data, expected2, 3000)
	cache.Put(expected3.data, expected3, 2000)

	actual1 := cache.Get(expected1.data)
	actual2 := cache.Get(expected2.data)
	actual3 := cache.Get(expected3.data)

	if actual1 == nil || actual2 == nil || actual3 == nil {
		t.Fail()
	}

	if actual1.data != expected1.data || actual2.data != expected2.data || actual3.data != expected3.data {
		t.Fail()
	}

	actual1 = cache.Get(expected1.data)
	actual2 = cache.Get(expected2.data)
	actual3 = cache.Get(expected3.data)

	if actual1 == nil || actual2 == nil || actual3 == nil {
		t.Fail()
	}

	if actual1.data != expected1.data || actual2.data != expected2.data || actual3.data != expected3.data {
		t.Fail()
	}

	cache.Purge(time.Now().Add(5 * time.Second).UnixMilli())

	actual1 = cache.Get(expected1.data)
	actual2 = cache.Get(expected2.data)
	actual3 = cache.Get(expected3.data)

	if actual1 != nil || actual2 != nil || actual3 != nil {
		t.Fail()
	}
}

func TestNegative1(t *testing.T) {
	cache := gocache.NewCache[string, DummyType]()

	nonExisting := cache.Get("non-existing")
	if nonExisting != nil {
		t.Fail()
	}
}
