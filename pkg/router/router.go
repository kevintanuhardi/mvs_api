package router

import (
	"net/http"

	"gitlab.warungpintar.co/sales-platform/brook/pkg/middleware"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/response"
)

type HTTPRouter interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	OPTIONS(path string, handler HandlerFunc)
}
type Registrator interface {
	AddMiddlewareWrapper(wrapper ...middleware.Wrapper) Registrator
	Register(method, path string, handler HandlerFunc)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
type httpRouter struct {
	router Registrator
	prefix string
}

func New(prefix string, router Registrator) HTTPRouter {
	return &httpRouter{
		router: router,
		prefix: prefix,
	}
}

func (h *httpRouter) GET(path string, handler HandlerFunc) {
	h.register(http.MethodGet, path, handler)
}

func (h *httpRouter) POST(path string, handler HandlerFunc) {
	h.register(http.MethodPost, path, handler)
}

func (h *httpRouter) PUT(path string, handler HandlerFunc) {
	h.register(http.MethodPut, path, handler)
}

func (h *httpRouter) OPTIONS(path string, handler HandlerFunc) {
	h.register(http.MethodOptions, path, handler)
}

func (h *httpRouter) register(method, path string, handler HandlerFunc) {
	fullpath := h.prefix + path
	h.router.Register(method, fullpath, handler)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) response.HTTPResponse
