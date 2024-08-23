package container

import (
	"errors"
	"fmt"

	"github.com/Kurichi/go-template/internal/app/config"
	"github.com/Kurichi/go-template/pkg/database"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
	cfg  *config.Config
	db   *database.DB
}

func NewApp(e *echo.Echo, cfg *config.Config, db *database.DB) *App {
	return &App{
		echo: e,
		cfg:  cfg,
		db:   db,
	}
}

func (a *App) Run() (err error) {

	return a.echo.Start(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}
