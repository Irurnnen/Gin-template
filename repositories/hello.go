package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type HelloRepository struct {
	dbPool *pgxpool.Pool
	logger *zerolog.Logger
}

type HelloRepositoryInterface interface {
	GetHelloMessage() (*HelloEntity, error)
}

func NewHelloRepository(dbPool *pgxpool.Pool, logger *zerolog.Logger) *HelloRepository {
	return &HelloRepository{
		dbPool: dbPool,
		logger: logger,
	}
}

func (r *HelloRepository) GetHelloMessage() (*HelloEntity, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.dbPool.QueryRow(context.Background(), query).Scan(&message)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get hello message")
		return nil, err
	}
	return &HelloEntity{Message: message}, nil
}
