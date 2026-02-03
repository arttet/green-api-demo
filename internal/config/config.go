// Package config provides application configuration structures and utilities.
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/cors"
)

const (
	defaultPortMaxAge = 300

	defaultPort              = 8080
	defaultReadTimeout       = 5 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleTimeout       = 15 * time.Second
	defaultReadHeaderTimeout = 2 * time.Second
)

// AppConfig holds the application's configuration parameters.
type AppConfig struct {
	CORS       cors.Options
	APIBaseURL string
	Server     ServerConfig
}

// ServerConfig holds HTTP server timeout settings.
type ServerConfig struct {
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	Port              int
}

// AppConfigBuilder is a builder for constructing AppConfig instances.
type AppConfigBuilder struct {
	config AppConfig
}

// NewAppConfigBuilder creates a new AppConfigBuilder with default configuration values.
func NewAppConfigBuilder() *AppConfigBuilder {
	return &AppConfigBuilder{
		config: AppConfig{
			CORS: cors.Options{
				AllowedOrigins:   []string{"http://localhost:5173"},
				AllowedMethods:   []string{"GET", "POST"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
				ExposedHeaders:   []string{},
				AllowCredentials: true,
				MaxAge:           defaultPortMaxAge,
				Debug:            false,
			},
			APIBaseURL: "https://api.green-api.com",
			Server: ServerConfig{
				ReadTimeout:       defaultReadTimeout,
				WriteTimeout:      defaultWriteTimeout,
				IdleTimeout:       defaultIdleTimeout,
				ReadHeaderTimeout: defaultReadHeaderTimeout,
				Port:              defaultPort,
			},
		},
	}
}

// WithPortFromEnv sets the application port from an environment variable.
// It parses the environment variable value and validates it as a port number.
func (b *AppConfigBuilder) WithPortFromEnv(key string) *AppConfigBuilder {
	if p := os.Getenv(key); p != "" {
		if port, err := strconv.Atoi(p); err == nil && port > 0 && port <= 65535 {
			b.config.Server.Port = port
		}
	}

	return b
}

// Build finalizes the AppConfig construction and returns the configured AppConfig.
func (b *AppConfigBuilder) Build() *AppConfig {
	return &b.config
}
