package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/stretchr/testify/require"
)

func TestDelegates(t *testing.T) {
	fileName := filepath.Join("test_data", "delegates.bin")
	buf, err := os.ReadFile(fileName)
	require.NoError(t, err)
	var out DelegatesList
	_, err = encoding.Decode(buf, &out, encoding.Dynamic())
	require.NoError(t, err)
}
