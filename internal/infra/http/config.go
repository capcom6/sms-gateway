package http

var ConfigDefault = Config{
	Listen: ":4000",
}

type Config struct {
	Listen string
}

// Helper function to set default values
func configDefault(config Config) Config {
	// Override default config
	if config.Listen == "" {
		config.Listen = ConfigDefault.Listen
	}

	return config
}
