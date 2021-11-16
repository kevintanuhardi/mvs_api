package logger

import (
	"testing"
)

func TestWith(t *testing.T) {
	originalLogger := getLogger()
	module := &logger{
		info:   originalLogger.Info,
		errorF: originalLogger.Info,
		fatal:  originalLogger.Info,
		with:   originalLogger.With,
	}
	module.With(LogServiceContext(&ServiceContext{
		Service: "Test",
		Version: "1.0",
	})).Info("Test")
}
func TestFatal(t *testing.T) {
	originalLogger := getLogger()
	module := &logger{
		info:   originalLogger.Info,
		errorF: originalLogger.Info,
		fatal:  originalLogger.Info,
		with:   originalLogger.With,
	}
	module.Fatal("Test")
}
