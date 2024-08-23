//go:build wireinject

package app

import (
	"github.com/Kurichi/go-template/internal/app/config"
	"github.com/Kurichi/go-template/internal/app/container"
	"github.com/Kurichi/go-template/pkg/database"
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

func provideDBConfig(cfg *config.DBConfig) *database.Config {
	return &database.Config{}
}
