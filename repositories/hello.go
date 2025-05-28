package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type HelloRepository struct {
	dbPool *pgxpool.Pool
	logger *zerolog.Logger
	redis  *redis.Client
}

type HelloRepositoryInterface interface {
	GetHelloMessage(context.Context) (*HelloEntity, error)
	GetHelloMessageWithCache(context.Context) (*HelloEntity, error)
}

func NewHelloRepository(dbPool *pgxpool.Pool, logger *zerolog.Logger, redis *redis.Client) *HelloRepository {
	return &HelloRepository{
		dbPool: dbPool,
		redis:  redis,
		logger: logger,
	}
}

func (r *HelloRepository) GetHelloMessage(ctx context.Context) (*HelloEntity, error) {
	var message string
	query := "SELECT 'Hello World' AS message"
	err := r.dbPool.QueryRow(ctx, query).Scan(&message)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get hello message")
		return nil, err
	}
	return &HelloEntity{Message: message}, nil
}

func (r *HelloRepository) GetHelloMessageWithCache(ctx context.Context) (*HelloEntity, error) {
	// Try to get value from cache
	val, err := r.redis.Get(ctx, "hello_message").Result()
	if err == nil {
		r.logger.Debug().Msg("Cache hit for hello_message")
		return &HelloEntity{Message: val}, nil
	}

	// If cache is empty - get from DB
	var message string
	query := "SELECT 'Hello World' AS message"
	err = r.dbPool.QueryRow(ctx, query).Scan(&message)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get hello message from DB")
		return nil, err
	}

	// Save to cache
	err = r.redis.Set(ctx, "hello_message", message, 5*time.Minute).Err()
	if err != nil {
		r.logger.Warn().Err(err).Msg("Failed to write to cache hello message")
	}
	r.logger.Debug().Msg("Cache set for hello_message")
	return &HelloEntity{Message: message}, nil
}
