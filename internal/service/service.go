package service

import (
	"time"

	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/models"
	"github.com/google/uuid"
)

type Authorization interface {
	SignUp(input models.SignUpInput) (uuid.UUID, error)
	SignIn(input models.SignInInput) (string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
}

type TodoList interface {
	Create(userID uuid.UUID, list models.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository, signingKey string, tockenTTL time.Duration) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization, signingKey, tockenTTL),
		TodoList:      NewTodoListService(repo),
	}
}
