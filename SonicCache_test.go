package sonic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSonicCache(t *testing.T) {
	tests := []struct {
		name    string
		options SonicOptions
	}{
		{
			"With Options",
			SonicOptions{
				Capacity: 30,
				TTL:      30 * time.Second,
			},
		},
		{
			"Without Options added",
			SonicOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSonicCache(tt.options)
			assert.NotNil(t, got)
		})
	}
}

var (
	cache = NewSonicCache(SonicOptions{
		Capacity: 5,
		TTL:      2 * time.Second,
	})
	k = "Test"
)

func TestSonicCache_Add(t *testing.T) {
	cache.Add(k, "Hello, World")
	assert.True(t, cache.Exists(k))
	cache.Purge()
}

func TestSonicCache_Get(t *testing.T) {
	cache.Add(k, "Hello, World")
	v, ok := cache.Get(k)
	assert.NotNil(t, v)
	assert.True(t, ok)
	cache.Purge()
}

func TestSonicCache_Exists(t *testing.T) {
	t.Run("Run test with existing entry", func(t *testing.T) {
		cache.Add(k, "Hello, World")
		assert.True(t, cache.Exists(k))
		cache.Purge()
	})
	t.Run("Run test with none existing entry", func(t *testing.T) {
		assert.False(t, cache.Exists(k))
	})
}

func TestSonicCache_PeekAll(t *testing.T) {
	t.Run("Adds two entries, retrieves two entries", func(t *testing.T) {
		cache.Add(k, "Hello, World")
		cache.Add("test2", "Goodnight, Mood")
		keys := []string{}
		for key, _ := range cache.PeekAll() {
			keys = append(keys, key.(string))
		}
		assert.Contains(t, keys, k)
		assert.Contains(t, keys, "test2")
	})
}

func TestSonicCache_Purge(t *testing.T) {
	t.Run("Purges all existing entries", func(t *testing.T) {
		cache.Add(k, "Hello, World")
		cache.Purge()
		assert.False(t, cache.Exists(k))
	})
}
