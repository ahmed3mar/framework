package events

type RetrievingKey struct {
	CacheEvent
}

func NewRetrievingKey(storeName, key string, tags ...string) *RetrievingKey {
	return &RetrievingKey{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
	}
}
