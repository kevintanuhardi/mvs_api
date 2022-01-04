package public

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kevintanuhardi/mvs_api/domain/user/dto"
	"github.com/kevintanuhardi/mvs_api/pkg/response"
)


func (p *Public) RegisterUser(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	v := validator.New()
	var request dto.RegisterUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	err = v.Struct(request)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return response.NewJSONResponse().SetError(e)
		}
	}

	_, err = p.service.User.UserRegister(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

func (p *Public) Login(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	v := validator.New()
	var request dto.LoginRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	err = v.Struct(request)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return response.NewJSONResponse().SetError(e)
		}
	}

	_, err = p.service.User.Login(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}
