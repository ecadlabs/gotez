package protocol

import (
	"fmt"

	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

func NewDelegateInfo(proto core.Protocol) (delegate core.DelegateInfo, err error) {
	switch proto {
	case core.Proto016PtMumbai:
		delegate = new(proto_016_PtMumbai.DelegateInfo)
	case core.Proto015PtLimaPt:
		delegate = new(proto_015_PtLimaPt.DelegateInfo)
	case core.Proto014PtKathma:
		delegate = new(proto_014_PtKathma.DelegateInfo)
	case core.Proto013PtJakart:
		delegate = new(proto_013_PtJakart.DelegateInfo)
	case core.Proto012Psithaca:
		delegate = new(proto_012_Psithaca.DelegateInfo)
	default:
		return nil, fmt.Errorf("gotez: NewDelegateInfo: unknown protocol version %d", proto)
	}
	return
}
