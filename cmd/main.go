package main

import (
	"log"

	"github.com/Desiatiy10/todo-app/cmd/server"
	"github.com/Desiatiy10/todo-app/internal/handler"
)

var port string = ":8080"

func main() {
	server := new(server.Server)

	handlers := new(handler.Handler)

	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %v", err)
	}

	log.Printf("Starting server on %s", port)

}
