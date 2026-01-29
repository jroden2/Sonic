package sonic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSonicOptions_Sanitise(t *testing.T) {
	tests := []struct {
		name       string
		options    SonicOptions
		expOptions SonicOptions
	}{
		{
			"With 0 capacity - Expects 30",
			SonicOptions{
				Capacity: 0,
				TTL:      30 * time.Second,
			},
			SonicOptions{
				Capacity: 30,
				TTL:      30 * time.Second,
			},
		},
		{
			"With 0.5s TTL - Expects 2s",
			SonicOptions{
				Capacity: 5,
				TTL:      500 * time.Millisecond,
			},
			SonicOptions{
				Capacity: 5,
				TTL:      2 * time.Second,
			},
		},
		{
			"With 0 capacity, 0.5s TTL - Expects 30 capacity, 2s TTL",
			SonicOptions{
				Capacity: 0,
				TTL:      500 * time.Millisecond,
			},
			SonicOptions{
				Capacity: 30,
				TTL:      2 * time.Second,
			},
		},
		{
			"With 50 capacity, 30s TTL - Expects 50 capacity, 30s TTL",
			SonicOptions{
				Capacity: 50,
				TTL:      30 * time.Second,
			},
			SonicOptions{
				Capacity: 50,
				TTL:      30 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.options.Sanitise()
			assert.Equal(t, tt.expOptions.Capacity, tt.options.Capacity)
			assert.Equal(t, tt.expOptions.TTL, tt.options.TTL)
		})
	}
}
