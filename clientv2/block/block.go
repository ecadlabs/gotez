package block

import (
	"context"

	"github.com/ecadlabs/gotez/v2"
	client "github.com/ecadlabs/gotez/v2/clientv2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/latest"
)

//go:generate go run generate.go

type MetadataMode int

const (
	MetadataDefault MetadataMode = iota
	MetadataAlways
	MetadataNever
)

func (m MetadataMode) String() string {
	switch m {
	case MetadataAlways:
		return "always"
	case MetadataNever:
		return "never"
	default:
		return "default"
	}
}

type SimpleRequest struct {
	Chain string
	Block string
}

type BlockRequest struct {
	Chain    string
	Block    string
	Metadata MetadataMode
	Protocol *gotez.ProtocolHash
}

type ContractRequest struct {
	Chain string
	Block string
	ID    core.ContractID
}

type ContextRequest struct {
	Chain    string
	Block    string
	Protocol *gotez.ProtocolHash
}

type RunOperationRequest struct {
	Chain   string
	Block   string
	Payload *latest.RunOperationRequest
}

type BasicBlockInfo struct {
	Hash         *gotez.BlockHash
	Protocol     *gotez.ProtocolHash
	NextProtocol *gotez.ProtocolHash
}

// BasicInfo returns hash and protocol of the block (usually head) to be used for sequent requests
func BasicInfo(ctx context.Context, cl *client.Client, chain string, block string) (*BasicBlockInfo, error) {
	hash, err := Hash(ctx, cl, &SimpleRequest{
		Chain: chain,
		Block: block,
	})
	if err != nil {
		return nil, err
	}

	proto, err := Protocols(ctx, cl, &SimpleRequest{
		Chain: chain,
		Block: hash.String(),
	})
	if err != nil {
		return nil, err
	}

	return &BasicBlockInfo{
		Hash:         hash,
		Protocol:     proto.Protocol,
		NextProtocol: proto.NextProtocol,
	}, nil
}
