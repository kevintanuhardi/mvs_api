# Basic Metrics Middleware

## Usage as middleware

### Golang Chi
```go

    import "github.com/kevintanuhardi/mvs_api/pkg/middleware"

    func main() {
        r := chi.NewRouter()
        r.Use(middleware.Metrics("merchant_myapp", []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}...))
        r.Get("/hello", s.hello)
        r.Post("/foo", s.foo)
        _ = http.ListenAndServe(":8080", r)
    } 
```

## Details

- Metrics
```bash
    # HELP merchant_myapp_failed_sync_product_to_mbs The total number of product sync failed to mbs request
    # TYPE merchant_myapp_failed_sync_product_to_mbs counter
    merchant_myapp_failed_sync_product_to_mbs 1
    # HELP merchant_myapp_req_fe_grpc_total The total number of incoming transaction in garfield via grpc
    # TYPE merchant_myapp_req_fe_grpc_total counter
    merchant_myapp_req_fe_grpc_total 0
    # HELP merchant_myapp_req_fe_http_total The total number of incoming transaction in garfield via http
    # TYPE merchant_myapp_req_fe_http_total counter
    merchant_myapp_req_fe_http_total 2
    # HELP merchant_myapp_request_duration_milliseconds How long it took to process the request, partitioned by status code, method and HTTP path.
    # TYPE merchant_myapp_request_duration_milliseconds histogram
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="GET",path="/myapp/v1/hello",service="merchant_myapp",le="300"} 0
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="GET",path="/myapp/v1/hello",service="merchant_myapp",le="1200"} 1
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="GET",path="/myapp/v1/hello",service="merchant_myapp",le="5000"} 1
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="GET",path="/myapp/v1/hello",service="merchant_myapp",le="+Inf"} 1
    merchant_myapp_request_duration_milliseconds_sum{code="OK",method="GET",path="/myapp/v1/hello",service="merchant_myapp"} 1168.779527
    merchant_myapp_request_duration_milliseconds_count{code="OK",method="GET",path="/myapp/v1/hello",service="merchant_myapp"} 1
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp",le="300"} 0
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp",le="1200"} 1
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp",le="5000"} 1
    merchant_myapp_request_duration_milliseconds_bucket{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp",le="+Inf"} 1
    merchant_myapp_request_duration_milliseconds_sum{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp"} 907.541947
    merchant_myapp_request_duration_milliseconds_count{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp"} 1
    # HELP merchant_myapp_requests_total How many HTTP requests processed, partitioned by status code, method and HTTP path.
    # TYPE merchant_myapp_requests_total counter
    merchant_myapp_requests_total{code="OK",method="GET",path="/myapp/v1/myapp/hello",service="merchant_myapp"} 1
    merchant_myapp_requests_total{code="OK",method="POST",path="/myapp/v1/foo",service="merchant_myapp"} 1
```

- Query

```yml
# HTTP response codes
rate(merchant_myapp_requests_total[5m])

# HTTP requests total
rate(merchant_myapp_requests_total[5m])

# HTTP request latency count
rate(merchant_myapp_request_duration_milliseconds_count[5m])

# HTTP latency percentiles
P99: histogram_quantile(0.99, sum(rate(merchant_myapp_request_duration_milliseconds_bucket[5m])) by (le))
P95: histogram_quantile(0.95, sum(rate(merchant_myapp_request_duration_milliseconds_bucket[5m])) by (le))
P50: histogram_quantile(0.50, sum(rate(merchant_myapp_request_duration_milliseconds_bucket[5m])) by (le))

```