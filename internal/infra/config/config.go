package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
}

func New(params Param) any {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		params.Logger.Error("Error loading .env file", zap.Error(err))
	}

	configPath := "config.yml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		params.Logger.Error("Error reading config file", zap.Error(err))
	}

	err = yaml.Unmarshal(yamlFile, params.Config)
	if err != nil {
		params.Logger.Error("Error unmarshalling config file", zap.Error(err))
	}

	return params.Config
}
