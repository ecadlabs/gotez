// Package client is a very limited Tezos RPC client library
package client

import (
	"context"

	tz "github.com/ecadlabs/gotez/v2"
	clientv2 "github.com/ecadlabs/gotez/v2/clientv2"
	"github.com/ecadlabs/gotez/v2/clientv2/block"
	"github.com/ecadlabs/gotez/v2/clientv2/monitor"
	"github.com/ecadlabs/gotez/v2/clientv2/utils"
	"github.com/ecadlabs/gotez/v2/protocol"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/latest"
)

type Logger = clientv2.Logger
type Error = clientv2.Error

type BlockInfo = protocol.BlockInfo
type BlockHeaderInfo = protocol.BlockHeaderInfo
type BigUint = tz.BigUint
type ChainID = tz.ChainID
type OperationWithOptionalMetadata = latest.OperationWithOptionalMetadata
type Constants = core.Constants
type BlockShellHeader = core.ShellHeader
type OperationHash = tz.OperationHash
type BlockProtocols = core.BlockProtocols
type BlockHash = tz.BlockHash

type MetadataMode = block.MetadataMode
type SimpleRequest = block.SimpleRequest
type BlockRequest = block.BlockRequest
type ContractRequest = block.ContractRequest
type ContextRequest = block.ContextRequest
type RunOperationRequest = block.RunOperationRequest
type InjectOperationRequest = utils.InjectOperationRequest
type InjectRequestPayload = utils.InjectRequestPayload
type BasicBlockInfo = block.BasicBlockInfo
type HeadsRequest = monitor.HeadsRequest
type Head = monitor.Head
type Flag = clientv2.Flag
type BootstrappedResponse = utils.BootstrappedResponse
type SyncState = utils.SyncState

const (
	SyncStateSynced   = utils.SyncStateSynced
	SyncStateUnsynced = utils.SyncStateUnsynced
	SyncStateStuck    = utils.SyncStateStuck
)

type Client clientv2.Client

func (client *Client) BlockHash(ctx context.Context, r *SimpleRequest) (*BlockHash, error) {
	return block.Hash(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) BlockProtocols(ctx context.Context, r *SimpleRequest) (*BlockProtocols, error) {
	return block.Protocols(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) BlockShellHeader(ctx context.Context, r *SimpleRequest) (*BlockShellHeader, error) {
	return block.ShellHeader(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) BlockHeader(ctx context.Context, r *BlockRequest) (BlockHeaderInfo, error) {
	return block.Header(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) Block(ctx context.Context, r *BlockRequest) (BlockInfo, error) {
	return block.Block(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) ContractBalance(ctx context.Context, r *ContractRequest) (BigUint, error) {
	return block.ContractBalance(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) ContractBalanceAndFrozenBonds(ctx context.Context, r *ContractRequest) (BigUint, error) {
	return block.ContractBalanceAndFrozenBonds(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) ContractCounter(ctx context.Context, r *ContractRequest) (BigUint, error) {
	return block.ContractCounter(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) RunOperation(ctx context.Context, r *RunOperationRequest) (*OperationWithOptionalMetadata, error) {
	return block.RunOperation(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) Constants(ctx context.Context, r *ContextRequest) (Constants, error) {
	return block.Constants(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) InjectOperation(ctx context.Context, r *InjectOperationRequest) (*OperationHash, error) {
	return utils.InjectOperation(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) Heads(ctx context.Context, r *HeadsRequest) (<-chan *Head, <-chan error, error) {
	return monitor.Heads(ctx, (*clientv2.Client)(client), r)
}

func (client *Client) IsBootstrapped(ctx context.Context, r *ChainID) (*BootstrappedResponse, error) {
	return utils.IsBootstrapped(ctx, (*clientv2.Client)(client), r)
}

// BasicBlockInfo returns hash and protocol of the block (usually head) to be used for sequent requests
func (client *Client) BasicBlockInfo(ctx context.Context, chain string, blockID string) (*BasicBlockInfo, error) {
	hash, err := block.Hash(ctx, (*clientv2.Client)(client), &block.SimpleRequest{
		Chain: chain,
		Block: blockID,
	})
	if err != nil {
		return nil, err
	}

	proto, err := block.Protocols(ctx, (*clientv2.Client)(client), &block.SimpleRequest{
		Chain: chain,
		Block: hash.String(),
	})
	if err != nil {
		return nil, err
	}

	return &BasicBlockInfo{
		Hash:     hash,
		Protocol: proto.Protocol,
	}, nil
}
