package httpserver

import (
	"go.uber.org/fx"
	"github.com/labstack/echo/v4"
)

var Module = fx.Options(
	fx.Provide(echo.New),
)
