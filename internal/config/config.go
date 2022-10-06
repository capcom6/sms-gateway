package config

import (
	"errors"
	"io/fs"
	"os"
	"sync"

	microbase "bitbucket.org/soft-c/gomicrobase"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP     microbase.HTTPConfig     `yaml:"http"`
	Database microbase.DatabaseConfig `yaml:"database"`
}

var instance *Config
var once = sync.Once{}

func newConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			errorLog.Println(err)
		}
	}

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config.yml"
	}

	config := Config{}

	if err := microbase.LoadConfig(path, &config); err != nil {
		errorLog.Fatalf("Can't load config from %s: %s", path, err.Error())
	}

	return &config
}

func GetConfig() Config {
	once.Do(func() {
		instance = newConfig()
	})

	return *instance
}
