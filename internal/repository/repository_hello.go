package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HelloRepository struct {
	dbPool *pgxpool.Pool
	logger *zap.Logger
}

type HelloRepositoryInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloRepository(dbPool *pgxpool.Pool, logger *zap.Logger) *HelloRepository {
	return &HelloRepository{
		dbPool: dbPool,
		logger: logger,
	}
}

func (r *HelloRepository) GetHelloMessage() (string, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.dbPool.QueryRow(context.Background(), query).Scan(&message)
	if err != nil {
		zap.Error(err)
		return "", err
	}
	return message, nil
}
