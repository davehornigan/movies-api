package handler

import (
	"encoding/json"
	"fmt"
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	"github.com/davehornigan/movies-api/generated/tmdb"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

var acceptedSortingOptions = map[apiserver.SortOption]bool{
	apiserver.Popularity:  true,
	apiserver.Rating:      true,
	apiserver.ReleaseDate: true,
	apiserver.Title:       true,
}

func (h *Handler) GetDiscoverMovies(c echo.Context, params apiserver.GetDiscoverMoviesParams) error {
	requestParams := tmdb.GetDiscoverMoviePaginatedParams{
		SortBy:                getDiscoveryMovieSortBy(params.SortBy),
		CertificationCountry:  params.CertificationCountry,
		Certification:         params.Certification,
		CertificationLte:      params.CertificationLte,
		CertificationGte:      params.CertificationGte,
		IncludeAdult:          params.IncludeAdult,
		IncludeVideo:          params.IncludeVideo,
		Language:              params.Language,
		Page:                  params.Page,
		PrimaryReleaseYear:    params.PrimaryReleaseYear,
		PrimaryReleaseDateGte: params.PrimaryReleaseDateGte,
		PrimaryReleaseDateLte: params.PrimaryReleaseDateLte,
		ReleaseDateGte:        params.ReleaseDateGte,
		ReleaseDateLte:        params.ReleaseDateLte,
		VoteCountGte:          params.VoteCountGte,
		VoteCountLte:          params.VoteCountLte,
		VoteAverageGte:        params.RatingGte,
		VoteAverageLte:        params.RatingLte,
		WithCast:              getFilterStringByParam(params.WithCast, params.WithCastCondition),
		WithCrew:              getFilterStringByParam(params.WithCrew, params.WithCrewCondition),
		WithCompanies:         getFilterStringByParam(params.WithCompanies, params.WithCompaniesCondition),
		WithGenres:            getFilterStringByParam(params.WithGenres, params.WithGenresCondition),
		WithKeywords:          getFilterStringByParam(params.WithKeywords, params.WithKeywordsCondition),
		WithPeople:            getFilterStringByParam(params.WithPeople, params.WithPeopleCondition),
		Year:                  params.Year,
		WithoutGenres:         getFilterString(params.WithoutGenres, apiserver.AND),
		WithRuntimeGte:        params.WithDurationGte,
		WithRuntimeLte:        params.WithDurationLte,
		WithReleaseType:       nil,
		WithOriginalLanguage:  getFilterString(params.WithOriginalLanguage, apiserver.OR),
		WithoutKeywords:       getFilterString(params.WithoutKeywords, apiserver.AND),
		Region:                params.Region,
	}
	r, err := h.tmdbClient.GetDiscoverMoviePaginated(c.Request().Context(), &requestParams)

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
	bytes, _ := json.Marshal(requestParams)
	log.Println(string(bytes))

	return c.JSON(http.StatusOK, BuildMovieList(response.JSON200))
}

func getDiscoveryMovieSortBy(by *apiserver.SortBy) *tmdb.GetDiscoverMoviePaginatedParamsSortBy {
	if by.Option == nil || acceptedSortingOptions[*by.Option] == false {
		return nil
	}
	order := apiserver.Desc
	if by.Order != nil {
		order = *by.Order
	}
	sort := fmt.Sprintf("%s.%s", *by.Option, order)
	result := tmdb.GetDiscoverMoviePaginatedParamsSortBy(sort)

	return &result
}

func getFilterString(values *apiserver.ArrayString, condition apiserver.FilterCondition) *string {
	getDelim := func(c apiserver.FilterCondition) string {
		if c == apiserver.AND {
			return ","
		}

		return "|"
	}

	return joinList(values, getDelim(condition))
}

func getFilterStringByParam(values *apiserver.ArrayString, condition *apiserver.FilterCondition) *string {
	iCondition := apiserver.AND
	if condition != nil {
		iCondition = *condition
	}

	return getFilterString(values, iCondition)
}

func joinList(list *[]string, glue string) *string {
	if list == nil {
		return nil
	}
	res := strings.Join(*list, glue)

	return &res
}
