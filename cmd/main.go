package main

import (
	"context"
	api "github.com/davehornigan/movies-api"
	"github.com/davehornigan/movies-api/pkg/handler"
	"github.com/gin-gonic/gin"
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
	gin.SetMode(config.Environment)

	handlers := new(handler.Handler)

	server := new(api.Server)
	go func() {
		if err := server.Run(config.ServerPort, handlers.InitRoutes()); err != nil {
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
