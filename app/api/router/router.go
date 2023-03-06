package router

import (
	"app/api/handler"
	auth "app/api/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339}:[${method}] ${status} ${path} ${error}` + "\n",
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 2,
	}))
	e.Use(middleware.Recover())

	e.GET("/healthz", handler.Healthz)

	g := e.Group("/api")
	g.Use(auth.FirebaseAuthMiddleware("service-account-file.json"))

	g.GET("/", func(c echo.Context) error {
		uid := c.Request().Context().Value("uid").(string)
		return c.JSON(http.StatusOK, map[string]string{"uid": uid})
	})

	return e
}
