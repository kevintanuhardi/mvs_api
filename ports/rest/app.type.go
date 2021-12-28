package rest

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/kevintanuhardi/mvs_api/adapter"
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain"
	"github.com/kevintanuhardi/mvs_api/pkg/metricserver"
	"github.com/kevintanuhardi/mvs_api/pkg/router"
	"github.com/kevintanuhardi/mvs_api/pkg/webservice"
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
type generateWebRegistrator func(service domain.DomainService) []webservice.WebRegistrator
type InitTracer func(serviceName string) (opentracing.Tracer, io.Closer, error)
