package response

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	BasicResponse
	JSONBody JSONBody
	Error    error
}
type JSONBody struct {
	*Error
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{
		BasicResponse: BasicResponse{
			ContentType: JSONContentType,
			StatusCode:  http.StatusOK,
		},
	}
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.JSONBody.Data = data
	return r
}

func (r *JSONResponse) SetMessage(message string) *JSONResponse {
	r.JSONBody.Message = message
	return r
}

func (r *JSONResponse) SetError(err error) *JSONResponse {
	if respErr, ok := err.(*Error); ok {
		r.JSONBody.Error = respErr
	} else {
		// when unspecified error is provided it will categorize the response as internal server error
		r.JSONBody.Error = NewError(err.Error(), http.StatusInternalServerError)
	}
	return r
}

func (r *JSONResponse) WriteResponse(w http.ResponseWriter) {
	b, err := json.Marshal(r.JSONBody)
	if err != nil {
		JSONBody := JSONBody{
			Error: NewError(err.Error(), http.StatusInternalServerError),
		}
		b, _ = json.Marshal(JSONBody)
	}
	r.Body = b
	if r.JSONBody.Error != nil {
		r.StatusCode = r.JSONBody.ErrorCode
	}
	r.BasicResponse.WriteResponse(w)
}
