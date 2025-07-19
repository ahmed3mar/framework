package events

type RetrievingManyKeys struct {
	CacheEvent
	Keys []string
}

func NewRetrievingManyKeys(storeName string, keys []string, tags ...string) *RetrievingManyKeys {
	return &RetrievingManyKeys{
		CacheEvent: NewCacheEvent(storeName, "multiple_keys", tags...),
		Keys:       keys,
	}
}
