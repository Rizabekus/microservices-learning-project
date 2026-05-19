package postgres

import (
	"context"

	"github.com/Rizabekus/microservices-learning-project/auth/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/storage/postgres/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) CreateSession(ctx context.Context, session *domain.Session) error {
	return r.queries.CreateSession(ctx, db.CreateSessionParams{
		ID: pgtype.UUID{
			Bytes: session.ID,
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: session.UserID,
			Valid: true,
		},
		Token: session.Token,

		ExpiresAt: pgtype.Timestamptz{
			Time:  session.ExpiresAt,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamptz{
			Time:  session.CreatedAt,
			Valid: true,
		},

		RevokedAt: pgtype.Timestamptz{
			Valid: false,
		},
	})
}

func (r *Repository) GetSessionByToken(ctx context.Context, token string) (*domain.Session, error) {
	sessionData, err := r.queries.GetSessionByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	session := &domain.Session{
		ID:        sessionData.ID.Bytes,
		UserID:    sessionData.UserID.Bytes,
		Token:     sessionData.Token,
		ExpiresAt: sessionData.ExpiresAt.Time,
		CreatedAt: sessionData.CreatedAt.Time,
	}
	if sessionData.RevokedAt.Valid {
		session.RevokedAt = &sessionData.RevokedAt.Time
	}
	return session, nil
}
