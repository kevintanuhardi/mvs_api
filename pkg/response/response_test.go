package response

import (
	"net/http/httptest"
	"testing"
)

func TestWriteResponseBasic(t *testing.T) {
	basicResponse := &BasicResponse{
		Body:        make([]byte, 0),
		StatusCode:  200,
		ContentType: "text",
	}
	basicResponse.WriteResponse(httptest.NewRecorder())
}
func TestWriteResponseBasic_Error(t *testing.T) {
	basicResponse := &BasicResponse{
		Body:        nil,
		StatusCode:  200,
		ContentType: "text",
	}
	w := httptest.NewRecorder()
	basicResponse.WriteResponse(w)
}
