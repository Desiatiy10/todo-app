package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Desiatiy10/todo-app/errs"
	"github.com/Desiatiy10/todo-app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(user models.User) (uuid.UUID, error) {
	var id uuid.UUID

	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	const query = `INSERT INTO users (id, name, username, password_hash) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`

	row := p.db.QueryRow(query, user.ID, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (p *AuthPostgres) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	const query = `SELECT id, name, username, password_hash 
		FROM users 
		WHERE username = $1`

	err := p.db.Get(&user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%w: username=%s", errs.ErrUserNotFound, username)
		}
		return models.User{}, fmt.Errorf("database error: %w", err)
	}

	return user, nil
}
