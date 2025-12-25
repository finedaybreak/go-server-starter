package asyn_queue

import (
	"context"
	"fmt"
	"go-server-starter/internal/config"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Client struct {
	*asynq.Client
	logger *zap.Logger
}

func NewClient(config config.AsynQConfig, logger *zap.Logger) (*Client, error) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port),
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})
	return &Client{Client: client, logger: logger}, nil
}

// Enqueue 入队任务（便捷方法）
func (c *Client) Enqueue(ctx context.Context, task *asynq.Task, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	info, err := c.Client.EnqueueContext(ctx, task, opts...)
	if err != nil {
		c.logger.Error("failed to enqueue task", zap.String("type", task.Type()), zap.Error(err))
		return nil, err
	}
	c.logger.Info("task enqueued", zap.String("id", info.ID), zap.String("type", task.Type()), zap.String("queue", info.Queue))
	return info, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
