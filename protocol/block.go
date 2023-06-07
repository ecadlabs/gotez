package protocol

import (
	"fmt"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

type BlockInfoProtocolData interface {
	GetHeader() core.BlockHeader
	GetMetadata() tz.Option[core.BlockMetadata]
	GetOperations() [][]core.OperationsGroup
}

type BlockInfo struct {
	ChainID      *tz.ChainID           `json:"chain_id"`
	Hash         *tz.BlockHash         `json:"hash"`
	ProtocolData BlockInfoProtocolData `json:"protocol_data"`
}

type blockInfoPreamble struct {
	ChainID *tz.ChainID
	Hash    *tz.BlockHash
}

func (info *BlockInfo) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	var p1 blockInfoPreamble
	data, err = encoding.Decode(data, &p1, encoding.Ctx(ctx))
	if err != nil {
		return nil, err
	}

	info.ChainID = p1.ChainID
	info.Hash = p1.Hash

	var p2 core.ShellHeader
	if _, err = encoding.Decode(data, &p2, encoding.Ctx(ctx), encoding.Dynamic()); err != nil {
		return nil, err
	}

	p, ok := ctx.Get(core.ProtocolVersionCtxKey).(core.Protocol)
	if !ok {
		p = p2.Proto
	}

	switch p {
	case core.Proto016PtMumbai:
		info.ProtocolData = new(proto_016_PtMumbai.BlockInfoProtocolData)
	case core.Proto015PtLimaPt:
		info.ProtocolData = new(proto_015_PtLimaPt.BlockInfoProtocolData)
	case core.Proto014PtKathma:
		info.ProtocolData = new(proto_014_PtKathma.BlockInfoProtocolData)
	case core.Proto013PtJakart:
		info.ProtocolData = new(proto_013_PtJakart.BlockInfoProtocolData)
	case core.Proto012Psithaca:
		info.ProtocolData = new(proto_012_Psithaca.BlockInfoProtocolData)

	default:
		return nil, fmt.Errorf("gotez: BlockInfoContents.DecodeTZ: unknown protocol version %d", p2.Proto)
	}

	return encoding.Decode(data, info.ProtocolData, encoding.Ctx(ctx))
}

type BlockHeaderInfo struct {
	ChainID      *tz.ChainID      `json:"chain_id"`
	Hash         *tz.BlockHash    `json:"hash"`
	ProtocolData core.BlockHeader `json:"protocol_data"`
}

func (info *BlockHeaderInfo) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	var p1 blockInfoPreamble
	data, err = encoding.Decode(data, &p1, encoding.Ctx(ctx))
	if err != nil {
		return nil, err
	}

	info.ChainID = p1.ChainID
	info.Hash = p1.Hash

	var p2 core.ShellHeader
	if _, err = encoding.Decode(data, &p2, encoding.Ctx(ctx)); err != nil {
		return nil, err
	}

	p, ok := ctx.Get(core.ProtocolVersionCtxKey).(core.Protocol)
	if !ok {
		p = p2.Proto
	}

	switch p {
	case core.Proto016PtMumbai:
		info.ProtocolData = new(proto_016_PtMumbai.BlockHeader)
	case core.Proto015PtLimaPt:
		info.ProtocolData = new(proto_015_PtLimaPt.BlockHeader)
	case core.Proto014PtKathma:
		info.ProtocolData = new(proto_014_PtKathma.BlockHeader)
	case core.Proto013PtJakart:
		info.ProtocolData = new(proto_013_PtJakart.BlockHeader)
	case core.Proto012Psithaca:
		info.ProtocolData = new(proto_012_Psithaca.BlockHeader)

	default:
		return nil, fmt.Errorf("gotez: BlockHeaderInfoContents.DecodeTZ: unknown protocol version %d", p2.Proto)
	}

	return encoding.Decode(data, info.ProtocolData, encoding.Ctx(ctx))
}
