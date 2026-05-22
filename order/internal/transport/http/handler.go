package http

import (
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/order/internal/service"
	"github.com/Rizabekus/microservices-learning-project/order/internal/transport/http/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	Service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dto CreateOrderDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		logger.Log.Error("failed to decode request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userID, ok := middlewares.UserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("failed to extract user ID from context")
		http.Error(w, "Failed to extract user ID", http.StatusUnauthorized)
		return
	}
	err = h.Service.CreateOrder(ctx, userID, dto.Amount, dto.Currency)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orderID := chi.URLParam(r, "id")
	if orderID == "" {
		logger.Log.Error("missing order id in request")
		http.Error(w, "Missing order id", http.StatusBadRequest)
		return
	}
	userID, ok := middlewares.UserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("failed to extract user ID from context")
		http.Error(w, "Failed to extract user ID", http.StatusUnauthorized)
		return
	}
	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		logger.Log.Error("invalid order id format", "error", err)
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}
	order, err := h.Service.GetOrder(ctx, userID, orderUUID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		logger.Log.Error("failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetMyOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := middlewares.UserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("failed to extract user ID from context")
		http.Error(w, "Failed to extract user ID", http.StatusUnauthorized)
		return
	}
	orders, err := h.Service.ListUserOrders(ctx, userID)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		logger.Log.Error("failed to encode response", "error", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orderID := chi.URLParam(r, "id")
	if orderID == "" {
		logger.Log.Error("missing order id in request")
		http.Error(w, "Missing order id", http.StatusBadRequest)
		return
	}
	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		logger.Log.Error("invalid order id format", "error", err)
		http.Error(w, "Invalid order id", http.StatusBadRequest)
		return
	}
	userID, ok := middlewares.UserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("failed to extract user ID from context")
		http.Error(w, "Failed to extract user ID", http.StatusUnauthorized)
		return
	}
	err = h.Service.CancelOrder(ctx, userID, orderUUID)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
