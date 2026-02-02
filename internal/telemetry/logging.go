package telemetry

import (
	"log/slog"
	"os"
)

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
