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

type App struct {
	cfg    *config.AppConfig
	handle *handler.GreenAPIProxy
	logger *slog.Logger
}

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

func (a *App) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/v1/api/proxy/", a.handle.ServeHTTP)
	router.HandleFunc("/health", telemetry.HealthHandler)

	wrappedRouter := middleware.Logging(a.logger)(router)

	corsMiddleware := cors.New(a.cfg.CORSConfig)
	wrappedRouter = corsMiddleware.Handler(wrappedRouter)

	addr := ":" + strconv.Itoa(a.cfg.Port)
	a.logger.Info("server listening",
		slog.Group("http",
			slog.String("address", addr),
		),
	)

	return http.ListenAndServe(addr, wrappedRouter)
}
