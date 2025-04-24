package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const DriverName = "pgx"

// Repository wraps the database connection and provides methods to interact with it.
type Repository struct {
	DB *sqlx.DB
}

// NewRepository initializes a new Repository instance with the given sqlx.DB connection.
func NewRepository(DSN string) (*Repository, error) {
	// TODO: check DSN

	db, err := sqlx.Connect(DriverName, DSN)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: db,
	}, nil

}

func (r *Repository) Ping() error {
	return r.DB.Ping()
}

// Close closes the database connection.
func (r *Repository) Close() error {
	return r.DB.Close()
}
