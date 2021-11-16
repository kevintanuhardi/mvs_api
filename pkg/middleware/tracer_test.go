package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/opentracing/opentracing-go"
)

func TestTrace(t *testing.T) {
	tracer := opentracing.GlobalTracer()
	rf := Trace(tracer, TraceConfig{
		SkipURLPath: []string{
			"/v1/healthz",
			"/metrics",
		},
	})
	wrappedHandler := rf(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(501)
	}))
	req := &http.Request{
		Body: ioutil.NopCloser(new(bytes.Buffer)),
		URL: &url.URL{
			Path: "api",
		},
	}
	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)
}

func TestTraceSkipPath(t *testing.T) {
	tracer := opentracing.GlobalTracer()
	rf := Trace(tracer, TraceConfig{
		SkipURLPath: []string{
			"/v1/healthz",
			"/metrics",
		},
	})
	wrappedHandler := rf(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	}))
	req := &http.Request{
		Body: ioutil.NopCloser(new(bytes.Buffer)),
		URL: &url.URL{
			Path: "/metrics",
		},
	}
	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)
}
