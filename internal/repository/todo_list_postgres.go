package repository

import (
	"github.com/Desiatiy10/todo-app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (p *TodoListPostgres) Create(userID uuid.UUID, list models.TodoList) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	const createListQuery = `INSERT INTO todo_lists (title, description) VALUES ($1, $2) RETURNING id`

	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	const createUsersListQuery = `INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2)`

	_, err = tx.Exec(createUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
