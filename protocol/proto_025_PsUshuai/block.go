package proto_025_PsUshuai

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type BlockInfo struct {
	ChainID    *tz.ChainID                          `json:"chain_id"`
	Hash       *tz.BlockHash                        `json:"hash"`
	Header     BlockHeader                          `tz:"dyn" json:"header"`
	Metadata   tz.Option[BlockMetadata]             `json:"metadata"`
	Operations []core.OperationsList[GroupContents] `tz:"dyn" json:"operations"`
}

func (block *BlockInfo) GetChainID() *tz.ChainID     { return block.ChainID }
func (block *BlockInfo) GetHash() *tz.BlockHash      { return block.Hash }
func (block *BlockInfo) GetHeader() core.BlockHeader { return &block.Header }
func (block *BlockInfo) GetMetadata() tz.Option[core.BlockMetadata] {
	if m, ok := block.Metadata.CheckUnwrapPtr(); ok {
		return tz.Some[core.BlockMetadata](m)
	}
	return tz.None[core.BlockMetadata]()
}

func (block *BlockInfo) GetOperations() [][]core.OperationsGroup {
	out := make([][]core.OperationsGroup, len(block.Operations))
	for i, list := range block.Operations {
		out[i] = list.GetGroups()
	}
	return out
}
