package config

type Config struct {
	HTTP     HTTP      `yaml:"http"`
	Database Database  `yaml:"database"`
	FCM      FCMConfig `yaml:"fcm"`
	Tasks    Tasks     `yaml:"tasks"`
}

type HTTP struct {
	Listen string `yaml:"listen" envconfig:"HTTP__LISTEN"`
}

type Database struct {
	Dialect  string `yaml:"dialect" envconfig:"DATABASE__DIALECT"`
	Host     string `yaml:"host" envconfig:"DATABASE__HOST"`
	Port     int    `yaml:"port" envconfig:"DATABASE__PORT"`
	User     string `yaml:"user" envconfig:"DATABASE__USER"`
	Password string `yaml:"password" envconfig:"DATABASE__PASSWORD"`
	Database string `yaml:"database" envconfig:"DATABASE__DATABASE"`
	Timezone string `yaml:"timezone" envconfig:"DATABASE__TIMEZONE"`
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
	HTTP: HTTP{
		Listen: ":3000",
	},
	Database: Database{
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
