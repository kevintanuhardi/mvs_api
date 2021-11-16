package render

import (
	"net/http/httptest"
	"testing"
)

func TestRender(t *testing.T) {
	w := httptest.NewRecorder()
	Response(w, 200, nil, nil, nil)
}
