package config

import (
	shared "github.com/capcom6/sms-gateway/internal/shared/config"
)

type Config struct {
	HTTP     shared.HTTP     `yaml:"http"`
	Database shared.Database `yaml:"database"`
}
