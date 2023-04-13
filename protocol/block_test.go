package protocol

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ecadlabs/gotez/encoding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var blocks = []string{
	"332534",
	"332090",
	"332530",
	"332075",
	"332470",
	"332053",
	"332066",
	"332091",
	"327682",
	"332064",
	"298154",
	"298135",
	"41157",
	"39524",
	"41821",
	"332093",
}

func TestBlock(t *testing.T) {
	for _, block := range blocks {
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
			//fmt.Printf("%# v\n", pretty.Formatter(&block, pretty.OptStringer(true), pretty.OptTextMarshaler(true), pretty.OptMaxDepth(20)))
		})
	}
}
