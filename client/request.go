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
type DelegateInfo = core.DelegateInfo
type BigUint = tz.BigUint
type OperationWithOptionalMetadata = core.OperationWithOptionalMetadata[latest.OperationWithOptionalMetadataContents]
type Constants = core.Constants
type BlockShellHeader = core.ShellHeader

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

type BlockRequest struct {
	Chain    string
	Block    string
	Metadata MetadataMode
	Protocol tz.Option[core.Protocol]
}

type ContractRequest struct {
	Chain string
	Block string
	ID    core.ContractID
}

type ContextRequest struct {
	Chain    string
	Block    string
	Protocol core.Protocol
}

type RunOperationRequest struct {
	Chain   string
	Block   string
	Payload *latest.RunOperationRequest
}

func newConstants(p core.Protocol) (Constants, error) { return protocol.NewConstants(p) }
