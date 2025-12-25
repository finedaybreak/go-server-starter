package middleware

import (
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
}

func DefaultCORSConfig() CORSConfig {
	// 允许客户端发送的请求头
	allowHeaders := []string{
		// 标准 HTTP headers
		"Content-Type",
		"Content-Length",
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		"Authorization",
		"Cache-Control",
		"X-Requested-With",
		"X-Request-ID",
		"Locale",
	}

	exposeHeaders := []string{
		"X-Request-ID",
		"new-token",
	}

	return CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     allowHeaders,
		ExposeHeaders:    exposeHeaders,
		AllowCredentials: false,
		MaxAge:           172800, // 48小时
	}
}

func CORSWithConfig(config CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var allowOrigin = false

		origin := c.Request.Header.Get("Origin")

		if len(config.AllowOrigins) > 0 {
			if slices.Contains(config.AllowOrigins, "*") {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				allowOrigin = false
			} else if slices.Contains(config.AllowOrigins, origin) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				allowOrigin = true
			} else {
				allowOrigin = false
			}
		}

		if config.AllowCredentials && allowOrigin {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if len(config.AllowHeaders) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
		}

		if len(config.AllowMethods) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))
		}

		if len(config.ExposeHeaders) > 0 {
			c.Writer.Header().Set("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ", "))
		}

		if config.MaxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.MaxAge))
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return CORSWithConfig(DefaultCORSConfig())
}
