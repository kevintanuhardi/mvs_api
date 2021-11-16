package tracing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTracing(t *testing.T) {
	tracer, closer, err := Init("dummyService")
	defer closer.Close()
	require.NoError(t, err)
	require.NotNil(t, tracer)
	require.NotNil(t, closer)
}
