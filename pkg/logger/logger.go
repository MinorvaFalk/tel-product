package logger

import (
	"os"
	"tel/product/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogger() {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleErrors := zapcore.Lock(os.Stderr)
	consoleDebugging := zapcore.Lock(os.Stdout)

	cores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewConsoleEncoder(
			zap.NewProductionEncoderConfig()),
			consoleErrors,
			highPriority,
		),
	}

	if config.ReadConfig().Env == "development" {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			consoleDebugging,
			lowPriority,
		))
	}

	core := zapcore.NewTee(cores...)

	log = zap.New(core)
}

func Zap() *zap.Logger {
	if log == nil {
		InitLogger()
	}

	return log
}
