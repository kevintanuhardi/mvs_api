package metricserver

import (
	"fmt"
	"log"

	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router/adapter"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/webservice"
)

type RunMetricServerFunc func(cfg *Config)

func RunMetricServer(cfg *Config) {
	server := cfg.NewWebService(
		fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		cfg.Router,
		NewHandler(),
	)
	log.Printf("Serving metric server on 0.0.0.0:%d\n", cfg.Port)
	err := server.Run()
	if err != nil {
		log.Println("Failed To run Metric Server because of", err.Error())
	}
}

func GetDefaultConfig(port int) *Config {
	return &Config{
		NewWebService: webservice.NewWebService,
		Port:          port,
		Router:        adapter.UseChiRouter(),
	}
}

type Config struct {
	Port          int
	NewWebService newWebServiceFunc
	Router        router.Registrator
}
type newWebServiceFunc func(port string, routerRegistrator router.Registrator, registrators ...webservice.WebRegistrator) webservice.WebService
