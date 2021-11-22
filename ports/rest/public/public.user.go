package public

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/domain/user/entity"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)


func (p *Public) RegisterUser(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	var request entity.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	fmt.Printf("not a body: %+v", request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	err = p.service.User.UserRegister(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData("Success")
}
