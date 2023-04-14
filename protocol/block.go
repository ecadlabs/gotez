package protocol

import (
	"fmt"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/proto"
	"github.com/ecadlabs/gotez/protocol/proto_016_PtMumbai"
	"github.com/ecadlabs/gotez/protocol/shell"
)

type ProtocolBlockHeader interface {
	ProtocolBlockHeader()
}

type BlockProtocolData interface {
	BlockProtocolData()
}

type BlockHeader struct {
	shell.BlockHeader
	ProtocolData ProtocolBlockHeader
}

func (header *BlockHeader) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	data, err = encoding.Decode(data, &header.BlockHeader, encoding.Ctx(ctx))
	if err != nil {
		return nil, err
	}
	header.ProtocolData, err = getProtocolBlockHeaderImpl(header.BlockHeader.Proto)
	if err != nil {
		return nil, err
	}
	return encoding.Decode(data, header.ProtocolData, encoding.Ctx(ctx))
}

type BlockInfo struct {
	Contents BlockInfoContents `tz:"dyn"`
}

type BlockInfoContents struct {
	ChainID      *tz.ChainID
	Hash         *tz.BlockHash
	Header       BlockHeader `tz:"dyn"`
	ProtocolData BlockProtocolData
}

func (block *BlockInfoContents) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	type blockInfoPrefix struct {
		ChainID *tz.ChainID
		Hash    *tz.BlockHash
	}

	var pre blockInfoPrefix
	data, err = encoding.Decode(data, &pre, encoding.Ctx(ctx))
	block.ChainID = pre.ChainID
	block.Hash = pre.Hash
	if err != nil {
		return nil, err
	}
	data, err = encoding.Decode(data, &block.Header, encoding.Ctx(ctx), encoding.Dynamic())
	if err != nil {
		return nil, err
	}
	block.ProtocolData, err = getBlockProtocolDataImpl(block.Header.BlockHeader.Proto)
	if err != nil {
		return nil, err
	}
	return encoding.Decode(data, block.ProtocolData, encoding.Ctx(ctx))
}

func getProtocolBlockHeaderImpl(p proto.Protocol) (ProtocolBlockHeader, error) {
	switch p {
	case proto.Proto016PtMumbai:
		return new(proto_016_PtMumbai.ProtocolBlockHeader), nil
	default:
		return nil, fmt.Errorf("gotez: unknown protocol version %d", p)
	}
}

func getBlockProtocolDataImpl(p proto.Protocol) (BlockProtocolData, error) {
	switch p {
	case proto.Proto016PtMumbai:
		return new(proto_016_PtMumbai.BlockProtocolData), nil
	default:
		return nil, fmt.Errorf("gotez: unknown protocol version %d", p)
	}
}
