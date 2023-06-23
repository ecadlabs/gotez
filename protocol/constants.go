package protocol

import (
	"fmt"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

func NewConstants(proto *tz.ProtocolHash) (constants core.Constants, err error) {
	switch *proto {
	case core.Proto016PtMumbai:
		constants = new(proto_016_PtMumbai.Constants)
	case core.Proto015PtLimaPt:
		constants = new(proto_015_PtLimaPt.Constants)
	case core.Proto014PtKathma:
		constants = new(proto_014_PtKathma.Constants)
	case core.Proto013PtJakart:
		constants = new(proto_013_PtJakart.Constants)
	case core.Proto012Psithaca:
		constants = new(proto_012_Psithaca.Constants)
	default:
		return nil, fmt.Errorf("gotez: NewConstants: unknown protocol version %d", proto)
	}
	return
}
