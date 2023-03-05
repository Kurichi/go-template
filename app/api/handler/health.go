package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthzHandler struct {
	Message string `json:"message"`
}

func Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, &HealthzHandler{
		Message: "ok",
	})
}
