package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

// logger
var (
	logger *zap.Logger
)

// provides logger
func init() {
	var err error
	if os.Getenv("USENSE_ENVIRONMENT") == "DEVELOPMENT" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
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
