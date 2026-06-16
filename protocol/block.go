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

type BlockInfo interface {
	GetChainID() *tz.ChainID
	GetHash() *tz.BlockHash
	GetHeader() core.BlockHeader
	GetMetadata() tz.Option[core.BlockMetadata]
	GetOperations() [][]core.OperationsGroup
}

type BlockHeaderInfo interface {
	GetChainID() *tz.ChainID
	GetHash() *tz.BlockHash
	core.BlockHeader
}

func NewBlockInfo(proto *tz.ProtocolHash) (BlockInfo, error) {
	switch *proto {
	case core.Proto025PsUshuai:
		return new(proto_025_PsUshuai.BlockInfo), nil
	case core.Proto024PtTALLiN:
		return new(proto_024_PtTALLiN.BlockInfo), nil
	case core.Proto023PtSeouLo:
		return new(proto_023_PtSeouLo.BlockInfo), nil
	case core.Proto022PsRiotum:
		return new(proto_022_PsRiotum.BlockInfo), nil
	case core.Proto021PsQuebec:
		return new(proto_021_PsQuebec.BlockInfo), nil
	case core.Proto019PtParisB, core.Proto020PsParisC:
		return new(proto_019_PtParisB.BlockInfo), nil
	case core.Proto018Proxford:
		return new(proto_018_Proxford.BlockInfo), nil
	case core.Proto017PtNairob:
		return new(proto_017_PtNairob.BlockInfo), nil
	case core.Proto016PtMumbai:
		return new(proto_016_PtMumbai.BlockInfo), nil
	case core.Proto015PtLimaPt:
		return new(proto_015_PtLimaPt.BlockInfo), nil
	case core.Proto014PtKathma:
		return new(proto_014_PtKathma.BlockInfo), nil
	case core.Proto013PtJakart:
		return new(proto_013_PtJakart.BlockInfo), nil
	case core.Proto012Psithaca:
		return new(proto_012_Psithaca.BlockInfo), nil
	default:
		return nil, fmt.Errorf("gotez: NewBlockInfo: unknown protocol %v", proto)
	}
}

func NewBlockHeaderInfo(proto *tz.ProtocolHash) (BlockHeaderInfo, error) {
	switch *proto {
	case core.Proto025PsUshuai:
		return new(proto_025_PsUshuai.BlockHeaderInfo), nil
	case core.Proto024PtTALLiN:
		return new(proto_024_PtTALLiN.BlockHeaderInfo), nil
	case core.Proto023PtSeouLo:
		return new(proto_023_PtSeouLo.BlockHeaderInfo), nil
	case core.Proto022PsRiotum:
		return new(proto_022_PsRiotum.BlockHeaderInfo), nil
	case core.Proto021PsQuebec:
		return new(proto_021_PsQuebec.BlockHeaderInfo), nil
	case core.Proto019PtParisB, core.Proto020PsParisC:
		return new(proto_019_PtParisB.BlockHeaderInfo), nil
	case core.Proto018Proxford:
		return new(proto_018_Proxford.BlockHeaderInfo), nil
	case core.Proto017PtNairob:
		return new(proto_017_PtNairob.BlockHeaderInfo), nil
	case core.Proto016PtMumbai:
		return new(proto_016_PtMumbai.BlockHeaderInfo), nil
	case core.Proto015PtLimaPt:
		return new(proto_015_PtLimaPt.BlockHeaderInfo), nil
	case core.Proto014PtKathma:
		return new(proto_014_PtKathma.BlockHeaderInfo), nil
	case core.Proto013PtJakart:
		return new(proto_013_PtJakart.BlockHeaderInfo), nil
	case core.Proto012Psithaca:
		return new(proto_012_Psithaca.BlockHeaderInfo), nil
	default:
		return nil, fmt.Errorf("gotez: NewBlockHeaderInfo: unknown protocol %v", proto)
	}
}
