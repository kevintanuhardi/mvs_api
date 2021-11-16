package response

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	err := NewError("Error", 100)
	require.Equal(t, err.Error(), "Error with code 100")
}
