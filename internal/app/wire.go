//go:build wireinject

package app

import (
	"strconv"

	"github.com/SAMPO-BU/sampo-server/internal/app/config"
	"github.com/SAMPO-BU/sampo-server/internal/app/container"
	"github.com/SAMPO-BU/sampo-server/pkg/database"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func New() (*container.App, error) {
	wire.Build(
		echo.New,
		config.New,
		config.NewDBConfig,
		provideDBConfig,
		container.NewApp,
		database.New,
	)

	return nil, nil
}

func provideDBConfig(cfg *config.DBConfig) (*database.Config, error) {
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		return nil, err
	}

	return &database.Config{
		Host:     cfg.Host,
		Port:     uint(port),
		User:     cfg.User,
		Password: cfg.Password,
		DBName:   cfg.DBName,
	}, nil
}
