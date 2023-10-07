package handler

import (
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"net/http"
)
import apiserver "github.com/davehornigan/movies-api/generated/api-server"

func (h *Handler) GetMoviesListType(c echo.Context, listType apiserver.MovieListType, params apiserver.GetMoviesListTypeParams) error {
	if listType == apiserver.Popular {
		return GetPopularMovies(h, c, params)
	} else if listType == apiserver.Upcoming {
		return GetUpcomingMovies(h, c, params)
	}

	return GetTopRatedMovies(h, c, params)
}

func GetPopularMovies(h *Handler, c echo.Context, params apiserver.GetMoviesListTypeParams) error {
	requestParams := tmdb.GetMoviePopularPaginatedParams{
		Page:     params.Page,
		Language: params.Language,
		Region:   params.CountryCode,
	}
	r, err := h.tmdbClient.GetMoviePopularPaginated(c.Request().Context(), &requestParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetMoviePopularPaginatedResponse(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}

	if response.StatusCode() != http.StatusOK {
		return c.JSON(response.StatusCode(), "External error")
	}

	return c.JSON(http.StatusOK, BuildMovieList(response.JSON200))
}

func GetUpcomingMovies(h *Handler, c echo.Context, params apiserver.GetMoviesListTypeParams) error {
	requestParams := tmdb.GetMovieUpcomingPaginatedParams{
		Page:     params.Page,
		Language: params.Language,
		Region:   params.CountryCode,
	}
	r, err := h.tmdbClient.GetMovieUpcomingPaginated(c.Request().Context(), &requestParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetMovieUpcomingPaginatedResponse(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}

	if response.StatusCode() != http.StatusOK {
		return c.JSON(response.StatusCode(), "External error")
	}

	return c.JSON(http.StatusOK, BuildMovieList(response.JSON200))
}

func GetTopRatedMovies(h *Handler, c echo.Context, params apiserver.GetMoviesListTypeParams) error {
	requestParams := tmdb.GetMovieTopRatedPaginatedParams{
		Page:     params.Page,
		Language: params.Language,
		Region:   params.CountryCode,
	}
	r, err := h.tmdbClient.GetMovieTopRatedPaginated(c.Request().Context(), &requestParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetMovieUpcomingPaginatedResponse(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}

	if response.StatusCode() != http.StatusOK {
		return c.JSON(response.StatusCode(), "External error")
	}

	return c.JSON(http.StatusOK, BuildMovieList(response.JSON200))
}

func BuildMovieList(resp *tmdb.MoviePaginated) apiserver.MovieListResponse {
	movieList := make([]apiserver.MovieShortResponse, 0)
	for _, movie := range *resp.Results {
		movieList = append(movieList, apiserver.MovieShortResponse{
			Duration:    0,
			Id:          *movie.Id,
			IsAdult:     *movie.Adult,
			Poster:      movie.PosterPath,
			Rating:      float64(*movie.Popularity),
			ReleaseDate: nil,
			Title:       *movie.Title,
		})
	}
	listResponse := apiserver.MovieListResponse{
		Items:      movieList,
		Page:       *resp.Page,
		TotalItems: *resp.TotalResults,
		TotalPages: *resp.TotalPages,
	}

	return listResponse
}
