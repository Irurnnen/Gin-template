package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type HelloRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

type HelloRepositoryInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloRepository(db *sqlx.DB, logger *zap.Logger) *HelloRepository {
	return &HelloRepository{
		db:     db,
		logger: logger,
	}
}

func (r *HelloRepository) GetHelloMessage() (string, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.db.Get(&message, query)
	if err != nil {
		zap.Error(err)
		return "", err
	}
	return message, nil
}
