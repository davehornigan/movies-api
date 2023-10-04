package main

import (
	"context"
	api "github.com/davehornigan/movies-api"
	"github.com/davehornigan/movies-api/oas"
	"github.com/davehornigan/movies-api/pkg/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := api.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %s", err.Error())
	}

	handlers := new(handler.Handler)

	server := new(api.Server)
	router := handlers.InitRoutes()
	oas.RegisterHandlersWithBaseURL(router, handlers, "/api")
	go func() {
		if err := server.Run(config.ServerPort, router); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Printf("Server started on port: %s", config.ServerPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on http server shutdown: %s", err.Error())
	}
}
