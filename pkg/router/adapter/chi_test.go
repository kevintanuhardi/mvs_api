package adapter

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)

func TestChiUseChiRouter(t *testing.T) {
	router := UseChiRouter()
	require.NotNil(t, router.router)
}

func TestChiRegisterWithMiddleware(t *testing.T) {
	router := UseChiRouter()
	router.AddMiddlewareWrapper(func(next http.Handler) http.Handler {
		return next
	}).Register(http.MethodGet, "/ping", PING)
	router.ServeHTTP(httptest.NewRecorder(), &http.Request{
		URL: &url.URL{},
	})
}
func PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
func TestChiHandleWithResponse(t *testing.T) {
	resp := handleWithResponse(PING)
	resp.ServeHTTP(httptest.NewRecorder(), &http.Request{})
}
