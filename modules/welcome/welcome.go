package welcome

import (
	"go.uber.org/zap"
	"github.com/labstack/echo/v4"
	"github.com/moadkey/gosvelte/system/gosvelte"
)

type Handler struct {
	mux      *echo.Echo
	logger   *zap.SugaredLogger
	gosvelte *gosvelte.GoSvelte
}

func New(s *echo.Echo, logger *zap.SugaredLogger, goSvelte *gosvelte.GoSvelte) *Handler {
	h := Handler{s, logger, goSvelte}
	h.registerRoutes()

	return &h
}

func (h *Handler) registerRoutes() {
	h.mux.GET("/", h.welcome)
}

func (h *Handler) welcome(ctx echo.Context) error {
	lang := "en"
	dir  := "ltr"
	name := "GoSvelte"
	data := gosvelte.Data{"lang": lang, "dir": dir, "name": name}

	goSvelteServeReq := gosvelte.ServeReq{
		SvelteFileName:				"welcome",
		//SvelteOptionsTag:			"gosvelte-welcome-page",
		HtmlGlobalLangAttribute:	lang,
		HtmlGlobalDirAttribute:		dir,
		Title:						name,
		Data:						data,
	}

	err := h.gosvelte.Serve(ctx.Response().Writer, goSvelteServeReq)
	if err != nil {
		h.logger.Error(err)
		return err
	}

	return nil
}