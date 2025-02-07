package session

import (
	frameworkconfig "github.com/goravel/framework/config"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/session"
	"github.com/goravel/framework/errors"
)

var (
	SessionFacade session.Manager
	ConfigFacade  config.Config
)

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(frameworkconfig.BindingSession, func(app foundation.Application) (any, error) {
		c := app.MakeConfig()
		if c == nil {
			return nil, errors.ConfigFacadeNotSet.SetModule(errors.ModuleSession)
		}

		j := app.GetJson()
		if j == nil {
			return nil, errors.JSONParserNotSet.SetModule(errors.ModuleSession)
		}

		return NewManager(c, j), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	SessionFacade = app.MakeSession()
	ConfigFacade = app.MakeConfig()
}
