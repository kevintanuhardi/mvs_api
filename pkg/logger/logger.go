package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instance *zap.Logger
var once sync.Once

func getLogger() *zap.Logger {
	once.Do(func() {
		instance = newLogger()
	})
	return instance
}

func newLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	cfg.Encoding = "json"
	cfg.EncoderConfig = EncoderConfig
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.DisableCaller = true

	logger, err := cfg.Build(
		serviceContextOpts(),
	)
	if err != nil {
		zap.S().Errorf("failed to create new logger with error: %s", err)
		panic(err)
	}
	return logger
}

type Logger interface {
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	With(fields ...zapcore.Field) Logger
}
type logFunc func(msg string, fields ...zapcore.Field)
type logWithFunc func(fields ...zapcore.Field) *zap.Logger
type logger struct {
	info   logFunc
	errorF logFunc
	fatal  logFunc
	with   logWithFunc
}

func (l *logger) Info(msg string, fields ...zapcore.Field) {
	l.info(msg, fields...)
}
func (l *logger) Error(msg string, fields ...zapcore.Field) {
	l.errorF(msg)
}
func (l *logger) Fatal(msg string, fields ...zapcore.Field) {
	l.fatal(msg, fields...)
}
func (l *logger) With(fields ...zapcore.Field) Logger {
	originalLogger := l.with(fields...)
	return &logger{
		info:   originalLogger.Info,
		errorF: originalLogger.Error,
		fatal:  originalLogger.Fatal,
		with:   originalLogger.With,
	}
}
