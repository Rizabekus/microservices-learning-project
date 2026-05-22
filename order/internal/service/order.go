package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Rizabekus/microservices-learning-project/order/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/logger"
	"github.com/google/uuid"
)

type Repository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status domain.OrderStatus) error
}
type Service interface {
	CreateOrder(ctx context.Context, userID uuid.UUID, amount int64, currency string) error
	GetOrder(ctx context.Context, userID uuid.UUID, orderID uuid.UUID) (*domain.Order, error)
	ListUserOrders(ctx context.Context, userID uuid.UUID) ([]*domain.Order, error)
	CancelOrder(ctx context.Context, userID uuid.UUID, orderID uuid.UUID) error
}

type service struct {
	Repository Repository
}

func New(repository Repository) Service {
	return &service{
		Repository: repository,
	}
}
func (s *service) CreateOrder(ctx context.Context, userID uuid.UUID, amount int64, currency string) error {

	order, err := domain.NewOrder(userID, amount, currency)
	if err != nil {
		logger.Log.Error("failed to create order", "error", err)
		return err
	}
	err = s.Repository.CreateOrder(ctx, order)
	if err != nil {
		logger.Log.Error("failed to save order", "error", err)
		return err
	}
	return nil
}

func (s *service) GetOrder(ctx context.Context, userID uuid.UUID, orderID uuid.UUID) (*domain.Order, error) {

	order, err := s.Repository.GetOrderByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return nil, ErrOrderNotFound
		}
		logger.Log.Error("failed to get order", "error", err)
		return nil, err
	}
	if order.UserID != userID {
		return nil, ErrForbidden
	}
	return order, nil
}

func (s *service) ListUserOrders(ctx context.Context, userID uuid.UUID) ([]*domain.Order, error) {
	orders, err := s.Repository.GetOrdersByUserID(ctx, userID)
	if err != nil {
		logger.Log.Error("failed to list user orders", "error", err)
		return nil, err
	}
	return orders, nil
}

func (s *service) CancelOrder(ctx context.Context, userID uuid.UUID, orderID uuid.UUID) error {
	order, err := s.Repository.GetOrderByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrOrderNotFound
		}
		logger.Log.Error("failed to get order", "error", err)
		return err
	}
	if order.UserID != userID {
		return ErrForbidden
	}
	if order.Status == domain.OrderStatusCancelled {
		return ErrOrderAlreadyCancelled
	}
	if order.Status == domain.OrderStatusPaid {
		return ErrCannotCancelPaidOrder
	}
	err = s.Repository.UpdateOrderStatus(ctx, orderID, domain.OrderStatusCancelled)

	if err != nil {
		logger.Log.Error("failed to update order status", "error", err)
		return err
	}
	return nil
}
