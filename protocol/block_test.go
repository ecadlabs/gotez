package protocol

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type blockTestData struct {
	name         string
	forceVersion bool
}

type protoTestData struct {
	proto  core.Protocol
	blocks []blockTestData
}

var testData = []protoTestData{
	{
		proto: core.Proto016PtMumbai,
		blocks: []blockTestData{
			{"3279466", false},
			{"181313", true},
			{"298135", true},
			{"298154", true},
			{"327682", true},
			{"332053", true},
			{"332064", true},
			{"332066", true},
			{"332075", true},
			{"332090", true},
			{"332091", true},
			{"332093", true},
			{"332470", true},
			{"332530", true},
			{"332534", true},
			{"39524", true},
			{"41157", true},
			{"41821", true},
		},
	},
}

func TestBlock(t *testing.T) {
	for _, protoData := range testData {
		t.Run(protoData.proto.Name(), func(t *testing.T) {
			for _, block := range protoData.blocks {
				t.Run(block.name, func(t *testing.T) {
					fileName := filepath.Join("test_data", protoData.proto.Name(), block.name+".bin")
					buf, err := os.ReadFile(fileName)
					require.NoError(t, err)
					var out BlockInfo
					ctx := encoding.NewContext()
					if block.forceVersion {
						ctx = ctx.Set(core.ProtocolVersionCtxKey, protoData.proto)
					}
					_, err = encoding.Decode(buf, &out, encoding.Ctx(ctx))
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
