package postgres

import (
	"context"
	"errors"

	"github.com/Rizabekus/microservices-learning-project/auth/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/storage/postgres/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) error {
	return r.queries.CreateUser(ctx, db.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: user.ID,
			Valid: true,
		},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone: pgtype.Text{
			String: user.MobileNumber,
			Valid:  user.MobileNumber != "",
		},
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt: pgtype.Timestamptz{
			Time:  user.CreatedAt,
			Valid: true,
		},
	})
}
func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	id, err := pgUUIDToUUID(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: user.Phone.String,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.Time,
	}, nil
}
func pgUUIDToUUID(p pgtype.UUID) (uuid.UUID, error) {
	if !p.Valid {
		return uuid.Nil, errors.New("invalid uuid")
	}

	return uuid.FromBytes(p.Bytes[:])
}
