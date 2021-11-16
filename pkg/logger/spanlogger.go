package logger

import (
	"time"

	"github.com/opentracing/opentracing-go"
	tag "github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap/zapcore"
)

type spanLogger struct {
	info       logFunc
	errorF     logFunc
	fatal      logFunc
	warn       logFunc
	with       logWithFunc
	span       opentracing.Span
	spanFields []zapcore.Field
}

func (sl *spanLogger) Info(msg string, fields ...zapcore.Field) {
	sl.logToSpan("info", msg, fields...)
	sl.info(msg, append(sl.spanFields, fields...)...)
}

func (sl *spanLogger) Error(msg string, fields ...zapcore.Field) {
	sl.logToSpan("error", msg, fields...)
	tag.Error.Set(sl.span, true)
	sl.errorF(msg, append(sl.spanFields, fields...)...)
}

func (sl *spanLogger) Warn(msg string, fields ...zapcore.Field) {
	sl.warn(msg, append(sl.spanFields, fields...)...)
}

func (sl *spanLogger) Fatal(msg string, fields ...zapcore.Field) {
	sl.logToSpan("fatal", msg, fields...)
	tag.Error.Set(sl.span, true)
	sl.fatal(msg, append(sl.spanFields, fields...)...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (sl *spanLogger) With(fields ...zapcore.Field) Logger {
	original := sl.with(fields...)
	return &spanLogger{
		info:       original.Info,
		errorF:     original.Error,
		fatal:      original.Fatal,
		warn:       original.Warn,
		with:       original.With,
		span:       sl.span,
		spanFields: sl.spanFields,
	}
}

func (sl *spanLogger) logToSpan(level, msg string, fields ...zapcore.Field) {
	initSize := 2 // add event and level
	fa := fieldAdapter(make([]log.Field, 0, initSize+len(fields)))
	fa = append(fa, log.String("event", msg), log.String("level", level))
	for _, field := range fields {
		field.AddTo(&fa)
	}
	sl.span.LogFields(fa...)
}

type fieldAdapter []log.Field

func (fa *fieldAdapter) AddBool(key string, value bool) {
	*fa = append(*fa, log.Bool(key, value))
}

func (fa *fieldAdapter) AddFloat64(key string, value float64) {
	*fa = append(*fa, log.Float64(key, value))
}

func (fa *fieldAdapter) AddFloat32(key string, value float32) {
	*fa = append(*fa, log.Float64(key, float64(value)))
}

func (fa *fieldAdapter) AddInt(key string, value int) {
	*fa = append(*fa, log.Int(key, value))
}

func (fa *fieldAdapter) AddInt64(key string, value int64) {
	*fa = append(*fa, log.Int64(key, value))
}

func (fa *fieldAdapter) AddInt32(key string, value int32) {
	*fa = append(*fa, log.Int64(key, int64(value)))
}

func (fa *fieldAdapter) AddInt16(key string, value int16) {
	*fa = append(*fa, log.Int64(key, int64(value)))
}

func (fa *fieldAdapter) AddInt8(key string, value int8) {
	*fa = append(*fa, log.Int64(key, int64(value)))
}

func (fa *fieldAdapter) AddUint(key string, value uint) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *fieldAdapter) AddUint64(key string, value uint64) {
	*fa = append(*fa, log.Uint64(key, value))
}

func (fa *fieldAdapter) AddUint32(key string, value uint32) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *fieldAdapter) AddUint16(key string, value uint16) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *fieldAdapter) AddUint8(key string, value uint8) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

// AddUintptr empty because not necessary for now
func (fa *fieldAdapter) AddUintptr(key string, value uintptr)                        {}
func (fa *fieldAdapter) AddArray(key string, marshaler zapcore.ArrayMarshaler) error { return nil }

// AddComplex128 empty because not necessary for now
func (fa *fieldAdapter) AddComplex128(key string, value complex128) {}

// AddComplex64 empty because not necessary for now
func (fa *fieldAdapter) AddComplex64(key string, value complex64)                  {}
func (fa *fieldAdapter) AddObject(key string, value zapcore.ObjectMarshaler) error { return nil }
func (fa *fieldAdapter) AddReflected(key string, value interface{}) error          { return nil }

// OpenNamespace empty because not necessary for now
func (fa *fieldAdapter) OpenNamespace(key string) {}

func (fa *fieldAdapter) AddDuration(key string, value time.Duration) {
	*fa = append(*fa, log.String(key, value.String()))
}

func (fa *fieldAdapter) AddTime(key string, value time.Time) {
	*fa = append(*fa, log.String(key, value.String()))
}

func (fa *fieldAdapter) AddBinary(key string, value []byte) {
	*fa = append(*fa, log.Object(key, value))
}

func (fa *fieldAdapter) AddByteString(key string, value []byte) {
	*fa = append(*fa, log.Object(key, value))
}

func (fa *fieldAdapter) AddString(key, value string) {
	if key != "" && value != "" {
		*fa = append(*fa, log.String(key, value))
	}
}
