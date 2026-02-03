// Package telemetry provides functionality for application observability.
package telemetry

import (
	"log/slog"
	"os"
)

// InitLoggerProvider initializes and returns a new slog.Logger instance
// with a JSON handler that writes to os.Stdout.
func InitLoggerProvider(level slog.Level) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:       level,
		AddSource:   true,
		ReplaceAttr: nil,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	return logger
}
