package redis

import (
	"context"

	"github.com/exceptionteapots/gin-template/config"
	"github.com/redis/go-redis/v9"
)

// NewClient creates new Redis client by config
func NewClient(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:       cfg.Address,
		Password:   cfg.Password,
		DB:         cfg.DB,
		Username:   cfg.User,
		MaxRetries: cfg.MaxRetries,
		// DialTimeout:  cfg.DialTimeout,
		// ReadTimeout:  cfg.Timeout,
		// WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return db, nil
}
