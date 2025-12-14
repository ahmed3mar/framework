package events

type CacheFlushed struct {
	CacheEvent
}

func NewCacheFlushed(storeName string, tags ...string) *CacheFlushed {
	return &CacheFlushed{
		CacheEvent: NewCacheEvent(storeName, "", tags...),
	}
}
