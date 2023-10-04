package main

import (
	api "github.com/davehornigan/movies-api"
	"github.com/davehornigan/movies-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(api.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
