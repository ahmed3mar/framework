package cache

import (
	"github.com/goravel/framework/contracts/cache"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/contracts/log"
)

type Application struct {
	cache.Driver
	events event.Instance
	config config.Config
	log    log.Log
	driver *Driver
	stores map[string]cache.Driver
}

func NewApplication(config config.Config, events event.Instance, log log.Log, store string) (*Application, error) {
	driver := NewDriver(config)
	instance, err := driver.New(store)
	if err != nil {
		return nil, err
	}

	repo := NewRepository(instance, events, store, config)

	return &Application{
		Driver: repo,
		events: events,
		config: config,
		log:    log,
		driver: driver,
		stores: map[string]cache.Driver{
			store: repo,
		},
	}, nil
}

func (app *Application) Store(name string) cache.Driver {
	if driver, exist := app.stores[name]; exist {
		return driver
	}

	instance, err := app.driver.New(name)
	if err != nil {
		app.log.Error(err)

		return nil
	}

	app.stores[name] = NewRepository(instance, app.events, name, app.config)

	return instance
}
