package sonic

import (
	"fmt"
	"sync"
	"time"
)

type sonicCache struct {
	entries map[any]Entry
	ttl     time.Duration
	mu      sync.Mutex
}

func NewSonicCache(opts ...SonicOptions) SonicCache {
	tOpts := SonicOptions{}
	if len(opts) != 0 {
		opts[0].Sanitise()
		tOpts = opts[0]
	}

	fmt.Printf("NewSonicCache %s\n", "Sonic Cache started... Gotta go fast")
	fmt.Printf("NewSonicCache %s %s\n", "Capacity", tOpts.Capacity)
	fmt.Printf("NewSonicCache %s %s\n", "TTL", tOpts.TTL)
	return &sonicCache{
		entries: make(map[any]Entry, tOpts.Capacity),
		ttl:     tOpts.TTL,
	}
}

type SonicCache interface {
	Add(K, V interface{})
	AddRaw(K interface{}, e *Entry)
	Get(K interface{}) (V interface{}, OK bool)
	Exists(K interface{}) bool
	PeekAll() map[any]any
	Purge()
	PurgeExpired()
	Close()
}

// Functions

func (c *sonicCache) Add(K, V interface{}) {
	c.AddRaw(K, &Entry{
		Key:    K,
		Value:  V,
		Expiry: time.Now().Add(c.ttl),
	})
}

func (c *sonicCache) AddRaw(K interface{}, e *Entry) {
	c.entries[K] = *e
}

func (c *sonicCache) Get(K interface{}) (V interface{}, OK bool) {
	v, ok := c.entries[K]
	return v, ok
}

func (c *sonicCache) Exists(K interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.entries[K]; !ok {
		return false
	}
	return true
}

func (c *sonicCache) PurgeExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for {
		TTE := time.Now()
		for _, e := range c.entries {
			if e.Expiry.Before(TTE) {
				fmt.Printf("PurgeExpired {} {}", "Purging", e.Value)
				delete(c.entries, e.Key)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (c *sonicCache) PeekAll() map[any]any {
	c.mu.Lock()
	defer c.mu.Unlock()

	retVal := make(map[any]any)
	for _, e := range c.entries {
		retVal[e.Key] = e.Value
	}
	return retVal
}

func (c *sonicCache) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for e := range c.entries {
		delete(c.entries, e)
	}
}

func (c *sonicCache) Close() {
	c.Purge()
	c.mu.Lock()
	c.mu.Unlock()
}
