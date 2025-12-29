package config

import (
	"time"
)

var DefaultConfig = Config{
	Server: ServerConfig{
		Port:          8080,
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
		MaxHeaderKB:   2048,
		SnowflakeNode: 1,
		APIPrefix:     "/api",
	},
	JWT: JWTConfig{
		Issuer:      "go-server-starter",
		TokenSecret: "tokenSecret@123",
		TokenExpires: MultiTokenExpireConfig{
			Web:             24 * time.Hour,
			Desktop:         15 * 24 * time.Hour,
			Mobile:          15 * 24 * time.Hour,
			ChromeExtension: 30 * 24 * time.Hour,
			API:             2 * 24 * time.Hour,
			Default:         1 * 24 * time.Hour,
		},
	},
	Database: DatabaseConfig{
		Host:            "localhost",
		Port:            3306,
		Username:        "root",
		Password:        "root",
		Name:            "test_db",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 10 * time.Second,
		Timezone:        "UTC",
		Charset:         "utf8mb4",
		ParseTime:       true,
	},
	Redis: RedisConfig{
		Host:     "localhost",
		Port:     6379,
		Password: "root",
		DB:       0,
	},
	AsynQ: AsynQConfig{
		RedisConfig: RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "root",
			DB:       1,
		},
		Concurrency: 10,
	},
	Logger: LoggerConfig{
		Level:         "info",
		FileDir:       "./logs",
		MaxSize:       100,
		MaxAge:        30,
		MaxBackups:    10,
		Compress:      false,
		ConsoleOutput: true,
	},
	GormLogger: GormLoggerConfig{
		Level:                     "info",
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true,
	},
}
