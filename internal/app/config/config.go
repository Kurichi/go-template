package config

type Config struct {
	Port uint
}

func New() (*Config, error) {
	return &Config{
		Port: uint(readInt("PORT", 8080)),
	}, nil
}
