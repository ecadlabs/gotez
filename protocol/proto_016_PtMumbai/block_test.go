package proto_016_PtMumbai

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ecadlabs/gotez/encoding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = []string{
	"181313",
	"298135",
	"298154",
	"327682",
	"332053",
	"332064",
	"332066",
	"332075",
	"332090",
	"332091",
	"332093",
	"332470",
	"332530",
	"332534",
	"39524",
	"41157",
	"41821",
}

func TestBlock(t *testing.T) {
	for _, block := range testData {
		t.Run(block, func(t *testing.T) {
			fileName := filepath.Join("test_data", block+".bin")
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
}
