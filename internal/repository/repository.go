package repository

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

const DriverName = "pgx"

type Repository struct {
	// db              *sqlx.DB
	dbPool          *pgxpool.Pool
	logger          *zap.Logger
	HelloRepository HelloRepositoryInterface
}

func NewRepository(DSN string, logger *zap.Logger) (*Repository, error) {
	dbPool, err := pgxpool.New(context.Background(), DSN)
	// db, err := sqlx.Connect(DriverName, DSN)
	if err != nil {
		return nil, err
	}

	return &Repository{
		// db:              db,
		dbPool:          dbPool,
		logger:          logger,
		HelloRepository: NewHelloRepository(dbPool, logger),
	}, nil
}

func NewRepositoryDB(dbPool *pgxpool.Pool, logger *zap.Logger) *Repository {
	return &Repository{
		// db:              db,
		dbPool:          dbPool,
		logger:          logger,
		HelloRepository: NewHelloRepository(dbPool, logger),
	}
}

func (r *Repository) Ping() error {
	return r.dbPool.Ping(context.Background())
}

func (r *Repository) Close() {
	r.dbPool.Close()
}
