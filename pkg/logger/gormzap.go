package logger

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	glogger "gorm.io/gorm/logger"
)

type gormLogger struct {
	ZapLogger            *zap.Logger
	SlowThreshold        time.Duration
	LogLevel             glogger.LogLevel
	IgnoreRecordNotFound bool

	msg string
}

func NewGormLogger() gormLogger {
	if log == nil {
		InitLogger()
	}

	return gormLogger{
		ZapLogger:            log,
		LogLevel:             glogger.Info,
		SlowThreshold:        500 * time.Millisecond,
		IgnoreRecordNotFound: true,
		msg:                  "gorm",
	}
}

func (l gormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	return gormLogger{
		ZapLogger:            l.ZapLogger,
		LogLevel:             l.LogLevel,
		SlowThreshold:        l.SlowThreshold,
		IgnoreRecordNotFound: l.IgnoreRecordNotFound,
	}
}

func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < glogger.Info {
		return
	}

	l.ZapLogger.Sugar().Infof(msg, data...)
}

func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < glogger.Warn {
		return
	}

	l.ZapLogger.Sugar().Warnf(msg, data...)
}

func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	// if l.LogLevel < glogger.Error {
	// 	return
	// }

	// l.ZapLogger.Sugar().Errorf(msg, data...)
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= glogger.Silent {
		return
	}

	elapsed := time.Since(begin)

	switch {
	// case err != nil && l.LogLevel >= glogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFound):
	// 	sql, rows := fc()
	// 	if rows == -1 {
	// 		l.ZapLogger.Error(l.msg,
	// 			zap.Error(err),
	// 			zap.Float64("time_elapsed", float64(elapsed.Nanoseconds())/1e6),
	// 			zap.String("sql", sql),
	// 		)
	// 	} else {
	// 		l.ZapLogger.Error(l.msg,
	// 			zap.Error(err),
	// 			zap.Float64("time_elapsed", float64(elapsed.Nanoseconds())/1e6),
	// 			zap.Int64("rows", rows),
	// 			zap.String("sql", sql),
	// 		)
	// 	}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= glogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.ZapLogger.Warn(l.msg,
				zap.String("warning", slowLog),
				zap.Float64("time_elapsed", float64(elapsed.Nanoseconds())/1e6),
				zap.String("sql", sql),
			)
		} else {
			l.ZapLogger.Warn(l.msg,
				zap.String("warning", slowLog),
				zap.Float64("time_elapsed", float64(elapsed.Nanoseconds())/1e6),
				zap.String("sql", sql),
				zap.Int64("rows", rows),
			)
		}
	case l.LogLevel == glogger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.ZapLogger.Info(l.msg,
				zap.Float64("time_elapsed", float64(elapsed.Nanoseconds())/1e6),
				zap.String("sql", sql),
			)
		} else {
			l.ZapLogger.Info(l.msg,
				zap.Float64("time_elapsed", float64(elapsed.Nanoseconds())/1e6),
				zap.String("sql", sql),
				zap.Int64("rows", rows),
			)
		}
	}

}
