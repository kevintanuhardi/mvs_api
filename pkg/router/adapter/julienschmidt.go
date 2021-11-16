package adapter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/middleware"
	"gitlab.warungpintar.co/sales-platform/brook/pkg/router"
)

type JulienSchmidtHTTPRouter struct {
	router  *httprouter.Router
	wrapper []middleware.Wrapper
}

func UseJulienHTTPRouter() router.Registrator {
	return &JulienSchmidtHTTPRouter{
		router: httprouter.New(),
	}
}

func (h *JulienSchmidtHTTPRouter) AddMiddlewareWrapper(wrapper ...middleware.Wrapper) router.Registrator {
	h.wrapper = append(h.wrapper, wrapper...)
	return h
}
func (h *JulienSchmidtHTTPRouter) Register(method, path string, handler router.HandlerFunc) {
	h.router.Handle(method, path, h.transformator(handler))
}

func (h *JulienSchmidtHTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *JulienSchmidtHTTPRouter) transformator(handler router.HandlerFunc) httprouter.Handle {
	wrappedHandler := handleWithResponse(handler)
	for _, v := range h.wrapper {
		wrappedHandler = v(wrappedHandler)
	}
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		wrappedHandler.ServeHTTP(rw, r)
	}
}
