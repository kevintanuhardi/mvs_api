package response

import (
	"errors"
	"math"
	"net/http/httptest"
	"testing"
)

func TestJSONResponseWriteResponse(t *testing.T) {
	resp := &JSONResponse{
		BasicResponse: BasicResponse{
			Body:        make([]byte, 0),
			StatusCode:  200,
			ContentType: "text",
		},
	}
	resp.WriteResponse(httptest.NewRecorder())
}

func TestJSONResponseWriteResponseWithError(t *testing.T) {
	resp := &JSONResponse{
		BasicResponse: BasicResponse{
			Body:        make([]byte, 0),
			StatusCode:  200,
			ContentType: "text",
		},
		JSONBody: JSONBody{
			Error: NewError("Something Bad Happen", 300),
		},
	}
	resp.WriteResponse(httptest.NewRecorder())
}

func TestJSONResponseWriteResponseErrorMarshall(t *testing.T) {
	resp := &JSONResponse{
		BasicResponse: BasicResponse{
			Body:        make([]byte, 0),
			StatusCode:  200,
			ContentType: "text",
		},
		JSONBody: JSONBody{
			Data: math.Inf(1),
		},
	}
	resp.WriteResponse(httptest.NewRecorder())
}

func TestJSONResponseSetKnownError(t *testing.T) {
	resp := NewJSONResponse()
	resp.SetError(ErrorBadRequest)
}

func TestJSONResponseSetUnKnownError(t *testing.T) {
	resp := NewJSONResponse()
	resp.SetError(errors.New("Something Bad Happen"))
}

func TestJSONResponseSetMessage(t *testing.T) {
	resp := NewJSONResponse()
	resp.SetMessage("message")
}

func TestJSONResponseSetData(t *testing.T) {
	resp := NewJSONResponse()
	resp.SetData([]int{1, 23, 4, 5, 67, 8910})
}
