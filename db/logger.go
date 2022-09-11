package db

import (
	"context"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	*zap.Logger
	LogLevel         logger.LogLevel
	SlowThreshold    time.Duration
	SkipCallerLookup bool
}

func initGormLogger(zapLogger *zap.Logger) *GormLogger {
	return &GormLogger{
		Logger:           zapLogger,
		LogLevel:         logger.Info,
		SlowThreshold:    100 * time.Millisecond,
		SkipCallerLookup: false,
	}
}

func (l *GormLogger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		default:
			return l.Logger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.Logger
}

// Info log Info for Gorm
func (l *GormLogger) Info(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	l.logger().Sugar().Debugf(str, args...)
}

// Warn log Warn for Gorm
func (l *GormLogger) Warn(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}

// Error log Error for Gorm
func (l *GormLogger) Error(_ context.Context, str string, args ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	l.logger().Sugar().Errorf(str, args...)
}
