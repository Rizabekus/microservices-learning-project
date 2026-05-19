package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTTokenManager struct {
	Secret string
}

func NewJWTTokenManager(secret string) *JWTTokenManager {
	return &JWTTokenManager{
		Secret: secret,
	}
}

func (m *JWTTokenManager) GenerateToken(userID uuid.UUID, tokenType string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"type":    tokenType,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.Secret))
}
