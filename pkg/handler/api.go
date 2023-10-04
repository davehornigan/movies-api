package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)
import apiserver "github.com/davehornigan/movies-api/generated/api-server"

func (h *Handler) GetMoviesIdDetails(c echo.Context, id apiserver.MovieId, params apiserver.GetMoviesIdDetailsParams) error {
	return c.String(http.StatusOK, "MovieDetails")
}
func (h *Handler) GetMoviesListType(c echo.Context, listType apiserver.MovieListType, params apiserver.GetMoviesListTypeParams) error {
	return c.String(http.StatusOK, "MovieList")
}
