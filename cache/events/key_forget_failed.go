package events

type KeyForgetFailed struct {
	CacheEvent
	Error error
}

func NewKeyForgetFailed(storeName, key string, err error, tags ...string) *KeyForgetFailed {
	return &KeyForgetFailed{
		CacheEvent: NewCacheEvent(storeName, key, tags...),
		Error:      err,
	}
}
