package service

import (
	"time"

	"github.com/Desiatiy10/todo-app/errs"
	"github.com/Desiatiy10/todo-app/internal/repository"
	"github.com/Desiatiy10/todo-app/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

type AuthService struct {
	repo       repository.Authorization
	signingKey string
	tokenTTL   time.Duration
}

func NewAuthService(repo repository.Authorization, signingKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		repo:       repo,
		signingKey: signingKey,
		tokenTTL:   tokenTTL,
	}
}

func (s *AuthService) SignUp(input models.SignUpInput) (uuid.UUID, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return uuid.Nil, errs.ErrFailedToHashPassword
	}

	user := models.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Username: input.Username,
		Password: string(passwordHash),
	}

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return uuid.Nil, errs.ErrFailedToCreateUser
	}

	return id, nil
}

func (s *AuthService) SignIn(input models.SignInInput) (string, error) {
	user, err := s.repo.GetUserByUsername(input.Username)
	if err != nil {
		return "", errs.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errs.ErrInvalidCredentials
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", errs.ErrFailedToGenerateToken
	}

	return token, nil
}

func (s *AuthService) generateToken(userID uuid.UUID) (string, error) {
	claims := &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.signingKey))
	if err != nil {
		return "", errs.ErrFailedToSignToken
	}

	return signedToken, nil
}
