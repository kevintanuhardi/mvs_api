
# Logger

Standarisasi logger untuk applikasi warung pintar. 
Logger dibagi menjadi dua `Bg` untuk background dan `For(ctx)` jika log ingin dimasukan kedalam jaeger.
Untuk menggunakan `For(ctx)` pastikan context sudah diisi trace id.

Logger menggunakan [zap](https://github.com/uber-go/zap) dan encodernya mengikuti standard dari stackdriver mengubah "level" menjadi "severity". 

```go
span, ctx := opentracing.StartSpanFromContext(ctx, "product_application_notifySyncDone")
defer span.Finish()
```

## USAGE

```go
import "gitlab.warungpintar.co/sales-platform/brook/pkg/logger"
```

- Environtment:
    = APP_NAME
    - APP_VERSION

### INFO

- Usage:
```go
    logger.Bg().Info("Info", zap.String("payload", "here"))
```

- Response

```json
{"severity":"INFO","eventTime":"2021-04-13T11:57:47.162+0700","message":"Info","serviceContext":{"service":"","version":""},"payload":"here"}
```

### INFO With Context

- Usage:
```go
    logger.For(ctx).Info("Info", zap.String("payload", "here"))
```

- Response
```json
{"severity":"INFO","eventTime":"2021-04-13T12:58:28.255+0700","message":"Info","serviceContext":{"service":"","version":""},"trace_id":"3ecec1bd1d8d97db","span_id":"3ecec1bd1d8d97db","payload":"here"}
```

### ERROR With Context

- Usage:
```go
    logger.For(ctx).Error("Info", zap.String("payload", "here"))
```

- Response
```json
{"severity":"ERROR","eventTime":"2021-04-13T12:58:28.255+0700","message":"Error","serviceContext":{"service":"","version":""},"trace_id":"3ecec1bd1d8d97db","span_id":"3ecec1bd1d8d97db","payload":"here","stacktrace":"gitlab.warungpintar.co/sales-platform/brook/pkg/logger.spanLogger.Error\n\t/Users/kit/Programming/projects/warungpintar/kit/logger/spanlogger.go:27\ngitlab.warungpintar.co/sales-platform/brook/pkg/logger_test.ExampleFor\n\t/Users/kit/Programming/projects/warungpintar/kit/logger/for_test.go:22\ntesting.runExample\n\t/usr/local/Cellar/go/1.15.7_1/libexec/src/testing/run_example.go:62\ntesting.runExamples\n\t/usr/local/Cellar/go/1.15.7_1/libexec/src/testing/example.go:44\ntesting.(*M).Run\n\t/usr/local/Cellar/go/1.15.7_1/libexec/src/testing/testing.go:1346\nmain.main\n\t_testmain.go:45\nruntime.main\n\t/usr/local/Cellar/go/1.15.7_1/libexec/src/runtime/proc.go:204"}
```