package public

import (
	"net/http"

	"github.com/kevintanuhardi/mvs_api/domain"
	"github.com/kevintanuhardi/mvs_api/pkg/response"
	"github.com/kevintanuhardi/mvs_api/pkg/router"
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
	r.GET("/healthz", p.PING)
	// USER routes
	r.POST("/user", p.RegisterUser)
	r.POST("/user/login", p.Login)
	// INVOICE routes
	r.POST("/invoice/excel", p.Login)
}

func (p *Public) PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
