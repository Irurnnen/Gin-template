package repository

import (
	"github.com/Irurnnen/gin-template/internal/database"
)

type HelloRepository struct {
	provider *database.Provider
}

type HelloRepositoryInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloRepository(provider *database.Provider) *HelloRepository {
	return &HelloRepository{
		provider: provider,
	}
}

func (r *HelloRepository) GetHelloMessage() (string, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.provider.GetDB().Get(&message, query)
	if err != nil {
		return "", err
	}
	return message, nil
}
