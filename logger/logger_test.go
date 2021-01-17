package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	Info("Hello", zap.Bool("working", true))
}

func TestWarn(t *testing.T) {
	Warn("Warning", zap.Bool("working", true))
}

func TestDebug(t *testing.T) {
	Debug("Debug", zap.Bool("working", true))
}

func TestError(t *testing.T) {
	Error("Error", zap.Bool("working", true))
}
