package events

import "time"

type WritingManyKeys struct {
	CacheEvent
	Values map[string]interface{}
	TTL    *time.Duration
}

func NewWritingManyKeys(storeName string, values map[string]interface{}, ttl *time.Duration, tags ...string) *WritingManyKeys {
	return &WritingManyKeys{
		CacheEvent: NewCacheEvent(storeName, "multiple_keys", tags...),
		Values:     values,
		TTL:        ttl,
	}
}
