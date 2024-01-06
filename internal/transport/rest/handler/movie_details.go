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

	genres := make([]apiserver.Genre, 0)
	for _, genre := range *response.JSON200.Genres {
		genres = append(genres, apiserver.Genre{
			Id:   *genre.Id,
			Name: *genre.Name,
		})
	}
	countries := make([]apiserver.Country, 0)
	for _, country := range *response.JSON200.ProductionCountries {
		countries = append(countries, apiserver.Country{
			Code: country.Iso6391,
			Name: country.Name,
		})
	}
	companies := make([]apiserver.Company, 0)
	for _, company := range *response.JSON200.ProductionCompanies {
		companies = append(companies, apiserver.Company{
			Id:   company.Id,
			Name: company.Name,
			Logo: company.LogoPath,
		})
	}

	movieDetails := apiserver.MovieDetailsResponse{
		Budget:      *response.JSON200.Budget,
		Companies:   companies,
		Countries:   countries,
		Duration:    *response.JSON200.Runtime,
		Genres:      genres,
		Id:          id,
		IsAdult:     *response.JSON200.Adult,
		Overview:    *response.JSON200.Overview,
		Popularity:  *response.JSON200.Popularity,
		Poster:      response.JSON200.PosterPath,
		Rating:      *response.JSON200.VoteAverage,
		ReleaseDate: response.JSON200.ReleaseDate,
		Revenue:     *response.JSON200.Revenue,
		Status:      *response.JSON200.Status,
		Tagline:     *response.JSON200.Tagline,
		Title:       *response.JSON200.Title,
	}

	return c.JSON(http.StatusOK, movieDetails)
}
