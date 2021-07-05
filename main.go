// persons-service service
//
// Schemes: http
// Host: localhost:8080
// BasePath: /v1
// Version: 1.0
//
// Security:
//     - api_key:
//
// SecurityDefinitions:
//  api_key:
//   type: apiKey
//   name: api-key
//   in: header
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/disturb16/go-sqlite-service/dbutils"
	"github.com/disturb16/go-sqlite-service/internal/api"
	"github.com/disturb16/go-sqlite-service/internal/api/healthcheck"
	v1 "github.com/disturb16/go-sqlite-service/internal/api/v1"
	"github.com/disturb16/go-sqlite-service/internal/persons"
	"github.com/disturb16/go-sqlite-service/internal/persons/repository"
	"github.com/disturb16/go-sqlite-service/internal/persons/repository/rediscache"
	"github.com/disturb16/go-sqlite-service/internal/persons/service"
	"github.com/disturb16/go-sqlite-service/settings"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sanservices/apicore/validator"
	logger "github.com/sanservices/apilogger/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// Set logger to custom one
		fx.Logger(logger.New()),

		fx.Provide(
			// Provide new instances of structs
			// Empty context
			context.Background,
			// Logger
			logger.New,
			// Settings
			settings.New,

			dbutils.New,

			rediscache.New,

			// Repo
			repository.New,
			// Service
			service.New,
			// Validator
			validator.NewValidator,
			// New server
			api.NewServer,
			// Add all handlers here
			func(cfg *settings.Settings, srv persons.Service, vld *validator.Validator) []api.Handler {
				return []api.Handler{
					healthcheck.NewHandler(),     // For Healthchecks
					v1.NewHandler(cfg, srv, vld), // v1
				}
			},
		),
		fx.Invoke(
			// Use logger to initialize it globally
			func(ctx context.Context, l *logger.Logger) {
				logger.Info(ctx, logger.LogCatStartUp, "Initializing the app")
			},

			provideLifeCycleHooks,

			// Register routes
			api.RegisterRoutes,
		),
	)

	app.Run()
}

func provideLifeCycleHooks(lc fx.Lifecycle, e *echo.Echo, cfg *settings.Settings, db *sqlx.DB) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			logger.Infof(ctx, logger.LogCatRouterInit, "server running on port: %d", cfg.Service.Port)

			go func() {
				log.Fatal(
					e.Start(fmt.Sprintf(":%d", cfg.Service.Port)),
				)
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Info(ctx, logger.LogCatDebug, "Closing database...")
			db.Close()

			logger.Info(ctx, logger.LogCatDebug, "Server is shutting down...")
			return e.Shutdown(ctx)
		},
	})
}
