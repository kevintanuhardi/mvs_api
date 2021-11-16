package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceContextClone(t *testing.T) {
	module := &ServiceContext{}
	resp := module.Clone()
	require.Equal(t, resp.Service, module.Service)
	require.Equal(t, resp.Version, module.Version)
}
