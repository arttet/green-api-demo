package config

import (
	"os"
	"strconv"

	"github.com/rs/cors"
)

type AppConfig struct {
	Port       int
	APIBaseURL string
	CORSConfig cors.Options
}

type AppConfigBuilder struct {
	config AppConfig
}

func NewAppConfigBuilder() *AppConfigBuilder {
	return &AppConfigBuilder{
		config: AppConfig{
			Port:       8080,
			APIBaseURL: "https://api.green-api.com",
			CORSConfig: cors.Options{
				AllowedOrigins:   []string{"http://localhost:5173"},
				AllowedMethods:   []string{"GET", "POST"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
				ExposedHeaders:   []string{},
				AllowCredentials: true,
				MaxAge:           300,
				Debug:            true,
			},
		},
	}
}

func (b *AppConfigBuilder) WithPortFromEnv(key string) *AppConfigBuilder {
	if p := os.Getenv(key); p != "" {
		if port, err := strconv.Atoi(p); err == nil && port > 0 && port <= 65535 {
			b.config.Port = port
		}
	}

	return b
}

func (b *AppConfigBuilder) Build() *AppConfig {
	return &b.config
}
