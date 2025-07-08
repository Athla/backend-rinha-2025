package config

// Holds all env values
type Config struct {
	ConversionRate int
}

func New() *Config {
	return &Config{
		ConversionRate: 100,
	}
}
