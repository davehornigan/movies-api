package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)
import apiserver "github.com/davehornigan/movies-api/generated/api-server"

func (h *Handler) GetMoviesListType(c echo.Context, listType apiserver.MovieListType, params apiserver.GetMoviesListTypeParams) error {
	return c.String(http.StatusOK, "MovieList")
}
