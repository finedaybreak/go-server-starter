package redis

import (
	"fmt"
	"time"
)

// ==================== Key 前缀 ====================
const (
	RedisKeyUserRoles = "auth:roles:%s"
)

func KeyAuthRoles(uniCode string) string {
	return fmt.Sprintf(RedisKeyUserRoles, uniCode)
}

// ==================== TTL 配置 ====================
const (
	TTLAuthRoles = 5 * time.Minute
)
