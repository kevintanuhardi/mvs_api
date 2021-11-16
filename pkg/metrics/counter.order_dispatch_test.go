package metrics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDispatchCounter(t *testing.T) {
	require.NotNil(t, OrderDispatchCounter)
}
