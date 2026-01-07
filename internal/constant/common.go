package constant

import (
	"fmt"
	"time"
)

const (
	MAX_PAGE_SIZE           = 500                    // max page size
	DEFAULT_PAGE_SIZE       = 20                     // default page size
	TTLAuthRoles            = 5 * time.Minute        // ttl of auth roles
	REDIS_KEY_OF_RATE_LIMIT = "api:rate_limit:%s:%s" // redis key of rate limit: zone:ip
	REDIS_KEY_OF_AUTH_ROLES = "auth:roles:%s"        // redis key of auth roles: uniCode
)

func RedisKeyOfRateLimit(zone string, ip string) string {
	return fmt.Sprintf(REDIS_KEY_OF_RATE_LIMIT, zone, ip)
}

func RedisKeyOfAuthRoles(uniCode string) string {
	return fmt.Sprintf(REDIS_KEY_OF_AUTH_ROLES, uniCode)
}
