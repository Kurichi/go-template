package router

import (
	"app/api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${method}] ${status} ${path} ${error}` + "\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/healthz", handler.Healthz)

	return e
}
