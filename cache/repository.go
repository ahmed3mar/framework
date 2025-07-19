package cache

import (
	"context"
	"time"

	"github.com/goravel/framework/cache/events"
	"github.com/goravel/framework/contracts/cache"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/testing/docker"
)

type Repository struct {
	store     cache.Driver
	events    event.Instance
	storeName string
	config    config.Config
	tags      []string
}

func NewRepository(store cache.Driver, events event.Instance, storeName string, config config.Config) *Repository {
	return &Repository{
		store:     store,
		events:    events,
		storeName: storeName,
		config:    config,
		tags:      []string{},
	}
}

// Add an item in the cache if the key does not exist.
func (r *Repository) Add(key string, value any, t time.Duration) bool {
	return r.store.Add(key, value, t)
}

// Decrement decrements the value of an item in the cache.
func (r *Repository) Decrement(key string, value ...int64) (int64, error) {
	return r.store.Decrement(key, value...)
}

// Forever add an item in the cache indefinitely.
func (r *Repository) Forever(key string, value any) bool {
	r.fireEvent(events.NewWritingKey(r.storeName, key, value, nil, r.tags...))

	success := r.store.Forever(key, value)

	if success {
		r.fireEvent(events.NewKeyWritten(r.storeName, key, value, nil, r.tags...))
	} else {
		r.fireEvent(events.NewKeyWriteFailed(r.storeName, key, value, nil, nil, r.tags...))
	}

	return success
}

// Forget removes an item from the cache.
func (r *Repository) Forget(key string) bool {
	r.fireEvent(events.NewForgettingKey(r.storeName, key, r.tags...))

	success := r.store.Forget(key)

	if success {
		r.fireEvent(events.NewKeyForgotten(r.storeName, key, r.tags...))
	} else {
		r.fireEvent(events.NewKeyForgetFailed(r.storeName, key, nil, r.tags...))
	}

	return success
}

// Flush remove all items from the cache.
func (r *Repository) Flush() bool {
	r.fireEvent(events.NewCacheFlushing(r.storeName, r.tags...))

	success := r.store.Flush()

	if success {
		r.fireEvent(events.NewCacheFlushed(r.storeName, r.tags...))
	} else {
		r.fireEvent(events.NewCacheFlushFailed(r.storeName, nil, r.tags...))
	}

	return success
}

// Get retrieve an item from the cache by key.
func (r *Repository) Get(key string, def ...any) any {
	r.fireEvent(events.NewRetrievingKey(r.storeName, key, r.tags...))

	value := r.store.Get(key, def...)

	if value != nil && (len(def) == 0 || value != def[0]) {
		r.fireEvent(events.NewCacheHit(r.storeName, key, value, r.tags...))
	} else {
		r.fireEvent(events.NewCacheMissed(r.storeName, key, r.tags...))
	}

	return value
}

// GetBool retrieves an item from the cache by key as a boolean.
func (r *Repository) GetBool(key string, def ...bool) bool {
	return r.store.GetBool(key, def...)
}

// GetInt retrieves an item from the cache by key as an integer.
func (r *Repository) GetInt(key string, def ...int) int {
	return r.store.GetInt(key, def...)
}

// GetInt64 retrieves an item from the cache by key as a 64-bit integer.
func (r *Repository) GetInt64(key string, def ...int64) int64 {
	return r.store.GetInt64(key, def...)
}

// GetString retrieves an item from the cache by key as a string.
func (r *Repository) GetString(key string, def ...string) string {
	return r.store.GetString(key, def...)
}

// Has check an item exists in the cache.
func (r *Repository) Has(key string) bool {
	return r.store.Has(key)
}

// Increment increments the value of an item in the cache.
func (r *Repository) Increment(key string, value ...int64) (int64, error) {
	return r.store.Increment(key, value...)
}

// Lock get a lock instance, the lock will not be expired if the second parameter is not set.
func (r *Repository) Lock(key string, t ...time.Duration) cache.Lock {
	return r.store.Lock(key, t...)
}

// Put Driver an item in the cache for a given time.
func (r *Repository) Put(key string, value any, t time.Duration) error {
	r.fireEvent(events.NewWritingKey(r.storeName, key, value, &t, r.tags...))

	err := r.store.Put(key, value, t)

	if err == nil {
		r.fireEvent(events.NewKeyWritten(r.storeName, key, value, &t, r.tags...))
	} else {
		r.fireEvent(events.NewKeyWriteFailed(r.storeName, key, value, &t, err, r.tags...))
	}

	return err
}

// Pull retrieve an item from the cache and delete it.
func (r *Repository) Pull(key string, def ...any) any {
	value := r.Get(key, def...)
	r.Forget(key)
	return value
}

// Remember gets an item from the cache, or execute the given Closure and store the result.
func (r *Repository) Remember(key string, ttl time.Duration, callback func() (any, error)) (any, error) {
	return r.store.Remember(key, ttl, callback)
}

// RememberForever get an item from the cache, or execute the given Closure and store the result forever.
func (r *Repository) RememberForever(key string, callback func() (any, error)) (any, error) {
	return r.store.RememberForever(key, callback)
}

// WithContext returns a new Cache instance with the given context.
func (r *Repository) WithContext(ctx context.Context) cache.Driver {
	return &Repository{
		store:     r.store.WithContext(ctx),
		events:    r.events,
		storeName: r.storeName,
		config:    r.config,
		tags:      r.tags,
	}
}

// Docker gets the docker driver.
func (r *Repository) Docker() (docker.CacheDriver, error) {
	return r.store.Docker()
}

// Tags returns a new repository instance with tags applied.
func (r *Repository) Tags(tags ...string) cache.Repository {
	return &Repository{
		store:     r.store,
		events:    r.events,
		storeName: r.storeName,
		config:    r.config,
		tags:      tags,
	}
}

// Many retrieve multiple items from the cache by key.
func (r *Repository) Many(keys []string) map[string]any {
	r.fireEvent(events.NewRetrievingManyKeys(r.storeName, keys, r.tags...))

	result := make(map[string]any)
	for _, key := range keys {
		result[key] = r.Get(key)
	}

	return result
}

// PutMany store multiple items in the cache for a given time.
func (r *Repository) PutMany(values map[string]any, t time.Duration) error {
	r.fireEvent(events.NewWritingManyKeys(r.storeName, values, &t, r.tags...))

	for key, value := range values {
		if err := r.Put(key, value, t); err != nil {
			return err
		}
	}

	return nil
}

// fireEvent fires an event through the facades.Event system.
func (r *Repository) fireEvent(e event.Event) {
	if r.events != nil {
		r.events.Dispatch(e)
	}
}
