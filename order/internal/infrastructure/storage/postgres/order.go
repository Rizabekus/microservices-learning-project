package postgres

import (
	"context"

	"github.com/Rizabekus/microservices-learning-project/order/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/storage/postgres/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) CreateOrder(ctx context.Context, order *domain.Order) error {
	err := r.queries.CreateOrder(ctx,
		db.CreateOrderParams{
			ID: pgtype.UUID{
				Bytes: order.ID,
				Valid: true,
			},

			UserID: pgtype.UUID{
				Bytes: order.UserID,
				Valid: true,
			},
			Amount:   order.Amount,
			Currency: order.Currency,
			Status:   string(order.Status),
			CreatedAt: pgtype.Timestamptz{
				Time:  order.CreatedAt,
				Valid: true,
			},
		},
	)
	return err
}

func (r *Repository) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	order, err := r.queries.GetOrderByID(ctx, pgtype.UUID{
		Bytes: orderID,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}
	return &domain.Order{
		ID:        order.ID.Bytes,
		UserID:    order.UserID.Bytes,
		Amount:    order.Amount,
		Currency:  order.Currency,
		Status:    domain.OrderStatus(order.Status),
		CreatedAt: order.CreatedAt.Time,
	}, nil
}

func (r *Repository) GetOrdersByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Order, error) {
	orders, err := r.queries.GetOrdersByUserID(ctx, pgtype.UUID{
		Bytes: userID,
		Valid: true,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*domain.Order, len(orders))
	for i, order := range orders {
		result[i] = &domain.Order{
			ID:        order.ID.Bytes,
			UserID:    order.UserID.Bytes,
			Amount:    order.Amount,
			Currency:  order.Currency,
			Status:    domain.OrderStatus(order.Status),
			CreatedAt: order.CreatedAt.Time,
		}
	}
	return result, nil
}

func (r *Repository) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status domain.OrderStatus) error {
	return r.queries.UpdateOrderStatus(ctx, db.UpdateOrderStatusParams{
		ID:     pgtype.UUID{Bytes: orderID, Valid: true},
		Status: string(status),
	})
}
