package config

import (
	"fmt"
	"log/slog"
	"os"

	"gitlab.com/kdg-ti/the-lab/teams-25-26/26-de-uitgeruste-it-ers/algorithm/internal/metrics/env"
)

type Config struct {
	Port   string
	APIKey string
}

func LoadConfig() Config {
	port := env.Get("SERVER_PORT", "8081")

	apiKey, ok := env.GetOptional("API_KEY")
	if !ok {
		slog.Error("API_KEY is required but not set")
		os.Exit(1)
	}

	return Config{
		Port:   fmt.Sprintf(":%s", port),
		APIKey: apiKey,
	}
}
