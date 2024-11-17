package middleware

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"student-service/internal/config"
)

func NewLogger(cfg *config.Config) *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(getLogLevel(cfg.Logging.Level))

	logger, _ := config.Build()
	return logger
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
