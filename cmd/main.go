package main

import (
	"log"

	"github.com/Desiatiy10/todo-app/internal/handler"
	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/internal/service"
	"github.com/Desiatiy10/todo-app/server"
)

var port string = ":8080"

func main() {
	repo := repository.NewRepository()
	srvc := service.NewService(repo)
	handler := handler.NewHandler(srvc)
	server := new(server.Server)
	
	if err := server.Run(port, handler.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %v", err)
	}

	log.Printf("Starting server on %s", port)

}
