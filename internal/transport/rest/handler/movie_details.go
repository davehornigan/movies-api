package handler

import (
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetMoviesIdDetails(c echo.Context, id apiserver.MovieId, params apiserver.GetMoviesIdDetailsParams) error {
	return c.String(http.StatusOK, "MovieDetails")
}
