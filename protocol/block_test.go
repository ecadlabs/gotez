package protocol

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type protoTestData struct {
	proto  proto.Protocol
	blocks []string
}

var testData = []protoTestData{
	{
		proto: proto.Proto016PtMumbai,
		blocks: []string{
			"3279466",
		},
	},
}

func TestBlock(t *testing.T) {
	for _, protoData := range testData {
		t.Run(protoData.proto.Name(), func(t *testing.T) {
			for _, block := range protoData.blocks {
				t.Run(block, func(t *testing.T) {
					fileName := filepath.Join("test_data", protoData.proto.Name(), block+".bin")
					buf, err := os.ReadFile(fileName)
					require.NoError(t, err)
					var block BlockInfo
					_, err = encoding.Decode(buf, &block)
					if !assert.NoError(t, err) {
						if err, ok := err.(*encoding.Error); ok {
							fmt.Println(err.Path)
						}
					}
				})
			}
		})
	}
}
