package redis

import (
	"context"
	"fmt"
	"go-server-starter/internal/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	*redis.Client
	logger *zap.Logger
}

func NewRedis(config config.RedisConfig, logger *zap.Logger, ctx context.Context) (*Redis, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("connect to redis failed: %w", err)
	}
	return &Redis{db, logger}, nil
}
