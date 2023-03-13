package log

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logger *Logger
	level  logger.LogLevel
}

func (l *Logger) NewGormLogger() *GormLogger {
	gl := &GormLogger{
		logger: l,
		level:  1,
	}
	return gl
}

func (gl *GormLogger) LogMode(LogLevel logger.LogLevel) logger.Interface {
	newGormLogger := &GormLogger{
		logger: gl.logger,
		level:  LogLevel,
	}
	return newGormLogger
}

func (gl *GormLogger) Info(_ context.Context, str string, args ...any) {
	gl.logger.Info("Gorm", fmt.Sprintf(str, args...))
}

func (gl *GormLogger) Warn(_ context.Context, str string, args ...any) {
	gl.logger.Warn("Gorm", fmt.Sprintf(str, args...))
}

func (gl *GormLogger) Error(_ context.Context, str string, args ...any) {
	gl.logger.Error("Gorm", fmt.Sprintf(str, args...))
}

func (gl *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	gl.logger.Debug("Gorm", fmt.Sprintf("sql trace: begin %s, sql: `%s`, rowsAffected: %d, err: %s", begin, sql, rowsAffected, err))
}
