package events

type CacheHit struct {
	CacheEvent
	Value interface{}
}

func NewCacheHit(storeName, key string, value interface{}, tags ...string) *CacheHit {
	return &CacheHit{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
		Value:      value,
	}
}
