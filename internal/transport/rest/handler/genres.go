package handler

import (
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetGenresMovies(c echo.Context, params apiserver.GetGenresMoviesParams) error {
	requestParams := tmdb.GetAllMovieGenresListParams{
		Language: params.Language,
	}
	r, err := h.tmdbClient.GetAllMovieGenresList(c.Request().Context(), &requestParams)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetAllMovieGenresListResponse(r)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
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

	res := apiserver.Genres{
		Genres: genres,
	}

	return c.JSON(http.StatusOK, res)
}
