package webservice

import (
	"errors"
	"net/http"

	"github.com/kevintanuhardi/mvs_api/pkg/router"
)

type WebHandler struct {
	router             router.Registrator
	registrators       []WebRegistrator
	port               string
	httpListenAndServe httpListenAndServe
}

func NewWebService(port string, routerRegistrator router.Registrator, registrators ...WebRegistrator) WebService {
	return &WebHandler{
		router:             routerRegistrator,
		port:               port,
		registrators:       registrators,
		httpListenAndServe: http.ListenAndServe,
	}
}

func (w *WebHandler) Run() error {
	if w.port == "" {
		return errors.New("empty port defined")
	}
	for _, registrator := range w.registrators {
		registrator.Register(w.router)
	}
	return w.httpListenAndServe(w.port, w.router)
}

type httpListenAndServe func(addr string, handler http.Handler) error
