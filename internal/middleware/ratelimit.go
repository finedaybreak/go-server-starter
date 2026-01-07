package middleware

import (
	"fmt"
	"go-server-starter/internal/constant"
	"go-server-starter/internal/ctx"
	"go-server-starter/internal/exception"
	"go-server-starter/pkg/redis"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"go.uber.org/zap"
)

type RateLimit struct {
	redis  *redis.Redis
	logger *zap.Logger
}

func NewRateLimit(redis *redis.Redis, logger *zap.Logger) *RateLimit {
	return &RateLimit{redis: redis, logger: logger}
}

// RateLimit 限流中间件（每分钟限流 rate 次）
func (r *RateLimit) RateLimit(rate int, zones ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := ctx.FromGinCtx(c)
		ip := c.ClientIP()

		if len(zones) == 0 {
			zones = []string{"GLOBAL"}
		}

		zoneStr := strings.ToUpper(strings.Join(zones, ":"))

		// Create Redis key using IP as identifier
		key := constant.RedisKeyOfRateLimit(zoneStr, ip)

		limiter := redis_rate.NewLimiter(r.redis)
		res, err := limiter.Allow(ctx.Ctx, key, redis_rate.PerMinute(rate))
		if err != nil {
			r.logger.Error("failed to allow rate limit", zap.Error(err))
			c.Next()
			return
		}

		headerKey := fmt.Sprintf("X-RateLimit-%s-Remaining", zoneStr)
		c.Header(headerKey, strconv.FormatInt(int64(res.Remaining), 10))

		if res.Allowed == 0 {
			c.Header("Retry-After", strconv.FormatFloat(res.RetryAfter.Seconds(), 'f', 0, 64))
			ctx.ToError(exception.TooManyRequests)
			return
		}
		c.Next()
	}
}
