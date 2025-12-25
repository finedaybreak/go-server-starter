package database

import (
	"context"
	"errors"
	"fmt"
	"go-server-starter/internal/config"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql db: %w", err)
	}
	return sqlDB.Close()
}

func NewDB(config config.DatabaseConfig, logger logger.Interface, gormConfig *gorm.Config) (*DB, error) {
	// Validate connection pool parameters
	if config.MaxIdleConns > config.MaxOpenConns {
		return nil, errors.New("MaxIdleConns cannot be greater than MaxOpenConns")
	}
	if config.MaxOpenConns <= 0 {
		return nil, errors.New("MaxOpenConns must be greater than 0")
	}

	var dsn = GetMySQLDSN(config)

	dialector := mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})

	GConfig := gormConfig
	if GConfig == nil {
		GConfig = &gorm.Config{}
	}
	GConfig.Logger = logger

	db, err := gorm.Open(dialector, GConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w with dsn: %s", err, dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	return &DB{db}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}

func GetMySQLDSN(config config.DatabaseConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)

	queryParams := url.Values{}
	queryParams.Add("charset", config.Charset)
	if config.ParseTime {
		queryParams.Add("parseTime", "True")
	} else {
		queryParams.Add("parseTime", "False")
	}
	queryParams.Add("loc", config.Timezone)

	return fmt.Sprintf("%s?%s", dsn, queryParams.Encode())
}
