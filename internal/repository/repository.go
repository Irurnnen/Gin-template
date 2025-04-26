package repository

import (
	"github.com/Irurnnen/gin-template/internal/database"
	"go.uber.org/zap"
)

type Repository struct {
	provider        *database.Provider
	logger          *zap.Logger
	HelloRepository HelloRepositoryInterface
}

func NewRepository(provider *database.Provider, logger *zap.Logger) *Repository {
	return &Repository{
		provider:        provider,
		logger:          logger,
		HelloRepository: NewHelloRepository(provider, logger),
	}
}

func (r *Repository) Ping() error {
	return r.provider.Ping()
}

func (r *Repository) Close() error {
	return r.provider.Close()
}
