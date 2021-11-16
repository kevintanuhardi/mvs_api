package logger

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func For(ctx context.Context) Logger {
	loggerInstance := getLogger()
	if span := opentracing.SpanFromContext(ctx); span != nil {
		logger := &spanLogger{
			info:   loggerInstance.Info,
			errorF: loggerInstance.Error,
			fatal:  loggerInstance.Fatal,
			warn:   loggerInstance.Warn,
			with:   loggerInstance.With,
			span:   span,
		}
		if jaegerCtx, ok := span.Context().(jaeger.SpanContext); ok {
			logger.spanFields = []zapcore.Field{
				zap.String("trace_id", jaegerCtx.TraceID().String()),
				zap.String("span_id", jaegerCtx.SpanID().String()),
			}
		}
		return logger
	}
	return &logger{
		info:   loggerInstance.Info,
		errorF: loggerInstance.Error,
		fatal:  loggerInstance.Fatal,
		with:   loggerInstance.With,
	}
}
