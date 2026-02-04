package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Desiatiy10/todo-app/internal/handler"
	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/internal/service"
	"github.com/Desiatiy10/todo-app/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("failed initializing configs: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Warn("env file not found, using sytem env")
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
		logrus.Fatalf("failed to initialize db: %v", err)
	}

	signingKey := viper.GetString("signingKey.key")
	if signingKey == "" {
		logrus.Fatal("signing key is not set")
	}

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo, viper.GetString("signingKey.key"), viper.GetDuration("auth.tokenttl"))
	handler := handler.NewHandler(srvc)

	server := new(server.Server)

	go func() {
		logrus.Infof("Starting server on %s", viper.GetString("port-app"))
		if err := server.Run(viper.GetString("port-app"), handler.InitRoutes()); err != nil {
			logrus.Infof("error running http server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.Info("Shutting down server...")

	if err := server.Shutdown(ctx); err != nil {
		logrus.Errorf("server shutdown failed: %v", err)
	}

	logrus.Info("Server exited properly")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
