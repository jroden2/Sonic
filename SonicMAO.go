package sonic

import (
	"time"
)

// Options
type SonicOptions struct {
	Capacity int
	TTL      time.Duration
}

const (
	DefaultMaxCapacity = 30
	DefaultTimeToLive  = 2 * time.Second
)

func (o *SonicOptions) Sanitise() {
	if o.Capacity <= 1 {
		o.Capacity = DefaultMaxCapacity
	}
	if o.TTL <= 1*time.Second {
		o.TTL = DefaultTimeToLive
	}
}

type Entry struct {
	Key    interface{}
	Value  interface{}
	Expiry time.Time
}
