package repository

import (
	"github.com/Irurnnen/gin-template/internal/database"
)

type Repository struct {
	provider        *database.Provider
	HelloRepository HelloRepositoryInterface
}

func NewRepository(provider *database.Provider) *Repository {
	return &Repository{
		provider:        provider,
		HelloRepository: NewHelloRepository(provider),
	}
}

func (r *Repository) Ping() error {
	return r.provider.Ping()
}
