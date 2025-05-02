package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type HelloRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

type HelloRepositoryInterface interface {
	GetHelloMessage(ctx context.Context) (string, error)
}

func NewHelloRepository(db *sqlx.DB, logger *zap.Logger) *HelloRepository {
	return &HelloRepository{
		db:     db,
		logger: logger,
	}
}

func (r *HelloRepository) GetHelloMessage(ctx context.Context) (string, error) {
	ctx, span := otel.Tracer("repository").Start(ctx, "GetHelloMessage")
	defer span.End()

	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.db.GetContext(ctx, &message, query) // Используем GetContext для передачи контекста
	if err != nil {
		zap.Error(err)
		return "", err
	}
	return message, nil
}
