package config

import (
	"os"
)

// Config holds application configuration.
type Config struct {
	BaseURL string
	Port    string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() *Config {
	baseURL := getEnv("SWAPI_BASE_URL", "https://swapi.dev/api")
	port := getEnv("APP_PORT", "8080")

	return &Config{
		BaseURL: baseURL,
		Port:    port,
	}
}

// getEnv fetches an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
