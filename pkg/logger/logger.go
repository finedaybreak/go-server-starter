package logger

import (
	"go-server-starter/internal/config"
	"go-server-starter/internal/enum"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(config *config.LoggerConfig, mode enum.ServerMode) *zap.Logger {
	configLevel := ParseLogLevel(config.Level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Info 级别的 Core（配置级别 ~ Error之前）
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= configLevel && lvl < zapcore.ErrorLevel
	})

	// Error 级别的 Core（Error、Panic、Fatal），也要遵循配置的最低级别
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && lvl >= configLevel
	})

	// 构建所有的 cores
	cores := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(NewLumberjackLoggerWriter("info", config)),
			infoLevel,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(NewLumberjackLoggerWriter("error", config)),
			errorLevel,
		),
	}

	// 如果需要控制台输出，添加控制台 core
	if config.ConsoleOutput {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(configLevel),
		))
	}

	// 创建多路输出的 core
	core := zapcore.NewTee(cores...)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	if mode == enum.ServerModeProd && config.Compress == false {
		logger.Warn("Log compression disabled in production mode. This may cause excessive disk usage.")
	}

	return logger
}

func NewLumberjackLoggerWriter(level string, config *config.LoggerConfig) *lumberjack.Logger {
	loggerWriter := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   filepath.Join(config.FileDir, level+".log"),
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	return loggerWriter
}

// ParseLogLevel 解析字符串格式的日志级别
func ParseLogLevel(level string) zapcore.Level {
	var l zapcore.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		// 如果解析失败，默认使用 Info 级别
		return zapcore.InfoLevel
	}
	return l
}
