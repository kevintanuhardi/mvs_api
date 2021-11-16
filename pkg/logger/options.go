package logger

import (
	"os"

	"go.uber.org/zap"
)

func serviceContextOpts() zap.Option {
	serviceName := os.Getenv("APP_NAME")
	serviceVersion := os.Getenv("APP_VERSION")
	return zap.Fields(
		LogServiceContext(&ServiceContext{
			Service: serviceName,
			Version: serviceVersion,
		}),
	)
}
