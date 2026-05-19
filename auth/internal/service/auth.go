package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Rizabekus/microservices-learning-project/auth/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/logger"
	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateSession(ctx context.Context, session *domain.Session) error
	GetSessionByToken(ctx context.Context, token string) (*domain.Session, error)
}
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}
type TokenManager interface {
	GenerateToken(userID uuid.UUID, tokenType string, duration time.Duration) (string, error)
}
type Service interface {
	Register(ctx context.Context, FirstName string, LastName string, Email string, Password string, MobileNumber string) (string, string, error)
	Login(ctx context.Context, Email string, Password string) (string, string, error)
	Refresh(ctx context.Context, RefreshToken string) (string, error)
}

type service struct {
	Repository     Repository
	PasswordHasher PasswordHasher
	TokenManager   TokenManager
}

func New(repository Repository, passwordHasher PasswordHasher, tokenManager TokenManager) Service {
	return &service{
		Repository:     repository,
		PasswordHasher: passwordHasher,
		TokenManager:   tokenManager,
	}
}

func (s *service) Register(ctx context.Context, firstName string, lastName string, email string, password string, mobileNumber string) (string, string, error) {
	userExists, err := s.Repository.GetUserByEmail(ctx, email)

	if err != nil {
		logger.Log.Error("failed to check if user exists", "error", err)
		return "", "", err
	}

	if userExists != nil {
		return "", "", ErrUserAlreadyExists
	}
	user, err := domain.NewUser(firstName, lastName, email, password, mobileNumber)
	if err != nil {
		return "", "", err
	}
	passwordHash, err := s.PasswordHasher.Hash(password)
	if err != nil {
		logger.Log.Error("failed to hash password", "error", err)
		return "", "", err
	}
	user.PasswordHash = passwordHash
	err = s.Repository.CreateUser(ctx, user)
	if err != nil {
		logger.Log.Error("failed to create user", "error", err)
		return "", "", err
	}
	accessToken, err := s.TokenManager.GenerateToken(user.ID, "access", time.Hour)
	if err != nil {
		logger.Log.Error("failed to generate access token", "error", err)
		return "", "", err
	}
	refreshToken, err := s.TokenManager.GenerateToken(user.ID, "refresh", 24*time.Hour)
	if err != nil {
		logger.Log.Error("failed to generate refresh token", "error", err)
		return "", "", err
	}
	err = s.Repository.CreateSession(ctx, &domain.Session{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		logger.Log.Error("failed to create session", "error", err)
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (s *service) Login(ctx context.Context, email string, password string) (string, string, error) {
	user, err := s.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", ErrInvalidCredentials
		}
		logger.Log.Error("failed to get user by email", "error", err)
		return "", "", err
	}

	err = s.PasswordHasher.Compare(user.PasswordHash, password)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	accessToken, err := s.TokenManager.GenerateToken(user.ID, "access", time.Hour)
	if err != nil {
		logger.Log.Error("failed to generate access token", "error", err)
		return "", "", err
	}

	refreshToken, err := s.TokenManager.GenerateToken(user.ID, "refresh", 24*time.Hour)
	if err != nil {
		logger.Log.Error("failed to generate refresh token", "error", err)
		return "", "", err
	}

	err = s.Repository.CreateSession(ctx, &domain.Session{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		logger.Log.Error("failed to create session", "error", err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (string, error) {
	session, err := s.Repository.GetSessionByToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrInvalidToken
		}
		logger.Log.Error("failed to get session by token", "error", err)
		return "", err
	}
	if session.IsExpired() || session.IsRevoked() {
		return "", ErrInvalidToken
	}
	accessToken, err := s.TokenManager.GenerateToken(session.UserID, "access", time.Hour)
	if err != nil {
		logger.Log.Error("failed to generate access token", "error", err)
		return "", err
	}

	return accessToken, nil
}
