package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP     HTTP      `yaml:"http"`
	Database Database  `yaml:"database"`
	FCM      FCMConfig `yaml:"fcm"`
}

type HTTP struct {
	Listen string `yaml:"listen" env:"HTTP_LISTEN"`
}

type Database struct {
	Host     string `yaml:"host" env:"DATABASE_HOST"`
	Port     int    `yaml:"port" env:"DATABASE_PORT"`
	User     string `yaml:"user" env:"DATABASE_USER"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD"`
	Database string `yaml:"database" env:"DATABASE_DATABASE"`
}

type FCMConfig struct {
	CredentialsJSON string `yaml:"credentials_json"`
}

var defaultConfig = Config{
	HTTP: HTTP{
		Listen: ":3000",
	},
	Database: Database{
		Host:     "localhost",
		Port:     3306,
		User:     "sms",
		Password: "sms",
		Database: "sms",
	},
	FCM: FCMConfig{
		CredentialsJSON: "",
	},
}

func GetConfig(logger *zap.Logger) Config {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		logger.Error("Error loading .env file", zap.Error(err))
	}

	configPath := "config.yml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		logger.Error("Error reading config file", zap.Error(err))
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logger.Error("Error unmarshalling config file", zap.Error(err))
	}

	return config
}
