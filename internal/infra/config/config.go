package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

func LoadConfig(config any) error {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := loadFromYaml(config); err != nil {
		return err
	}

	if err := loadFromEnv(config); err != nil {
		return err
	}

	return nil
}

func loadFromYaml(config any) error {
	configPath := "config.yml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlFile, config); err != nil {
		return err
	}

	return nil
}

func loadFromEnv(config any) error {
	return envconfig.Process("", config)
}
