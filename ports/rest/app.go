package rest

import (
	"fmt"
	"log"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/opentracing/opentracing-go"
	"gitlab.warungpintar.co/sales-platform/brook/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/config"
	"gitlab.warungpintar.co/sales-platform/brook/domain"
	companyDomain "gitlab.warungpintar.co/sales-platform/brook/domain/company"
	companyMysql "gitlab.warungpintar.co/sales-platform/brook/domain/company/repository/mysql"
	otpDomain "gitlab.warungpintar.co/sales-platform/brook/domain/otp"
	otpMysql "gitlab.warungpintar.co/sales-platform/brook/domain/otp/repository/mysql"
	userDomain "gitlab.warungpintar.co/sales-platform/brook/domain/user"
	userMysql "gitlab.warungpintar.co/sales-platform/brook/domain/user/repository/mysql"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/metricserver"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/middleware"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
	routeradapter "gitlab.warungpintar.co/sales-platform/brook/pkg/router/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/tracing"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/webservice"
	"gitlab.warungpintar.co/sales-platform/brook/ports/rest/public"
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
			userMysql.NewRepository(db)),
		companyDomain.NewCompany(config.Config{},
			companyMysql.NewRepository(db)),
		otpDomain.NewOtp(config.Config{},
			userMysql.NewRepository(db),
			otpMysql.NewRepository(db)),
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
