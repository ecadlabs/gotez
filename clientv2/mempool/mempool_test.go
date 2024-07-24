package mempool_test

import (
	_ "embed"
	"testing"

	"github.com/ecadlabs/gotez/v2/clientv2/mempool"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/stretchr/testify/assert"
)

//go:embed test_data/mempool.bin
var mempoolSrc []byte

func TestPendingOperations(t *testing.T) {
	var out mempool.PendingOperationsResponse
	_, err := encoding.Decode(mempoolSrc, &out, encoding.Dynamic())
	assert.NoError(t, err)
}
