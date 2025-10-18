package pokecache

import (
	"context"
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mutex sync.RWMutex
	cancel context.CancelFunc
}

type cacheEntry struct {
	createdAt time.Time
	val      []byte
}

func NewCache() *Cache {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Cache{
		entry: make(map[string]cacheEntry),
		cancel: cancel,
	}
	go c.reapLoop(ctx, 5*time.Second, 5*time.Second)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:      val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, exists := c.entry[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(ctx context.Context, interval time.Duration, ttl time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mutex.Lock()
			for key, entry := range c.entry {
			if time.Since(entry.createdAt) > ttl {
				delete(c.entry, key)
			}
		}
			c.mutex.Unlock()
			
		case <-ctx.Done():
			return
		}
	}
}

func (c *Cache) Shutdown() {
	if c.cancel != nil {
		c.cancel()
	}
}
