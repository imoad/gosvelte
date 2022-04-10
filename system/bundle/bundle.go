package bundle

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"

	"github.com/moadkey/gosvelte/system/config"
	"github.com/moadkey/gosvelte/system/gosvelte"
	"github.com/moadkey/gosvelte/system/httpserver"
	"github.com/moadkey/gosvelte/system/logger"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	cfg *config.Config,
	mux *echo.Echo,
	goSvelte *gosvelte.GoSvelte,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				mux.Static("www/public/*filepath", "www/public/")
				logger.Info("Serving static files at /www/public")
				addr := fmt.Sprintf("%s%s", cfg.ApplicationConfig.Host, cfg.ApplicationConfig.Port)
				go http.ListenAndServe(addr, mux)
				logger.Info(cfg.ApplicationConfig.Name, " listening on ", addr)
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Onstop...")
				return logger.Sync()
			},
		},
	)
}

var Module = fx.Options(
	config.Module,
	logger.Module,
	httpserver.Module,
	gosvelte.Module,
	fx.Invoke(registerHooks),
)
