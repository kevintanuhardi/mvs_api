package public

import (
	"encoding/json"
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/domain/otp/dto"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)

func (p *Public) OTP(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	decoder := json.NewDecoder(r.Body)
	var request dto.SendOTPRequest
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	otp, err := p.service.Otp.SendOTP(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(otp)
}

func (p *Public) VerifyOTP(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	decoder := json.NewDecoder(r.Body)
	var request dto.VerifyOTPRequest
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	otp, err := p.service.Otp.Verify(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(otp)
}

func (p *Public) AuthLogin(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	decoder := json.NewDecoder(r.Body)
	var request dto.LoginRequest
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	otp, err := p.service.Otp.Login(r.Context(), &request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(otp)
}
