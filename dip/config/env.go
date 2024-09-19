package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Env struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

func LoadEnv() (*Env, error) {
	var ev Env
	if err := env.Parse(&ev); err != nil {
		log.Fatal("Error loading environment")
		return nil, err
	}

	return &ev, nil
}
