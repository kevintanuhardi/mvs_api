package intools

import (
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/domain/usecase"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
)

type Intools struct {
	service usecase.ServiceManager
	prefix  string
}

func NewHandler(service usecase.ServiceManager) *Intools {
	return &Intools{
		service: service,
		prefix:  "/intools",
	}
}

func (p *Intools) Register(rr router.Registrator) {
	r := router.New(p.prefix, rr)
	r.GET("/ping", p.PING)
	r.GET("/order", p.GetOrder)
}

func (p *Intools) PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
