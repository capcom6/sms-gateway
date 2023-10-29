package config

import (
	"sync"

	microbase "bitbucket.org/soft-c/gomicrobase"
)

type Config struct {
	HTTP     microbase.HTTPConfig     `yaml:"http"`
	Database microbase.DatabaseConfig `yaml:"database"`
	FCM      FCMConfig                `yaml:"fcm"`
}

var instance *Config
var once = sync.Once{}

func newConfig() *Config {
	config := Config{}

	if err := microbase.LoadConfig(&config); err != nil {
		errorLog.Fatalf("Can't load config from %s", err.Error())
	}

	return &config
}

func GetConfig() Config {
	once.Do(func() {
		instance = newConfig()
	})

	return *instance
}
