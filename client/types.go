package client

//go:generate go run generate.go

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/latest"
)

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
	Protocol *tz.ProtocolHash
}

type ContractRequest struct {
	Chain string
	Block string
	ID    core.ContractID
}

type ContextRequest struct {
	Chain    string
	Block    string
	Protocol *tz.ProtocolHash
}

type RunOperationRequest struct {
	Chain   string
	Block   string
	Payload *latest.RunOperationRequest
}

type InjectOperationRequest struct {
	Chain   string
	Async   Flag
	Payload *InjectRequestPayload
}

type InjectRequestPayload struct {
	Contents []byte `tz:"dyn"`
}

type BasicBlockInfo struct {
	Hash     *tz.BlockHash
	Protocol *tz.ProtocolHash
}

type HeadsRequest struct {
	Chain        string
	Protocol     *tz.ProtocolHash
	NextProtocol *tz.ProtocolHash
}

type Head struct {
	Hash *tz.BlockHash `json:"hash"`
	core.ShellHeader
	ProtocolData []byte `json:"protocol_data"` // not dyn, takes the rest
}

type Flag bool

func newConstants(p *tz.ProtocolHash) (Constants, error) { return protocol.NewConstants(p) }
func newBlockInfo(p *tz.ProtocolHash) (BlockInfo, error) { return protocol.NewBlockInfo(p) }
func newBlockHeaderInfo(p *tz.ProtocolHash) (BlockHeaderInfo, error) {
	return protocol.NewBlockHeaderInfo(p)
}
