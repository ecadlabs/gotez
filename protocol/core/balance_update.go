package core

import (
	"github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type BalanceUpdateOrigin interface {
	BalanceUpdateOrigin() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BalanceUpdateOrigin]{
		Variants: encoding.Variants[BalanceUpdateOrigin]{
			0: BalanceUpdateOriginBlockApplication{},
			1: BalanceUpdateOriginProtocolMigration{},
			2: BalanceUpdateOriginSubsidy{},
			3: BalanceUpdateOriginSimulation{},
			4: (*BalanceUpdateOriginDelayedOperation)(nil),
		},
	})
}

//json:origin=BalanceUpdateOrigin()
type BalanceUpdateOriginBlockApplication struct{}

func (BalanceUpdateOriginBlockApplication) BalanceUpdateOrigin() string { return "block_application" }

//json:origin=BalanceUpdateOrigin()
type BalanceUpdateOriginProtocolMigration struct{}

func (BalanceUpdateOriginProtocolMigration) BalanceUpdateOrigin() string { return "protocol_migration" }

//json:origin=BalanceUpdateOrigin()
type BalanceUpdateOriginSubsidy struct{}

func (BalanceUpdateOriginSubsidy) BalanceUpdateOrigin() string { return "subsidy" }

//json:origin=BalanceUpdateOrigin()
type BalanceUpdateOriginSimulation struct{}

func (BalanceUpdateOriginSimulation) BalanceUpdateOrigin() string { return "simulation" }

//json:origin=BalanceUpdateOrigin()
type BalanceUpdateOriginDelayedOperation struct {
	DelayedOperationHash gotez.OperationHash `json:"delayed_operation_hash"`
}

func (*BalanceUpdateOriginDelayedOperation) BalanceUpdateOrigin() string { return "delayed_operation" }

// not present in the binary protocol
type BalanceUpdateKind int

func (k BalanceUpdateKind) String() string {
	switch k {
	case BalanceUpdateKindContract:
		return "contract"
	case BalanceUpdateKindAccumulator:
		return "accumulator"
	case BalanceUpdateKindFreezer:
		return "freezer"
	case BalanceUpdateKindMinted:
		return "minted"
	case BalanceUpdateKindBurned:
		return "burned"
	case BalanceUpdateKindCommitment:
		return "commitment"
	case BalanceUpdateKindStaking:
		return "staking"
	default:
		return "<unknown>"
	}
}

func (k BalanceUpdateKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

const (
	BalanceUpdateKindContract BalanceUpdateKind = iota
	BalanceUpdateKindAccumulator
	BalanceUpdateKindFreezer
	BalanceUpdateKindMinted
	BalanceUpdateKindBurned
	BalanceUpdateKindCommitment
	BalanceUpdateKindStaking
)

type BalanceUpdate interface {
	GetContents() BalanceUpdateContents
	GetChange() int64
	GetOrigin() BalanceUpdateOrigin
}

type BalanceUpdateContents interface {
	BalanceUpdateCategory() string
	BalanceUpdateKind() BalanceUpdateKind
}

type WithBalanceUpdates interface {
	GetBalanceUpdates() []BalanceUpdate
}

type BalanceUpdateContract interface {
	BalanceUpdateContents
	GetContract() ContractID
}
