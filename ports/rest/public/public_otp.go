package public

import (
	"encoding/json"
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/domain/usecase"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)

func (p *Public) OTP(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	decoder := json.NewDecoder(r.Body)
	var request usecase.SendOTPRequest
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	otp, err := p.service.SendOTP(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(otp)
}

func (p *Public) VerifyOTP(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	decoder := json.NewDecoder(r.Body)
	var request usecase.VerifyOTPRequest
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	otp, err := p.service.Verify(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(otp)
}

func (p *Public) AuthLogin(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	decoder := json.NewDecoder(r.Body)
	var request usecase.LoginRequest
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	otp, err := p.service.Login(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(otp)
}
