package main

import (
	"context"
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	"github.com/davehornigan/movies-api/generated/tmdb"
	appConfig "github.com/davehornigan/movies-api/internal/app-config"
	"github.com/davehornigan/movies-api/internal/transport/rest"
	"github.com/davehornigan/movies-api/internal/transport/rest/handler"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := appConfig.LoadConfig(".env")
	if err != nil {
		logrus.Fatalf("cannot load app-config: %s", err.Error())
	}

	tokenProvider, tokenProviderErr := securityprovider.NewSecurityProviderBearerToken(config.TmdbAccessToken)
	if tokenProviderErr != nil {
		panic(tokenProviderErr)
	}

	tmdbClient, _ := tmdb.NewClient(config.TmdbServer, tmdb.WithRequestEditorFn(tokenProvider.Intercept))
	handlers := handler.NewHandler(tmdbClient)

	server := new(rest.Server)
	router := handlers.InitRoutes()
	apiserver.RegisterHandlersWithBaseURL(router, handlers, "/api")
	go func() {
		if err := server.Run(config.ServerPort, router); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Infof("Server started on port: %s. Environment: %s", config.ServerPort, config.Environment)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Server shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occurred on http server shutdown: %s", err.Error())
	}
}
