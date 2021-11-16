package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMetrics(t *testing.T) {
	rf := Metrics("dummy_service")
	wrappedHandler := rf(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	}))
	req := &http.Request{
		Body: ioutil.NopCloser(new(bytes.Buffer)),
		URL: &url.URL{
			Path: "",
		},
	}
	w := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(w, req)
}

func TestMetricsExcludedPath(t *testing.T) {
	rf := Metrics("dummy_service2")
	wrappedHandler := rf(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	}))
	req := &http.Request{
		Body: ioutil.NopCloser(new(bytes.Buffer)),
		URL: &url.URL{
			Path: "metrics",
		},
	}
	w := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(w, req)
}
