package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var AppConfiguration *AppConfig

type AppConfig struct {
	DBConnectionString string
	APIBaseURL         string
	EmailAPIKey        string
}

const (
	dBConnectionString = "DB_CONNECTION_STRING"
	apiBaseUrl         = "API_BASE_URL"
	emailApiKey        = "EMAIL_API_KEY"
)

func LoadEnvVars() (config *AppConfig) {
	envErrorMessages := []string{}

	dbConnectionString := os.Getenv(dBConnectionString)
	if dbConnectionString == "" {
		envErrorMessages = append(envErrorMessages, dBConnectionString)
	}

	apiBaseURL := os.Getenv(apiBaseUrl)
	if apiBaseURL == "" {
		envErrorMessages = append(envErrorMessages, apiBaseUrl)
	}

	emailAPIKey := os.Getenv(emailApiKey)
	if emailAPIKey == "" {
		envErrorMessages = append(envErrorMessages, emailApiKey)
	}

	if len(envErrorMessages) > 0 {
		panic(fmt.Sprintf("Missing required ENV vars: %v", envErrorMessages))
	}

	appConfig := AppConfig{}
	appConfig.DBConnectionString = dbConnectionString
	appConfig.APIBaseURL = apiBaseURL
	appConfig.EmailAPIKey = emailAPIKey

	AppConfiguration = &appConfig
	return AppConfiguration
}
