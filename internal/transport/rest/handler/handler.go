package handler

import (
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	tmdbClient tmdb.ClientInterface
}

func NewHandler(tmdbClient tmdb.ClientInterface) *Handler {
	return &Handler{tmdbClient: tmdbClient}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()
	log := logrus.New()
	router.Use(middleware.CORS())
	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"Method": values.Method,
				"URI":    values.URI,
				"Status": values.Status,
				"Error":  values.Error,
			}).Info("request")

			return nil
		},
	}))

	router.GET("/version", h.version)
	router.Static("/", "./web/static")

	return router
}
