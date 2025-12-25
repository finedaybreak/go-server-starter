package logger

import (
	"context"
	"errors"
	"fmt"
	"go-server-starter/internal/config"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var gormSourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	gormSourceDir = sourceDir(file)
}

type GormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewGormLogger(logger *zap.Logger, config config.GormLoggerConfig) (*GormLogger, error) {
	LogLevel, err := ParseStringGormLogLevel(config.Level)
	if err != nil {
		return nil, err
	}
	return &GormLogger{
		ZapLogger:                 logger,
		LogLevel:                  LogLevel,
		SlowThreshold:             config.SlowThreshold,
		SkipCallerLookup:          config.SkipCallerLookup,
		IgnoreRecordNotFoundError: config.IgnoreRecordNotFoundError,
	}, nil
}

func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= gormlogger.Info {
		l.ZapLogger.Sugar().Infof(msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= gormlogger.Warn {
		l.ZapLogger.Sugar().Warnf(msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...any) {
	if l.LogLevel >= gormlogger.Error {
		l.ZapLogger.Sugar().Errorf(msg, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
	}

	if !l.SkipCallerLookup {
		fields = append(fields, zap.String("file", FileWithLineNum()))
	}

	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		l.ZapLogger.Error("SQL Error", append(fields, zap.Error(err))...)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormlogger.Warn:
		l.ZapLogger.Warn("Slow SQL", append(fields, zap.Duration("threshold", l.SlowThreshold))...)
	case l.LogLevel == gormlogger.Info:
		l.ZapLogger.Info("SQL", fields...)
	}
}

func sourceDir(file string) string {
	dir := filepath.Dir(file)
	dir = filepath.Dir(dir)

	s := filepath.Dir(dir)
	if filepath.Base(s) != "gorm.io" {
		s = dir
	}
	return filepath.ToSlash(s) + "/"
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	pcs := [13]uintptr{}
	// the third caller usually from gorm internal
	len := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:len])
	for range len {
		// second return value is "more", not "ok"
		frame, _ := frames.Next()
		if (!strings.HasPrefix(frame.File, gormSourceDir) ||
			strings.HasSuffix(frame.File, "_test.go")) && !strings.HasSuffix(frame.File, ".gen.go") {
			return string(strconv.AppendInt(append([]byte(frame.File), ':'), int64(frame.Line), 10))
		}
	}

	return ""
}

func ParseStringGormLogLevel(level string) (gormlogger.LogLevel, error) {
	switch level {
	case "silent":
		return gormlogger.Silent, nil
	case "info":
		return gormlogger.Info, nil
	case "warn":
		return gormlogger.Warn, nil
	case "error":
		return gormlogger.Error, nil
	default:
		return gormlogger.Info, fmt.Errorf("invalid gorm log level: %s", level)
	}
}
