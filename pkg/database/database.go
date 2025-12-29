package database

import (
	"context"
	"errors"
	"fmt"
	"go-server-starter/internal/config"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
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

func NewDB(cfg config.DatabaseConfig, logger logger.Interface, gormConfig *gorm.Config) (*DB, error) {
	// Validate connection pool parameters
	if cfg.MaxIdleConns > cfg.MaxOpenConns {
		return nil, errors.New("MaxIdleConns cannot be greater than MaxOpenConns")
	}
	if cfg.MaxOpenConns <= 0 {
		return nil, errors.New("MaxOpenConns must be greater than 0")
	}

	loc, _ := time.LoadLocation(cfg.Timezone)

	// Connect to MySQL server without database name to create database if needed
	mysqlConfigWithoutDB := &mysqlDriver.Config{
		User:      cfg.Username,
		Passwd:    cfg.Password,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		ParseTime: cfg.ParseTime,
		Loc:       loc,
		Params:    map[string]string{"charset": cfg.Charset},
	}

	GConfig := gormConfig
	if GConfig == nil {
		GConfig = &gorm.Config{}
	}
	GConfig.Logger = logger

	// First, connect without database to create it if needed
	tempDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlConfigWithoutDB.FormatDSN(),
		DefaultStringSize: 256,
	}), GConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL server: %w", err)
	}

	// Create database if not exists
	createSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET %s", cfg.Name, cfg.Charset)
	if err := tempDB.Exec(createSQL).Error; err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	// Close temporary connection
	tempSqlDB, err := tempDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get temp sql db: %w", err)
	}
	tempSqlDB.Close()

	// Now connect with database name in DSN
	mysqlConfigWithDB := &mysqlDriver.Config{
		User:      cfg.Username,
		Passwd:    cfg.Password,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DBName:    cfg.Name,
		ParseTime: cfg.ParseTime,
		Loc:       loc,
		Params:    map[string]string{"charset": cfg.Charset},
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               mysqlConfigWithDB.FormatDSN(),
		DefaultStringSize: 256,
	}), GConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql db: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &DB{db}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
