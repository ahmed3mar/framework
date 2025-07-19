package events

type CacheFlushFailed struct {
	CacheEvent
	Error error
}

func NewCacheFlushFailed(storeName string, err error, tags ...string) *CacheFlushFailed {
	return &CacheFlushFailed{
		CacheEvent: NewCacheEvent(storeName, "", tags...),
		Error:      err,
	}
}
