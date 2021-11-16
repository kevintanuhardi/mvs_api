package intools

import (
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)

func (p *Intools) GetOrder(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	orders, err := p.service.OrderList(r.Context())
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(orders)
}
