package public

import (
	"encoding/json"
	"net/http"

	"github.com/kevintanuhardi/mvs_api/domain/user/entity"
	"github.com/kevintanuhardi/mvs_api/pkg/response"
)


func (p *Public) RegisterUser(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	var request entity.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	_, err = p.service.User.UserRegister(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

// func (p *Public) ActivateUser(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
// 	var request dto.UserActivateRequest

// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&request)
// 	if err != nil {
// 		return response.NewJSONResponse().SetError(err)
// 	}

// 	_, _, err = p.service.User.UserActivation(r.Context(), &request)
// 	if err != nil {
// 		return response.NewJSONResponse().SetError(err)
// 	}
// 	return response.NewJSONResponse().SetMessage("Success")
// }
