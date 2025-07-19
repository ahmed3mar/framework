package events

type CacheFlushing struct {
	CacheEvent
}

func NewCacheFlushing(storeName string, tags ...string) *CacheFlushing {
	return &CacheFlushing{
		CacheEvent: NewCacheEvent(storeName, "", tags...),
	}
}
