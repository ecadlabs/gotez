package monitor

import (
	"github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

//go:generate go run generate.go

type HeadsRequest struct {
	Chain        string
	Protocol     *gotez.ProtocolHash
	NextProtocol *gotez.ProtocolHash
}

type Head struct {
	Hash *gotez.BlockHash `json:"hash"`
	core.ShellHeader
	ProtocolData []byte `json:"protocol_data"` // not dyn, takes the rest
}
