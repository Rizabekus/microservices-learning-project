package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
	RevokedAt *time.Time
}

func NewSession(userID uuid.UUID, token string, expiresAt time.Time) *Session {
	return &Session{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now().UTC(),
	}
}
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

func (s *Session) IsRevoked() bool {
	return s.RevokedAt != nil
}
