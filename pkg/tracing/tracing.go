package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/kevintanuhardi/mvs_api/pkg/logger"
	"go.uber.org/zap"
)

const (
	samplerType  = "const"
	samplerParam = 1
)

func Init(serviceName string) (opentracing.Tracer, io.Closer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		logger.Bg().Error("cannot parse Jaeger env vars", zap.Error(err))
		return nil, nil, fmt.Errorf("jaeger init error: %v", err)
	}

	cfg.ServiceName = serviceName
	cfg.Sampler.Type = samplerType
	cfg.Sampler.Param = samplerParam
	cfg.Reporter = &config.ReporterConfig{
		LogSpans: true,
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		logger.Bg().Error("cannot initialize Jaeger Tracer", zap.Error(err))
		return tracer, closer, fmt.Errorf("jaeger init error: %v", err)
	}

	return tracer, closer, nil
}
