package postgres

import (
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/storage/postgres/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *db.Queries
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		queries: db.New(pool),
	}
}
