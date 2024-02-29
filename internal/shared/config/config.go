package config

type HTTP struct {
	Listen string `yaml:"listen" envconfig:"HTTP__LISTEN"`
}

type Database struct {
	Dialect  string `yaml:"dialect"  envconfig:"DATABASE__DIALECT"`
	Host     string `yaml:"host"     envconfig:"DATABASE__HOST"`
	Port     int    `yaml:"port"     envconfig:"DATABASE__PORT"`
	User     string `yaml:"user"     envconfig:"DATABASE__USER"`
	Password string `yaml:"password" envconfig:"DATABASE__PASSWORD"`
	Database string `yaml:"database" envconfig:"DATABASE__DATABASE"`
	Timezone string `yaml:"timezone" envconfig:"DATABASE__TIMEZONE"`
}
