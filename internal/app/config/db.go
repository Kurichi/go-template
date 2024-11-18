package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     readStr("DB_HOST", "localhost"),
		Port:     readStr("DB_PORT", "5432"),
		User:     readStr("DB_USER", "postgres"),
		Password: readStr("DB_PASSWORD", "password"),
		DBName:   readStr("DB_NAME", "postgres"),
	}
}

func readStr(key, value string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return value
}

func readInt(key string, value int) int {
	s, ok := os.LookupEnv(key)
	if !ok {
		return value
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return value
	}

	return v
}
