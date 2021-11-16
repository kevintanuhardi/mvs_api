package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logKeyServiceContext = "serviceContext"
)

type ServiceContext struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

func (s *ServiceContext) Clone() *ServiceContext {
	return &ServiceContext{
		Service: s.Service,
		Version: s.Version,
	}
}

func (s *ServiceContext) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddString("service", s.Service)
	e.AddString("version", s.Version)
	return nil
}

// LogServiceContext add service name and version
func LogServiceContext(ctx *ServiceContext) zapcore.Field {
	return zap.Object(logKeyServiceContext, ctx)
}
