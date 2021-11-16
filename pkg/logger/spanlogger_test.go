package logger

import (
	"context"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestSpanLoggerWith(t *testing.T) {
	originalLogger := getLogger()
	span, _ := opentracing.StartSpanFromContext(context.Background(), "test")
	defer span.Finish()
	module := &spanLogger{
		info:   originalLogger.Info,
		errorF: originalLogger.Info,
		fatal:  originalLogger.Info,
		warn:   originalLogger.Info,
		with:   originalLogger.With,
		span:   span,
	}

	module.With(LogServiceContext(&ServiceContext{
		Service: "Test",
		Version: "1.0",
	})).Info("Test")
}

func TestSpanLoggerFatal(t *testing.T) {
	originalLogger := getLogger()
	span, _ := opentracing.StartSpanFromContext(context.Background(), "test")
	defer span.Finish()

	module := &spanLogger{
		info:   originalLogger.Info,
		errorF: originalLogger.Info,
		fatal:  originalLogger.Info,
		with:   originalLogger.With,
		span:   span,
	}
	module.Fatal("Test")
}

func TestSpanLoggerWarn(t *testing.T) {
	originalLogger := getLogger()
	span, _ := opentracing.StartSpanFromContext(context.Background(), "test")
	defer span.Finish()

	module := &spanLogger{
		info:   originalLogger.Info,
		errorF: originalLogger.Info,
		fatal:  originalLogger.Info,
		warn:   originalLogger.Info,
		with:   originalLogger.With,
		span:   span,
	}
	module.Warn("Test")
}
func TestFieldAdapter_AddArray(t *testing.T) {
	module := new(fieldAdapter)
	err := module.AddArray("", zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
		return nil
	}))
	require.NoError(t, err)
}
func TestFieldAdapter_AddObject(t *testing.T) {
	module := new(fieldAdapter)
	err := module.AddObject("", &ServiceContext{
		Service: "test",
	})
	require.NoError(t, err)
}
func TestFieldAdapter_AddReflected(t *testing.T) {
	module := new(fieldAdapter)
	err := module.AddReflected("", &ServiceContext{
		Service: "test",
	})
	require.NoError(t, err)
}

func TestFieldAdapter_AddBinary(t *testing.T) {
	module := new(fieldAdapter)
	module.AddBinary("", []byte{})
}

func TestFieldAdapter_AddComplex128(t *testing.T) {
	module := new(fieldAdapter)
	module.AddComplex128("", complex128(10))
}

func TestFieldAdapter_AddComplex64(t *testing.T) {
	module := new(fieldAdapter)
	module.AddComplex64("", complex64(10))
}

func TestFieldAdapter_AddDuration(t *testing.T) {
	module := new(fieldAdapter)
	module.AddDuration("", time.Hour)
}

func TestFieldAdapter_AddFloat32(t *testing.T) {
	module := new(fieldAdapter)
	module.AddFloat32("", 12.0)
}

func TestFieldAdapter_AddFloat64(t *testing.T) {
	module := new(fieldAdapter)
	module.AddFloat64("", 12.0)
}

func TestFieldAdapter_AddInt(t *testing.T) {
	module := new(fieldAdapter)
	module.AddInt("", 12)
}

func TestFieldAdapter_AddInt8(t *testing.T) {
	module := new(fieldAdapter)
	module.AddInt8("", 12)
}

func TestFieldAdapter_AddInt16(t *testing.T) {
	module := new(fieldAdapter)
	module.AddInt16("", 12)
}

func TestFieldAdapter_AddInt32(t *testing.T) {
	module := new(fieldAdapter)
	module.AddInt32("", 12)
}

func TestFieldAdapter_AddInt64(t *testing.T) {
	module := new(fieldAdapter)
	module.AddInt64("", 12)
}

func TestFieldAdapter_AddString(t *testing.T) {
	module := new(fieldAdapter)
	module.AddString("", "")
}

func TestFieldAdapter_AddByteString(t *testing.T) {
	module := new(fieldAdapter)
	module.AddByteString("", []byte("test"))
}

func TestFieldAdapter_AddBool(t *testing.T) {
	module := new(fieldAdapter)
	module.AddBool("", false)
}

func TestFieldAdapter_AddUint(t *testing.T) {
	module := new(fieldAdapter)
	module.AddUint("", 12)
}

func TestFieldAdapter_AddUint8(t *testing.T) {
	module := new(fieldAdapter)
	module.AddUint8("", 12)
}

func TestFieldAdapter_AddUint16(t *testing.T) {
	module := new(fieldAdapter)
	module.AddUint16("", 12)
}

func TestFieldAdapter_AddUint32(t *testing.T) {
	module := new(fieldAdapter)
	module.AddUint32("", 12)
}

func TestFieldAdapter_AddUint64(t *testing.T) {
	module := new(fieldAdapter)
	module.AddUint64("", 12)
}

func TestFieldAdapter_AddTime(t *testing.T) {
	module := new(fieldAdapter)
	module.AddTime("", time.Now())
}
