package middleware

import (
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var URL = c.Request.URL
		headers := c.Request.Header
		headersMap := make(map[string]string)
		for k, v := range headers {
			headersMap[k] = strings.Join(v, ",")
		}
		c.Next()
		timeSpend := time.Since(start)
		timeSpendMs := float64(timeSpend.Microseconds()) / 1000.0
		logger.Info("request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", URL.Path),
			zap.String("query", URL.RawQuery),
			zap.Any("headers", headersMap),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Float64("time_spend_ms", timeSpendMs),
		)
	}
}

// ZapRecovery 恢复中间件
func ZapRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fields := []zap.Field{
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
					zap.String("ip", c.ClientIP()),
				}

				// 根据 stack 参数决定是否记录堆栈信息
				if stack {
					fields = append(fields, zap.String("stack", string(debug.Stack())))
				}

				logger.Error("panic recovered", fields...)
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
