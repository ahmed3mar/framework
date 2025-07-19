package events

import (
	"time"
)

type KeyWriteFailed struct {
	CacheEvent
	Value interface{}
	TTL   *time.Duration
	Error error
}

func NewKeyWriteFailed(storeName, key string, value interface{}, ttl *time.Duration, err error, tags ...string) *KeyWriteFailed {
	return &KeyWriteFailed{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
		Value:      value,
		TTL:        ttl,
		Error:      err,
	}
}
