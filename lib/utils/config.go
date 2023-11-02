package utils

import "os"

type AppConfig struct {
	DeepInfraApiKey string
	AppName         string
}

func NewConfig() *AppConfig {
	return &AppConfig{
		DeepInfraApiKey: os.Getenv("DEEPINFRA_API_KEY"),
		AppName:         os.Getenv("APP_NAME"),
	}
}
