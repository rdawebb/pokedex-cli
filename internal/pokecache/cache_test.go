package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	retrievedValue, ok := cache.Get(key)
	if !ok {
		t.Fatalf("Expected to find key %s in cache", key)
	}

	if string(retrievedValue) != string(value) {
		t.Fatalf("Expected value %s, got %s", value, retrievedValue)
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	// Wait for longer than the TTL (5 seconds)
	time.Sleep(6 * time.Second)

	_, ok := cache.Get(key)
	if ok {
		t.Fatalf("Expected key %s to have expired from cache", key)
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := []byte("testValue")

	done := make(chan bool)

	// Writer goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Add(key, value)
		}
		done <- true
	}()

	// Reader goroutine
	go func() {
		for i := 0; i < 100; i++ {
			cache.Get(key)
		}
		done <- true
	}()

	// Wait for both goroutines to finish
	<-done
	<-done
}