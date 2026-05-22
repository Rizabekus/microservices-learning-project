package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Rizabekus/microservices-learning-project/order/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/order/internal/service"
)

type APIError struct {
	Message string `json:"error"`
}

func MapError(err error) (int, string) {
	switch {
	// validation (domain)
	case errors.Is(err, domain.ErrInvalidUserID):
		return http.StatusBadRequest, "invalid user id"

	case errors.Is(err, domain.ErrInvalidAmount):
		return http.StatusBadRequest, "invalid amount"

	case errors.Is(err, domain.ErrInvalidCurrency):
		return http.StatusBadRequest, "invalid currency"

	// business logic (service)
	case errors.Is(err, service.ErrOrderNotFound):
		return http.StatusNotFound, "order not found"

	case errors.Is(err, service.ErrForbidden):
		return http.StatusForbidden, "order does not belong to user"

	case errors.Is(err, service.ErrOrderAlreadyCancelled):
		return http.StatusConflict, "order already cancelled"

	case errors.Is(err, service.ErrCannotCancelPaidOrder):
		return http.StatusConflict, "cannot cancel paid order"

	default:
		return http.StatusInternalServerError, "internal server error"
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
