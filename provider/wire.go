//go:build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitServer() *echo.Echo {
	wire.Build(
		InitEcho,
	)

	return nil
}
