package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// Test case where environment variable exists
	key := "TEST_ENV_VAR"
	expectedValue := "expectedValue"
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key) // Ensure the environment variable is cleared after the test

	value := getEnv(key, "defaultValue")
	if value != expectedValue {
		t.Errorf("getEnv(%q, %q) = %q; want %q", key, "defaultValue", value, expectedValue)
	}

	// Test case where environment variable does not exist
	defaultValue := "defaultValue"
	value = getEnv("NON_EXISTENT_ENV_VAR", defaultValue)
	if value != defaultValue {
		t.Errorf("getEnv(%q, %q) = %q; want %q", "NON_EXISTENT_ENV_VAR", defaultValue, value, defaultValue)
	}
}

func TestLoadConfig(t *testing.T) {
	// Set mock environment variables
	baseURLKey := "SWAPI_BASE_URL"
	baseURLValue := "https://mockapi.example.com"
	portKey := "APP_PORT"
	portValue := "9090"

	os.Setenv(baseURLKey, baseURLValue)
	os.Setenv(portKey, portValue)
	defer os.Unsetenv(baseURLKey)
	defer os.Unsetenv(portKey)

	config := LoadConfig()

	if config.BaseURL != baseURLValue {
		t.Errorf("LoadConfig().BaseURL = %q; want %q", config.BaseURL, baseURLValue)
	}

	if config.Port != portValue {
		t.Errorf("LoadConfig().Port = %q; want %q", config.Port, portValue)
	}

	// Test with default values (unset environment variables)
	os.Unsetenv(baseURLKey)
	os.Unsetenv(portKey)

	config = LoadConfig()

	if config.BaseURL != "https://swapi.dev/api" {
		t.Errorf("LoadConfig().BaseURL = %q; want %q", config.BaseURL, "https://swapi.dev/api")
	}

	if config.Port != "8080" {
		t.Errorf("LoadConfig().Port = %q; want %q", config.Port, "8080")
	}
}
