package main

import (
	"log"
	"os"

	"github.com/Desiatiy10/todo-app/internal/handler"
	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/internal/service"
	"github.com/Desiatiy10/todo-app/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed initializing configs: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed initializing env variables: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	repo := repository.NewRepository(db)
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
