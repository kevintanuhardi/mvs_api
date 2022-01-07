package rest

import (
	"fmt"
	"log"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/kevintanuhardi/mvs_api/adapter"
	"github.com/kevintanuhardi/mvs_api/config"
	"github.com/kevintanuhardi/mvs_api/domain"
	invoiceDomain "github.com/kevintanuhardi/mvs_api/domain/invoice"
	invoiceMysql "github.com/kevintanuhardi/mvs_api/domain/invoice/repository/mysql"
	productDomain "github.com/kevintanuhardi/mvs_api/domain/product"
	productMysql "github.com/kevintanuhardi/mvs_api/domain/product/repository/mysql"
	userDomain "github.com/kevintanuhardi/mvs_api/domain/user"
	userMysql "github.com/kevintanuhardi/mvs_api/domain/user/repository/mysql"
	"github.com/kevintanuhardi/mvs_api/pkg/metricserver"
	"github.com/kevintanuhardi/mvs_api/pkg/middleware"
	"github.com/kevintanuhardi/mvs_api/pkg/router"
	routeradapter "github.com/kevintanuhardi/mvs_api/pkg/router/adapter"
	"github.com/kevintanuhardi/mvs_api/pkg/tracing"
	"github.com/kevintanuhardi/mvs_api/pkg/webservice"
	"github.com/kevintanuhardi/mvs_api/ports/rest/public"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// Application for rest service
func Application(cfg *Config) error {
	db, err := AppWithGorm(cfg)
	if err != nil {
		return err
	}
	service := initService(db)
	trace, _, err := cfg.InitTracer(cfg.Cfg.App.Name)
	if err != nil {
		return err
	}

	server := cfg.NewWebService(
		fmt.Sprintf("0.0.0.0:%d", cfg.Cfg.Port.HTTP),
		cfg.GenerateRouter(trace),
		cfg.GenerateWebRegistrator(service)...,
	)
	go cfg.RunMetricServer(metricserver.GetDefaultConfig(cfg.Cfg.Port.HTTPMetric))
	log.Println("REST Server running on 0.0.0.0:", cfg.Cfg.Port.HTTP)
	return server.Run()
}
func GetDefaultConfig(cfg *config.Config) *Config {
	return &Config{
		Cfg:                    cfg,
		GormStarter:            adapter.NewGormStarter(),
		NewWebService:          webservice.NewWebService,
		GenerateRouter:         getRouter,
		GenerateWebRegistrator: getWebRegistrator,
		InitTracer:             tracing.Init,
		RunMetricServer:        metricserver.RunMetricServer,
	}
}
func getRouter(tracer opentracing.Tracer) router.Registrator {
	module := routeradapter.UseChiRouter()
	module.AddMiddlewareWrapper(
		chimiddleware.Recoverer,
		middleware.Metrics("brook_api"),
		middleware.Trace(
			tracer,
			middleware.TraceConfig{
				SkipURLPath: []string{
					"/healthz",
					"/metrics",
				},
			}),
	)

	return module
}
func getWebRegistrator(service domain.DomainService) []webservice.WebRegistrator {
	resp := []webservice.WebRegistrator{
		public.NewHandler(service),
	}
	return resp
}
func initService(db *gorm.DB) domain.DomainService {
	return domain.NewDomain(
		userDomain.NewUser(config.Config{},
			userMysql.NewRepository(db),
		),
		invoiceDomain.NewInvoice(config.Config{},
			invoiceMysql.NewRepository(db),
		),
		productDomain.NewProduct(config.Config{},
			productMysql.NewRepository(db),
		),
	)
}
func AppWithGorm(cfg *Config) (*gorm.DB, error) {
	db, err := cfg.GormStarter.ConnectToGorm(adapter.Config{
		Master:           cfg.Cfg.DB.MasterDSN,
		Replicas:         []string{cfg.Cfg.DB.ReplicaDSN},
		MaxIdleConns:     cfg.Cfg.DB.MaxIdle,
		MaxOpenConns:     cfg.Cfg.DB.MaxOpen,
		MaxLifetimeConns: cfg.Cfg.DB.MaxLifeTime,
	}, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
