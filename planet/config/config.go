package config

import (
	"os"
	"strconv"
)

type Config struct {
	Mongodb_connetion_string string
	Rest_max_retry           int
	Rest_wait_sec            int
	Rest_max_wait_sec        int
}

// InitConfig get all desired Env Variables and load
func (c *Config) InitConfig() {
	c.Mongodb_connetion_string = c.GetEnvAsStringOrFallback("MONGODB_CONNECTION_URL", "mongodb://localhost:27017")
	c.Rest_max_retry = c.GetEnvAsIntOrFallback("REST_CLIENT_RETRY_MAX", 2)
	c.Rest_wait_sec = c.GetEnvAsIntOrFallback("REST_CLIENT_WAIT_SECS", 2)
	c.Rest_max_wait_sec = c.GetEnvAsIntOrFallback("REST_CLIENT_MAX_WAIT_SECS", 10)
}

// GetEnvAsStringOrFallback returns the env variable for the given key
// and falls back to the given defaultValue if not set
func (c *Config) GetEnvAsStringOrFallback(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// GetEnvAsIntOrFallback returns the env variable (parsed as integer) for
// the given key and falls back to the given defaultValue if not set
func (c *Config) GetEnvAsIntOrFallback(key string, defaultValue int) int {
	if v := os.Getenv(key); v != "" {
		value, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue
		}
		return value
	}
	return defaultValue
}
