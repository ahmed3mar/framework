package events

type ForgettingKey struct {
	CacheEvent
}

func NewForgettingKey(storeName, key string, tags ...string) *ForgettingKey {
	return &ForgettingKey{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
	}
}
