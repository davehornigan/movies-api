package handler

import (
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)
import apiserver "github.com/davehornigan/movies-api/generated/api-server"

func (h *Handler) GetMoviesListType(c echo.Context, listType apiserver.MovieListType, params apiserver.GetMoviesListTypeParams) error {
	logrus.Info(listType)
	logrus.Info(*params.Page)
	logrus.Info(*params.Language)
	logrus.Info(*params.CountryCode)
	var response *http.Response
	var err error
	if listType == apiserver.Popular {
		requestParams := tmdb.GetMoviePopularPaginatedParams{
			Page:     params.Page,
			Language: params.Language,
			Region:   params.CountryCode,
		}
		response, err = h.tmdbClient.GetMoviePopularPaginated(c.Request().Context(), &requestParams)
	} else if listType == apiserver.Upcoming {
		requestParams := tmdb.GetMovieUpcomingPaginatedParams{
			Page:     params.Page,
			Language: params.Language,
			Region:   params.CountryCode,
		}
		response, err = h.tmdbClient.GetMovieUpcomingPaginated(c.Request().Context(), &requestParams)
	} else {
		requestParams := tmdb.GetMovieTopRatedPaginatedParams{
			Page:     params.Page,
			Language: params.Language,
			Region:   params.CountryCode,
		}
		response, err = h.tmdbClient.GetMovieTopRatedPaginated(c.Request().Context(), &requestParams)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	logrus.Info(response)

	return c.String(http.StatusOK, "MovieList")
}
