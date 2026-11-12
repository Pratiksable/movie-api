package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func Load() (Config, error) {
	cfg := Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.Port == "" {
		fmt.Println("No port found")
		cfg.Port = "8000"
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE URL IS REQUIRED.....")
	}
	return cfg, nil
}
