package service

import (
	"fmt"

	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(input models.User) (int, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: string(passwordHash),
	}

	return s.repo.CreateUser(user)
}
