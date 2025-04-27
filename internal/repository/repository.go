package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const DriverName = "pgx"

type Repository struct {
	db              *sqlx.DB
	logger          *zap.Logger
	HelloRepository HelloRepositoryInterface
}

func NewRepository(DSN string, logger *zap.Logger) (*Repository, error) {
	db, err := sqlx.Connect(DriverName, DSN)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:              db,
		logger:          logger,
		HelloRepository: NewHelloRepository(db, logger),
	}, nil
}

func (r *Repository) Ping() error {
	return r.db.Ping()
}

func (r *Repository) Close() error {
	return r.db.Close()
}
