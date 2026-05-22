package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Amount    int64
	Currency  string
	Status    OrderStatus
	CreatedAt time.Time
}

func NewOrder(userID uuid.UUID, amount int64, currency string) (*Order, error) {
	if userID == uuid.Nil {
		return nil, ErrInvalidUserID
	}

	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	//temporary
	if currency == "" {
		return nil, ErrInvalidCurrency
	}
	return &Order{
		ID:        uuid.New(),
		UserID:    userID,
		Amount:    amount,
		Currency:  currency,
		Status:    OrderStatusPending,
		CreatedAt: time.Now().UTC(),
	}, nil
}
