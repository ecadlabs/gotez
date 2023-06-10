package client

//go:generate go run generate.go

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type BlockInfo = protocol.BlockInfo
type BlockHeaderInfo = protocol.BlockHeaderInfo
type DelegateInfo = core.DelegateInfo
type BigUint = tz.BigUint

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
}

type DelegateRequest struct {
	Chain    string
	Block    string
	PKH      tz.PublicKeyHash
	Protocol core.Protocol
}

func newDelegateInfo(proto core.Protocol) (DelegateInfo, error) {
	return protocol.NewDelegateInfo(proto)
}

type ContractRequest struct {
	Chain    string
	Block    string
	ID       core.ContractID
	Protocol core.Protocol
}
