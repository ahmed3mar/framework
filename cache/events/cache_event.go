package events

import "github.com/goravel/framework/contracts/event"

type CacheEvent struct {
	StoreName string
	Key       string
	Tags      []string
}

func NewCacheEvent(storeName, key string, tags ...string) CacheEvent {
	return CacheEvent{
		StoreName: storeName,
		Key:       key,
		Tags:      tags,
	}
}

func (c *CacheEvent) SetTags(tags []string) *CacheEvent {
	c.Tags = tags
	return c
}

// Handle the event.
func (c CacheEvent) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
