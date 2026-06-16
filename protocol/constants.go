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
	"github.com/ecadlabs/gotez/v2/protocol/proto_017_PtNairob"
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
	"github.com/ecadlabs/gotez/v2/protocol/proto_019_PtParisB"
	"github.com/ecadlabs/gotez/v2/protocol/proto_021_PsQuebec"
	"github.com/ecadlabs/gotez/v2/protocol/proto_022_PsRiotum"
	"github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"
	"github.com/ecadlabs/gotez/v2/protocol/proto_024_PtTALLiN"
	"github.com/ecadlabs/gotez/v2/protocol/proto_025_PsUshuai"
)

func NewConstants(proto *tz.ProtocolHash) (constants core.Constants, err error) {
	switch *proto {
	case core.Proto025PsUshuai:
		constants = new(proto_025_PsUshuai.Constants)
	case core.Proto024PtTALLiN:
		constants = new(proto_024_PtTALLiN.Constants)
	case core.Proto023PtSeouLo:
		constants = new(proto_023_PtSeouLo.Constants)
	case core.Proto022PsRiotum:
		constants = new(proto_022_PsRiotum.Constants)
	case core.Proto021PsQuebec:
		constants = new(proto_021_PsQuebec.Constants)
	case core.Proto019PtParisB, core.Proto020PsParisC:
		constants = new(proto_019_PtParisB.Constants)
	case core.Proto018Proxford:
		constants = new(proto_018_Proxford.Constants)
	case core.Proto017PtNairob:
		constants = new(proto_017_PtNairob.Constants)
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
