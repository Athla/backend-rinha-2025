package config

import "os"

// Holds all env values
type Config struct {
	ConversionRate      int
	SqlConnString       string
	SqlDriver           string
	FallbackServiceAddr string
	DefaultServiceAddr  string
}

func New() *Config {
	c := new(Config)
	c.load()

	return c
}

func (c *Config) load() {
	c.SqlDriver = "sqlite3"
	c.SqlConnString = "payments.db"
	c.ConversionRate = 100
	c.DefaultServiceAddr = os.Getenv("PAYMENT_PROCESSOR_DEFAULT_ADDR")
	c.FallbackServiceAddr = os.Getenv("PAYMENT_PROCESSOR_FALLBACK_ADDR")
}
