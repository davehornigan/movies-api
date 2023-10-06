package main

import (
	"context"
	apiserver "github.com/davehornigan/movies-api/generated/api-server"
	appConfig "github.com/davehornigan/movies-api/internal/app-config"
	"github.com/davehornigan/movies-api/internal/transport/rest"
	"github.com/davehornigan/movies-api/internal/transport/rest/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := appConfig.LoadConfig(".env")
	if err != nil {
		log.Fatalf("cannot load app-config: %s", err.Error())
	}

	handlers := new(handler.Handler)

	server := new(rest.Server)
	router := handlers.InitRoutes()
	apiserver.RegisterHandlersWithBaseURL(router, handlers, "/api")
	go func() {
		if err := server.Run(config.ServerPort, router); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	router.Logger.Printf("Server started on port: %s. Environment: %s", config.ServerPort, config.Environment)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	router.Logger.Printf("Server shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		router.Logger.Fatalf("error occured on http server shutdown: %s", err.Error())
	}
}
