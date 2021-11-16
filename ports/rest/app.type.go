package rest

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"gitlab.warungpintar.co/sales-platform/brook/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain/usecase"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/metricserver"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/webservice"
)

type Config struct {
	Cfg                    *config.Config
	GormStarter            adapter.GormStarter
	NewWebService          newWebServiceFunc
	GenerateRouter         generateRouter
	GenerateWebRegistrator generateWebRegistrator
	InitTracer             InitTracer
	RunMetricServer        metricserver.RunMetricServerFunc
}
type newWebServiceFunc func(port string, routerRegistrator router.Registrator, registrators ...webservice.WebRegistrator) webservice.WebService
type generateRouter func(tracer opentracing.Tracer) router.Registrator
type generateWebRegistrator func(service usecase.ServiceManager) []webservice.WebRegistrator
type InitTracer func(serviceName string) (opentracing.Tracer, io.Closer, error)
