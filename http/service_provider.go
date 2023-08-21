package http

import (
	"github.com/goravel/framework/contracts/cache"
	"github.com/goravel/framework/contracts/config"
	consolecontract "github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/exception"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/log"
	"github.com/goravel/framework/contracts/translation"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/http/console"
)

const BindingRateLimiter = "goravel.rate_limiter"

type ServiceProvider struct{}

var (
	TranslationFacade translation.Translation
	CacheFacade       cache.Cache
	LogFacade         log.Log
	ExceptionFacade   exception.Exception
	RateLimiterFacade http.RateLimiter
	ValidationFacade  validation.Validation
	ConfigFacade      config.Config
)

func (http *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(BindingRateLimiter, func(app foundation.Application) (any, error) {
		return NewRateLimiter(), nil
	})
}

func (http *ServiceProvider) Boot(app foundation.Application) {
	CacheFacade = app.MakeCache()
	ConfigFacade = app.MakeConfig()
	RateLimiterFacade = app.MakeRateLimiter()
	ExceptionFacade = app.MakeException()
	TranslationFacade = app.MakeTranslation()

	http.registerCommands(app)
}

func (http *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RequestMakeCommand{},
		&console.ControllerMakeCommand{},
		&console.MiddlewareMakeCommand{},
	})
}
