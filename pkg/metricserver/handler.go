package metricserver

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/kevintanuhardi/mvs_api/pkg/response"
	"github.com/kevintanuhardi/mvs_api/pkg/router"
)

type Metric struct {
	prefix string
}

func NewHandler() *Metric {
	return &Metric{
		prefix: "",
	}
}
func (p *Metric) Register(rr router.Registrator) {
	r := router.New(p.prefix, rr)
	r.GET("/metrics", p.HandlerMetrics)
}
func (p *Metric) HandlerMetrics(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	promhttp.Handler().ServeHTTP(w, r)
	return nil
}
