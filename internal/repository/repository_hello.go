package repository

import "github.com/jmoiron/sqlx"

type HelloRepository struct {
	db *sqlx.DB
}

type HelloRepositoryInterface interface {
	GetHelloMessage() (string, error)
}

func NewHelloRepository(db *sqlx.DB) *HelloRepository {
	return &HelloRepository{
		db: db,
	}
}

func (r *HelloRepository) GetHelloMessage() (string, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.db.Get(&message, query)
	if err != nil {
		return "", err
	}
	return message, nil
}
