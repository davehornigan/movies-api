package handler

import (
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetMoviesIdDetails(c echo.Context, id apiserver.MovieId, params apiserver.GetMoviesIdDetailsParams) error {
	requestParams := tmdb.GetMovieDetailsParams{
		Language: params.Language,
	}
	r, err := h.tmdbClient.GetMovieDetails(c.Request().Context(), id, &requestParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetMovieDetailsResponse(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}

	if response.StatusCode() == http.StatusNotFound {
		return c.JSON(http.StatusNotFound, "Movie not found")
	}

	if response.StatusCode() != http.StatusOK {
		return c.JSON(response.StatusCode(), "External error")
	}

	movieDetails := apiserver.MovieDetailsResponse{
		Budget:      *response.JSON200.Budget,
		Companies:   nil,
		Countries:   nil,
		Duration:    *response.JSON200.Runtime,
		Genres:      nil,
		Id:          id,
		IsAdult:     *response.JSON200.Adult,
		Overview:    *response.JSON200.Overview,
		Popularity:  float64(*response.JSON200.Popularity),
		Poster:      response.JSON200.PosterPath,
		Rating:      float64(*response.JSON200.VoteAverage),
		ReleaseDate: nil,
		Revenue:     *response.JSON200.Revenue,
		Status:      *response.JSON200.Status,
		Tagline:     *response.JSON200.Tagline,
		Title:       *response.JSON200.Title,
	}

	return c.JSON(http.StatusOK, movieDetails)
}
