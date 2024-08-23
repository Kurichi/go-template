package config

type Config struct {
	Port uint
}

func New() (*Config, error) {
	return &Config{
		Port: 8080,
	}, nil
}
