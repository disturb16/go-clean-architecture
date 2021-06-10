package api

import (
	"github.com/disturb16/go-sqlite-service/settings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
)

// NewServer creates new echo server object and registers start and end of lifecycle of app
// to start echo on start and gracefully shut it down on exit
func NewServer(lc fx.Lifecycle, cfg *settings.Settings) *echo.Echo {
	e := echo.New()

	// avoid any native logging of echo, because we use custom library for logging
	e.HideBanner = true        // don't log the banner on startup
	e.HidePort = true          // hide log about port server started on
	e.Logger.SetLevel(log.OFF) // disable echo#Logger

	return e
}
