package protocol

import (
	"fmt"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_016_PtMumbai"
)

type BlockHeader struct {
	core.BlockHeader
	ProtocolData ProtocolBlockHeader
}

func (header *BlockHeader) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	data, err = encoding.Decode(data, &header.BlockHeader, encoding.Ctx(ctx))
	if err != nil {
		return nil, err
	}

	p, ok := ctx.Get(core.ProtocolVersionCtxKey).(core.Protocol)
	if !ok {
		p = header.Proto
	}

	switch p {
	//case proto.Proto015PtLimaPt:
	//	header.ProtocolData = new(proto_015_PtLimaPt.ProtocolBlockHeader)
	case core.Proto016PtMumbai:
		header.ProtocolData = new(proto_016_PtMumbai.ProtocolBlockHeader)
	default:
		return nil, fmt.Errorf("gotez: BlockHeader.DecodeTZ: unknown protocol version %d", header.Proto)
	}

	return encoding.Decode(data, header.ProtocolData, encoding.Ctx(ctx))
}

type ProtocolBlockHeader interface {
	ProtocolBlockHeader()
}

type BlockInfo struct {
	Contents BlockInfoContents `tz:"dyn"`
}

type BlockInfoContents struct {
	ChainID    *tz.ChainID
	Hash       *tz.BlockHash
	Header     BlockHeader `tz:"dyn"`
	Metadata   tz.Option[BlockMetadata]
	Operations []OperationsList `tz:"dyn"`
}

func (header *BlockInfoContents) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	type part1 struct {
		ChainID *tz.ChainID
		Hash    *tz.BlockHash
		Header  BlockHeader `tz:"dyn"`
	}
	type part2 struct {
		Metadata   tz.Option[BlockMetadata]
		Operations []OperationsList `tz:"dyn"`
	}
	var (
		p1 part1
		p2 part2
	)
	data, err = encoding.Decode(data, &p1, encoding.Ctx(ctx))
	header.ChainID = p1.ChainID
	header.Hash = p1.Hash
	header.Header = p1.Header
	if err != nil {
		return nil, err
	}
	if ctx.Get(core.ProtocolVersionCtxKey) == nil {
		ctx = ctx.Set(core.ProtocolVersionCtxKey, p1.Header.Proto)
	}

	data, err = encoding.Decode(data, &p2, encoding.Ctx(ctx))
	header.Metadata = p2.Metadata
	header.Operations = p2.Operations
	return data, err
}

type BlockMetadata struct {
	BlockMetadataContents `tz:"dyn"`
}

type BlockMetadataContents interface {
	BlockMetadataContents()
}

func init() {
	encoding.RegisterType(func(data []byte, ctx *encoding.Context) (BlockMetadataContents, []byte, error) {
		p, ok := ctx.Get(core.ProtocolVersionCtxKey).(core.Protocol)
		if !ok {
			return nil, nil, fmt.Errorf("gotez: protocol version must be passed to the decoder chain")
		}

		var out BlockMetadataContents
		switch p {
		case core.Proto016PtMumbai:
			out = new(proto_016_PtMumbai.BlockMetadataContents)
		default:
			return nil, nil, fmt.Errorf("gotez: unknown protocol version %d", p)
		}

		data, err := encoding.Decode(data, out, encoding.Ctx(ctx))
		return out, data, err
	})
}
