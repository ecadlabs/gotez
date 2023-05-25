package protocol

import (
	"fmt"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
	"github.com/ecadlabs/gotez/protocol/proto_016_PtMumbai"
)

type BlockInfoProtocolData interface {
	BlockHeader() *core.BlockHeader
	BlockMetadata() tz.Option[*core.BlockMetadataHeader]
}

type BlockInfo struct {
	Contents BlockInfoContents `tz:"dyn"`
}

type BlockInfoContents struct {
	ChainID      *tz.ChainID
	Hash         *tz.BlockHash
	ProtocolData BlockInfoProtocolData
}

func (info *BlockInfoContents) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	type part1 struct {
		ChainID *tz.ChainID
		Hash    *tz.BlockHash
	}

	type part2 struct {
		Level int32
		Proto core.Protocol
	}

	var p1 part1
	data, err = encoding.Decode(data, &p1, encoding.Ctx(ctx))
	if err != nil {
		return nil, err
	}

	info.ChainID = p1.ChainID
	info.Hash = p1.Hash

	var p2 part2
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
	default:
		return nil, fmt.Errorf("gotez: BlockInfoContents.DecodeTZ: unknown protocol version %d", p2.Proto)
	}

	return encoding.Decode(data, info.ProtocolData, encoding.Ctx(ctx))
}
