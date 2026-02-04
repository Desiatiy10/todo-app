package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"-" db:"password_hash"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
