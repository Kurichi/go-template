package config

import "os"

type Config struct {
	ServerPort string
}

func New() *Config {
	return &Config{
		ServerPort: LookupEnv("PORT", true, "8080"),
	}
}

func LookupEnv(key string, isPanic bool, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if isPanic {
		panic("env " + key + " not found")
	}
	return defaultValue
}
