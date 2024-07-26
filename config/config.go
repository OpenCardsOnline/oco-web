package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	DBConnectionString string
	APIBaseURL         string
	EmailAPIKey        string
}

func (_c *AppConfig) LoadEnvVars() {
	envErrorMessages := []string{}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		envErrorMessages = append(envErrorMessages, "DB_CONNECTION_STRING")
	}

	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		envErrorMessages = append(envErrorMessages, "API_BASE_URL")
	}

	emailAPIKey := os.Getenv("EMAIL_API_KEY")
	if emailAPIKey == "" {
		envErrorMessages = append(envErrorMessages, "EMAIL_API_KEY")
	}

	if len(envErrorMessages) > 0 {
		panic(fmt.Sprintf("Missing required ENV vars: %v", envErrorMessages))
	}

	_c.DBConnectionString = dbConnectionString
	_c.APIBaseURL = apiBaseURL
	_c.EmailAPIKey = emailAPIKey
}
