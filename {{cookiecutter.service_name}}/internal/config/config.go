package config

import (
	"os"
)

type Config struct {
	Port        string
	Environment string
	ServiceName string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "{{cookiecutter.server_port}}"),
		Environment: getEnv("ENVIRONMENT", "development"),
		ServiceName: "{{cookiecutter.service_name}}",
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
