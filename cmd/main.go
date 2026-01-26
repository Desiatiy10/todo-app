package main

import (
	"log"

	"github.com/Desiatiy10/todo-app/internal/handler"
	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/internal/service"
	"github.com/Desiatiy10/todo-app/server"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %v", err)
	}

	repo := repository.NewRepository()
	srvc := service.NewService(repo)
	handler := handler.NewHandler(srvc)
	server := new(server.Server)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %v", err)
	}

	log.Printf("Starting server on %s", viper.GetString("port"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
