package handler

import (
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"net/http"
	"time"
)

func (h *Handler) GetMoviesIdVideos(c echo.Context, id apiserver.MovieId, params apiserver.GetMoviesIdVideosParams) error {
	requestParams := tmdb.GetMovieVideosListParams{
		Language: params.Language,
	}
	r, err := h.tmdbClient.GetMovieVideosList(c.Request().Context(), id, &requestParams)

	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}
	response, err := tmdb.ParseGetMovieVideosListResponse(r)

	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error()) // TODO: Customize error
	}

	if response.StatusCode() == http.StatusNotFound {
		logrus.Error(err.Error())
		return c.JSON(http.StatusNotFound, "Movie not found")
	}

	if response.StatusCode() != http.StatusOK {
		logrus.Error(err.Error())
		return c.JSON(response.StatusCode(), "External error")
	}

	typeFilter := *params.Types
	if len(typeFilter) == 0 {
		typeFilter = append(typeFilter, apiserver.Trailer)
	}

	videos := make([]apiserver.MovieVideo, 0)
	for _, video := range *response.JSON200.Results {
		videoType := apiserver.VideoType(*video.Type)

		if !slices.Contains(typeFilter, videoType) || *video.Site != "YouTube" {
			continue
		}

		PublicationTimeParsed, err := time.Parse(time.RFC3339, *video.PublishedAt)
		var publicationDate *types.Date
		if err == nil {
			publicationDate = &types.Date{Time: PublicationTimeParsed}
		}

		videos = append(videos, apiserver.MovieVideo{
			Name:            *video.Name,
			Key:             *video.Key,
			Type:            videoType,
			Size:            *video.Size,
			IsOfficial:      *video.Official,
			PublicationDate: publicationDate,
		})
	}

	movieVideos := apiserver.MovieVideosResponse{
		Id:      id,
		Results: videos,
	}

	return c.JSON(http.StatusOK, movieVideos)
}
