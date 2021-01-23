package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// logger
var (
	logger *zap.Logger
)

// provides logger
func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Zap logger not initialized. %v", err)
	}
}

// Info wrapper for zap.Info
func Info(m string, fields ...zapcore.Field) {
	logger.Info(m, fields...)
}

// Warn wrapper for zap.Warn
func Warn(m string, fields ...zapcore.Field) {
	logger.Warn(m, fields...)
}

// Debug wrapper for zap.Debug
func Debug(m string, fields ...zapcore.Field) {
	logger.Debug(m, fields...)
}

// Error wrapper for zap.Error
func Error(m string, fields ...zapcore.Field) {
	logger.Error(m, fields...)
}
