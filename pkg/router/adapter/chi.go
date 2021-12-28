package adapter

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kevintanuhardi/mvs_api/pkg/middleware"
	"github.com/kevintanuhardi/mvs_api/pkg/router"
)

type ChiRouter struct {
	router  chi.Router
	wrapper []middleware.Wrapper
}

func UseChiRouter() *ChiRouter {
	return &ChiRouter{
		router: chi.NewRouter(),
	}
}
func (h *ChiRouter) AddMiddlewareWrapper(wrapper ...middleware.Wrapper) router.Registrator {
	h.wrapper = append(h.wrapper, wrapper...)
	return h
}
func (h *ChiRouter) Register(method, path string, handler router.HandlerFunc) {
	h.router.Method(method, path, h.transformator(handler))
}
func (h *ChiRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
func (h *ChiRouter) transformator(handler router.HandlerFunc) http.Handler {
	wrappedHandler := handleWithResponse(handler)
	for _, v := range h.wrapper {
		wrappedHandler = v(wrappedHandler)
	}
	return wrappedHandler
}
func handleWithResponse(handler router.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := handler(rw, r)
		if resp != nil {
			resp.WriteResponse(rw)
		}
	})
}
