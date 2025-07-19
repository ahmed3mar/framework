package events

import (
	"time"
)

type WritingKey struct {
	CacheEvent
	Value interface{}
	TTL   *time.Duration
}

func NewWritingKey(storeName, key string, value interface{}, ttl *time.Duration, tags ...string) *WritingKey {
	return &WritingKey{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
		Value:      value,
		TTL:        ttl,
	}
}
