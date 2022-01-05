package public

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevintanuhardi/mvs_api/domain/invoice/dto"
	"github.com/kevintanuhardi/mvs_api/pkg/response"
)

func (p *Public) CreateInvoice(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	// v := validator.New()
	var request *dto.CreateInvoiceRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	// err = v.Struct(request)

	// if err != nil {
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		return response.NewJSONResponse().SetError(e)
	// 	}
	// }

	invoiceResponse, err := p.service.Invoice.CreateInvoice(r.Context(), request)
	if err != nil {
		fmt.Println(err);
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(invoiceResponse)
}

func (p *Public) GetInvoiceXls(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	// v := validator.New()
	var request *dto.GetInvoiceXlsRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}

	// err = v.Struct(request)

	// if err != nil {
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		return response.NewJSONResponse().SetError(e)
	// 	}
	// }

	xls, filename, err := p.service.Invoice.GetInvoiceXls(r.Context(), request)
	if err != nil {
		fmt.Println(err);
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewFileResponse().SetFileWriter(xls, filename)
}