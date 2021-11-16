package adapter

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

func TestJulienSchmidtUseJulienHTTPRouter(t *testing.T) {
	router := UseJulienHTTPRouter()
	require.NotNil(t, router)
}

func TestJulienSchmidtRegisterWithMiddleware(t *testing.T) {
	router := UseJulienHTTPRouter()
	router.AddMiddlewareWrapper(func(next http.Handler) http.Handler {
		return next
	}).Register(http.MethodGet, "/ping", PING)
	router.ServeHTTP(httptest.NewRecorder(), &http.Request{
		URL: &url.URL{},
	})
}
func TestJulienSchmidtTransformator(t *testing.T) {
	router := &JulienSchmidtHTTPRouter{
		router: httprouter.New(),
	}
	resp := router.transformator(PING)
	resp(httptest.NewRecorder(), &http.Request{}, nil)
}
