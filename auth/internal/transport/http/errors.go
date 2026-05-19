package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Rizabekus/microservices-learning-project/auth/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/service"
)

type APIError struct {
	Message string `json:"error"`
}

func MapError(err error) (int, string) {
	switch {
	// validation (domain)
	case errors.Is(err, domain.ErrInvalidFirstName):
		return 400, "invalid first name"

	case errors.Is(err, domain.ErrInvalidLastName):
		return 400, "invalid last name"

	case errors.Is(err, domain.ErrInvalidEmail):
		return 400, "invalid email"

	case errors.Is(err, domain.ErrInvalidPassword):
		return 400, "invalid password"

	case errors.Is(err, domain.ErrInvalidMobileNumber):
		return 400, "invalid mobile number"

	// business logic (service)
	case errors.Is(err, service.ErrUserAlreadyExists):
		return 409, "user already exists"

	case errors.Is(err, service.ErrInvalidCredentials):
		return 401, "invalid credentials"

	case errors.Is(err, service.ErrSessionNotFound):
		return 404, "session not found"

	// fallback
	default:
		return 500, "internal server error"
	}
}

func handleError(w http.ResponseWriter, err error) {
	code, msg := MapError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(APIError{Message: msg})
	if err != nil {
		logger.Log.Error("failed to encode error response: %v", slog.String("err", err.Error()))
		http.Error(w, "failed to encode error response", http.StatusInternalServerError)
		return
	}
}
