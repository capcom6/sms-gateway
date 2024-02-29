package config

import (
	shared "github.com/capcom6/sms-gateway/internal/shared/config"
)

type Config struct {
	HTTP     shared.HTTP     `yaml:"http"`
	Database shared.Database `yaml:"database"`
	FCM      FCMConfig       `yaml:"fcm"`
	Tasks    Tasks           `yaml:"tasks"`
}

type FCMConfig struct {
	CredentialsJSON string `yaml:"credentials_json"`
	DebounceSeconds uint16 `yaml:"debounce_seconds"`
	TimeoutSeconds  uint16 `yaml:"timeout_seconds"`
}

type Tasks struct {
	Hashing HashingTask `yaml:"hashing"`
}

type HashingTask struct {
	IntervalSeconds uint16 `yaml:"interval_seconds"`
}

var defaultConfig = Config{
	HTTP: shared.HTTP{
		Listen: ":3000",
	},
	Database: shared.Database{
		Dialect:  "mysql",
		Host:     "localhost",
		Port:     3306,
		User:     "sms",
		Password: "sms",
		Database: "sms",
		Timezone: "UTC",
	},
	FCM: FCMConfig{
		CredentialsJSON: "",
	},
	Tasks: Tasks{
		Hashing: HashingTask{
			IntervalSeconds: uint16(15 * 60),
		},
	},
}
