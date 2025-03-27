package config

import (
	"fmt"

	"github.com/SlashNephy/auto-claimer/database"
	"github.com/SlashNephy/auto-claimer/pipeline"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/do/v2"
)

type Config struct {
	Database database.DatabaseConfig `envPrefix:"DATABASE_"`
	Batch    pipeline.BatchRedeemCodesConfig
}

func LoadConfig(do.Injector) (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return &config, nil
}
