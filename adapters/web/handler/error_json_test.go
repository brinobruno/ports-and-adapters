package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerJsonError(t *testing.T) {
	msg := "hello Json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"hello Json"}`), result)
}
