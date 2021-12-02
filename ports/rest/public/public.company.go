package public

import (
	"encoding/json"
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/domain/company/entity"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)


func (p *Public) RegisterCompany(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	var request entity.Company

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	err = p.service.Company.CompanyRegister(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}
