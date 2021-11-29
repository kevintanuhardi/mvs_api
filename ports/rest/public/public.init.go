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
	r.GET("/healthz", p.PING)
	r.POST("/user", p.RegisterUser)
	r.PUT("/user/activate", p.ActivateUser)
	r.POST("/company", p.RegisterCompany)
	r.POST("/otp/send", p.OTP)
	r.POST("/otp/verify", p.VerifyOTP)
	r.POST("/otp/login", p.AuthLogin)
}

func (p *Public) PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
