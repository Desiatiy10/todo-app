package service

import (
	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/models"
	"github.com/google/uuid"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(UserID uuid.UUID, list models.TodoList) (int, error) {
	return s.repo.Create(UserID, list)
}
