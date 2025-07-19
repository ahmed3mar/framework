package events

import (
	"time"
)

type KeyWritten struct {
	CacheEvent
	Value interface{}
	TTL   *time.Duration
}

func NewKeyWritten(storeName, key string, value interface{}, ttl *time.Duration, tags ...string) *KeyWritten {
	return &KeyWritten{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
		Value:      value,
		TTL:        ttl,
	}
}
