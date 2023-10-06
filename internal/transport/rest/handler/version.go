package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Version struct {
	Version string `json:"version" xml:"version"`
}

func (h *Handler) version(c echo.Context) error {
	resp := &Version{Version: "1.0"}

	return c.JSON(http.StatusOK, resp)
}
