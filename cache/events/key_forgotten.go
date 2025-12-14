package events

type KeyForgotten struct {
	CacheEvent
}

func NewKeyForgotten(storeName, key string, tags ...string) *KeyForgotten {
	return &KeyForgotten{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
	}
}
