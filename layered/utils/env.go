package utils

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

func LoadEnv() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Error loading environment")
		return nil, err
	}

	return &cfg, nil
}
