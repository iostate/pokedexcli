package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	data     map[string]cacheEntry
	interval time.Duration

	ticker *time.Ticker
	stopCh chan struct{}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
		ticker:   time.NewTicker(interval),
		stopCh:   make(chan struct{}),
	}
	// Clean cache based on interval duration
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	cacheEntry, ok := c.data[key]

	if !ok {
		return cacheEntry.val, false
	}

	return cacheEntry.val, true
}

// Clean up old cache entries past the interval
// specified in Cache
func (c *Cache) reapLoop() {
	defer c.ticker.Stop()
	for range c.ticker.C {
		c.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.data, key)
			}
		}
		c.Unlock()
	}
}
