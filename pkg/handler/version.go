package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) version(c echo.Context) error {
	return c.String(http.StatusOK, "Version")
}
