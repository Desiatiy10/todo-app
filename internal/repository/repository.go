package repository

import (
	"github.com/Desiatiy10/todo-app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (uuid.UUID, error)
	GetUserByUsername(username string) (models.User, error)
}

type TodoList interface {
	Create(userID uuid.UUID, list models.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
	}
}
