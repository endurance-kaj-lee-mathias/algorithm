package config

import (
	"fmt"
	"log/slog"
	"os"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/env"
)

type Config struct {
	Port   string
	APIKey string
}

func LoadConfig() Config {
	port := env.Get("SERVER_PORT", "8081")

	apiKey := env.Get("ALGO_API_KEY", "")
	if apiKey == "" {
		slog.Error("ALGO_API_KEY is not set, requests to this service will be rejected")
		os.Exit(1)
	}

	return Config{
		Port:   fmt.Sprintf(":%s", port),
		APIKey: apiKey,
	}
}
