package public

import (
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/domain"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
)

type Public struct {
	service domain.DomainService
	prefix  string
}

func NewHandler(service domain.DomainService) *Public {
	return &Public{
		service: service,
		prefix:  "/api",
	}
}
func (p *Public) Register(rr router.Registrator) {
	r := router.New(p.prefix, rr)
	r.GET("/ping", p.PING)
	r.POST("/user", p.RegisterUser)
	r.POST("/company", p.RegisterCompany)
}

func (p *Public) PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
