package handler

import (
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"net/http"
	"time"
)
import apiserver "github.com/davehornigan/movies-api/generated/api-server"

func (h *Handler) GetMoviesListType(c echo.Context, listType apiserver.MovieListType, params apiserver.GetMoviesListTypeParams) error {
	var getMovies func(h *Handler, c echo.Context, params apiserver.GetMoviesListTypeParams) error
	if listType == apiserver.Popular {
		getMovies = GetPopularMovies
	} else if listType == apiserver.Upcoming {
		getMovies = GetUpcomingMovies
	} else if listType == apiserver.NowPlaying {
		getMovies = GetNowPlayingMovies
	} else {
		getMovies = GetTopRatedMovies
	}

	return getMovies(h, c, params)
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
	response, err := tmdb.ParseGetMovieTopRatedPaginatedResponse(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}

	if response.StatusCode() != http.StatusOK {
		return c.JSON(response.StatusCode(), "External error")
	}

	return c.JSON(http.StatusOK, BuildMovieList(response.JSON200))
}

func GetNowPlayingMovies(h *Handler, c echo.Context, params apiserver.GetMoviesListTypeParams) error {
	requestParams := tmdb.GetMovieNowPlayingPaginatedParams{
		Page:     params.Page,
		Language: params.Language,
		Region:   params.CountryCode,
	}
	r, err := h.tmdbClient.GetMovieNowPlayingPaginated(c.Request().Context(), &requestParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetMovieNowPlayingPaginatedResponse(r)

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
		releaseDateParsed, err := time.Parse(time.DateOnly, *movie.ReleaseDate)
		var releaseDate *types.Date
		if err == nil {
			releaseDate = &types.Date{Time: releaseDateParsed}
		}

		movieList = append(movieList, apiserver.MovieShortResponse{
			Id:          *movie.Id,
			IsAdult:     *movie.Adult,
			Poster:      movie.PosterPath,
			Backdrop:    movie.BackdropPath,
			GenreIds:    movie.GenreIds,
			Rating:      *movie.VoteAverage,
			ReleaseDate: releaseDate,
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
