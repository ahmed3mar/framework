package events

type CacheMissed struct {
	CacheEvent
}

func NewCacheMissed(storeName, key string, tags ...string) *CacheMissed {
	return &CacheMissed{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
	}
}
