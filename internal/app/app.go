// Package app provides the main application structure and its operational logic.
package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	router := http.NewServeMux()
	router.HandleFunc("/v1/api/proxy/", a.handle.ServeHTTP)
	router.HandleFunc("/health", telemetry.HealthHandler)

	wrappedRouter := middleware.Logging(a.logger)(router)

	corsMiddleware := cors.New(a.cfg.CORS)
	wrappedRouter = corsMiddleware.Handler(wrappedRouter)

	addr := ":" + strconv.Itoa(a.cfg.Server.Port)
	server := &http.Server{
		Addr:              addr,
		Handler:           wrappedRouter,
		ReadTimeout:       a.cfg.Server.ReadTimeout,
		WriteTimeout:      a.cfg.Server.WriteTimeout,
		IdleTimeout:       a.cfg.Server.IdleTimeout,
		ReadHeaderTimeout: a.cfg.Server.ReadHeaderTimeout,
	}

	errChan := make(chan error, 1)

	go func() {
		a.logger.Info("server listening",
			slog.Group("http",
				slog.String("address", addr),
			),
		)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("server failed: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		a.logger.Info("received shutdown signal")
		stop()
	case err := <-errChan:
		a.logger.Error("critical server error", slog.Any("error", err))

		return err
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.cfg.Server.WriteTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server graceful shutdown encountered an error: %w", err)
	}

	a.logger.Info("server exited properly")

	return nil
}
