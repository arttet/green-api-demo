// Package app provides the main application structure and its operational logic.
package app

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/rs/cors"

	"github.com/arttet/green-api-demo/internal/config"
	"github.com/arttet/green-api-demo/internal/handler"
	"github.com/arttet/green-api-demo/internal/middleware"
	"github.com/arttet/green-api-demo/internal/telemetry"
)

// App holds the application's configuration, Green API proxy handler, and logger.
type App struct {
	cfg    *config.AppConfig
	handle *handler.GreenAPIProxy
	logger *slog.Logger
}

// New creates and initializes a new App instance.
func New(
	cfg *config.AppConfig,
	handle *handler.GreenAPIProxy,
	logger *slog.Logger,
) *App {
	return &App{
		cfg:    cfg,
		handle: handle,
		logger: logger,
	}
}

// Run starts the HTTP server and handles incoming requests.
func (a *App) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/v1/api/proxy/", a.handle.ServeHTTP)
	router.HandleFunc("/health", telemetry.HealthHandler)

	wrappedRouter := middleware.Logging(a.logger)(router)

	corsMiddleware := cors.New(a.cfg.CORS)
	wrappedRouter = corsMiddleware.Handler(wrappedRouter)

	addr := ":" + strconv.Itoa(a.cfg.Server.Port)
	a.logger.Info("server listening",
		slog.Group("http",
			slog.String("address", addr),
		),
	)

	srv := &http.Server{
		Addr:              addr,
		Handler:           wrappedRouter,
		ReadTimeout:       a.cfg.Server.ReadTimeout,
		WriteTimeout:      a.cfg.Server.WriteTimeout,
		IdleTimeout:       a.cfg.Server.IdleTimeout,
		ReadHeaderTimeout: a.cfg.Server.ReadHeaderTimeout,
	}

	return srv.ListenAndServe()
}
