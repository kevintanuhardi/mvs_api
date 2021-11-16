# Basic Tracing Middleware

## Usage as middleware

### Golang Chi
```go

    import "gitlab.warungpintar.co/sales-platform/brook/pkg/middleware"

    func main() {
        tracer, closer := tracing.Init("myapp")
	    opentracing.SetGlobalTracer(tracer)
	    defer closer.Close()

        r := chi.NewRouter()
        r.Use(middleware.Trace(
            tracer, 
            middleware.TraceConfig{
		        SkipURLPath: []string{
			        "/metrics", // skip tracer for /metrics
		        },
	        },
        )
        r.Get("/hello", s.hello)
        r.Post("/foo", s.foo)
        _ = http.ListenAndServe(":8080", r)
    } 
```