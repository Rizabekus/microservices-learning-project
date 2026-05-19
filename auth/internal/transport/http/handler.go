package http

import (
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/service"
)

type Handler struct {
	Service service.Service
}

func New(service service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userCreateDTO UserCreateDTO
	err := json.NewDecoder(r.Body).Decode(&userCreateDTO)
	if err != nil {
		logger.Log.Error("failed to decode request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	accessToken, refreshToken, err := h.Service.Register(
		ctx,
		userCreateDTO.FirstName,
		userCreateDTO.LastName,
		userCreateDTO.Email,
		userCreateDTO.Password,
		userCreateDTO.MobileNumber,
	)
	if err != nil {
		handleError(w, err)
		return
	}
	response := RegisterResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Log.Error("failed to encode response", "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dto LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		logger.Log.Error("failed to decode request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	access, refresh, err := h.Service.Login(ctx, dto.Email, dto.Password)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(LoginResponseDTO{
		AccessToken:  access,
		RefreshToken: refresh,
	})
	if err != nil {
		logger.Log.Error("failed to encode response", "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dto RefreshDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		logger.Log.Error("failed to decode request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	access, err := h.Service.Refresh(ctx, dto.RefreshToken)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(RefreshResponseDTO{
		AccessToken: access,
	})
	if err != nil {
		logger.Log.Error("failed to encode response", "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
