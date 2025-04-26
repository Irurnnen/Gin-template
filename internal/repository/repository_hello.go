package repository

import (
	"github.com/Irurnnen/gin-template/internal/database"
	"go.uber.org/zap"
)

type HelloRepository struct {
	provider *database.Provider
	logger   *zap.Logger
}

type HelloRepositoryInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloRepository(provider *database.Provider, logger *zap.Logger) *HelloRepository {
	return &HelloRepository{
		provider: provider,
		logger:   logger,
	}
}

func (r *HelloRepository) GetHelloMessage() (string, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.provider.GetDB().Get(&message, query)
	if err != nil {
		zap.Error(err)
		return "", err
	}
	return message, nil
}
