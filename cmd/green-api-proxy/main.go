package main

import (
	"log"
	"log/slog"

	"github.com/arttet/green-api-demo/internal/app"
	"github.com/arttet/green-api-demo/internal/config"
	"github.com/arttet/green-api-demo/internal/handler"
	"github.com/arttet/green-api-demo/internal/telemetry"
)

func main() {
	app, err := initialize()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err = app.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}

func initialize() (*app.App, error) {
	logger := telemetry.InitLoggerProvider(slog.LevelInfo)

	logger.Info("application starting")

	cfg := config.NewAppConfigBuilder().
		WithPortFromEnv("APP_PORT").
		Build()

	handler, err := handler.NewGreenAPIProxy(cfg.APIBaseURL, logger)
	if err != nil {
		return nil, err
	}

	app := app.New(
		cfg,
		handler,
		logger,
	)

	return app, nil
}
