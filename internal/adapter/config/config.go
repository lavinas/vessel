package config

import (
	"os"
)

// Config represents the configuration for the service
type Config struct {
}

// NewConfig creates a new Config
func NewConfig() *Config {
	return &Config{}
}

// Get is a method that gets the configuration value
func (c *Config) Get(key string) string {
	return os.Getenv(key)
}
