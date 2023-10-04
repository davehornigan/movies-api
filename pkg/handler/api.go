package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)
import oas "github.com/davehornigan/movies-api/oas"

func (h *Handler) GetMoviesIdDetails(c echo.Context, id oas.MovieId, params oas.GetMoviesIdDetailsParams) error {
	return c.String(http.StatusOK, "MovieDetails")
}
func (h *Handler) GetMoviesListType(c echo.Context, listType oas.MovieListType, params oas.GetMoviesListTypeParams) error {
	return c.String(http.StatusOK, "MovieList")
}
