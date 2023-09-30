package provider

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitEcho initializes echo server
func InitEcho() *echo.Echo {
	e := echo.New()
	SetupMiddleware(e)
	SetupCommonRoutes(e)
	return e
}

// SetupMiddleware sets up middleware for echo server
func SetupMiddleware(e *echo.Echo) {
	cors := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	gzip := middleware.GzipConfig{
		Level: 2,
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/health"
		},
	}

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(cors))
	e.Use(middleware.GzipWithConfig(gzip))
}

// SetupCommonRoutes sets up common routes for echo server
func SetupCommonRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}
