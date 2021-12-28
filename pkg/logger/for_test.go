package logger_test

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
	"github.com/kevintanuhardi/mvs_api/pkg/logger"
	"github.com/kevintanuhardi/mvs_api/pkg/tracing"
	"go.uber.org/zap"
)

func ExampleFor() {
	tracer, closer, _ := tracing.Init("hello")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	ctx := context.Background()
	span := tracer.StartSpan("say-hello")
	ctx = opentracing.ContextWithSpan(ctx, span)

	logger.For(ctx).Info("Info", zap.String("hello", "world"))
	logger.For(ctx).Error("Error", zap.String("hello", "world"))
}
func TestFor(t *testing.T) {
	tracer, closer, err := tracing.Init("hello")
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	require.NoError(t, err)
	ctx := context.Background()
	span := tracer.StartSpan("say-hello")
	ctx = opentracing.ContextWithSpan(ctx, span)

	logger.For(ctx).Info("Info", zap.String("hello", "world"))
	logger.For(ctx).Error("Error", zap.String("hello", "world"))
}
func TestWithoutSpan(t *testing.T) {
	logger.For(context.Background()).Info("Info", zap.String("hello", "world"))
	logger.For(context.Background()).Error("Error", zap.String("hello", "world"))
}
