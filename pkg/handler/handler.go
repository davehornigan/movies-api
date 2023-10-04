package handler

import "github.com/labstack/echo/v4"

type Handler struct {
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	api := router.Group("/api")
	{
		api.GET("/version", h.version)
	}

	return router
}
