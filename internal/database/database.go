package database

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const DriverName = "pgx"

// Provider represents database connection provider
type Provider struct {
	db *sqlx.DB
}

// NewProvider creates new database provider
func NewProvider(dsn string) (*Provider, error) {
	db, err := sqlx.Connect(DriverName, dsn)
	if err != nil {
		return nil, err
	}

	return &Provider{db: db}, nil
}

// GetDB returns database connection
func (p *Provider) GetDB() *sqlx.DB {
	return p.db
}

// Close closes database connection
func (p *Provider) Close() error {
	return p.db.Close()
}

// Ping checks database connection
func (p *Provider) Ping() error {
	return p.db.Ping()
}
