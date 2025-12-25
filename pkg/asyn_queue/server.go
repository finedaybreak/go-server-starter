package asyn_queue

import (
	"context"
	"fmt"
	"go-server-starter/internal/config"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Concurrency int            // 并发数
	Queues      map[string]int // 队列优先级
}

type Server struct {
	*asynq.Server
	mux    *asynq.ServeMux
	logger *zap.Logger
}

func NewServer(redisConfig config.AsynQConfig, serverConfig ServerConfig, logger *zap.Logger) *Server {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", redisConfig.RedisConfig.Host, redisConfig.RedisConfig.Port),
			Password: redisConfig.RedisConfig.Password,
			DB:       redisConfig.RedisConfig.DB,
		},
		asynq.Config{
			Concurrency: serverConfig.Concurrency,
			Queues:      serverConfig.Queues,
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				logger.Error("task processing failed",
					zap.String("type", task.Type()),
					zap.ByteString("payload", task.Payload()),
					zap.Error(err),
				)
			}),
		},
	)
	return &Server{Server: srv, mux: asynq.NewServeMux(), logger: logger}
}

// Handle 注册任务处理器
func (s *Server) Handle(pattern string, handler asynq.Handler) {
	s.mux.Handle(pattern, handler)
}

// HandleFunc 注册任务处理函数
func (s *Server) HandleFunc(pattern string, handler func(context.Context, *asynq.Task) error) {
	s.mux.HandleFunc(pattern, handler)
}

// Start 启动服务器
func (s *Server) Start() error {
	// s.logger.Info("starting asynq worker server...")
	return s.Server.Start(s.mux)
}

// Shutdown 关闭服务器
func (s *Server) Shutdown() {
	s.Server.Shutdown()
	s.logger.Info("asynq worker server stopped")
}
