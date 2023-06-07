package core

import "strconv"

type BalanceUpdateOrigin uint8

const (
	BalanceUpdateOriginBlockApplication BalanceUpdateOrigin = iota
	BalanceUpdateOriginProtocolMigration
	BalanceUpdateOriginSubsidy
	BalanceUpdateOriginSimulation
)

func (o BalanceUpdateOrigin) String() string {
	switch o {
	case BalanceUpdateOriginBlockApplication:
		return "block_application"
	case BalanceUpdateOriginProtocolMigration:
		return "protocol_migration"
	case BalanceUpdateOriginSubsidy:
		return "subsidy"
	case BalanceUpdateOriginSimulation:
		return "simulation"
	default:
		return strconv.FormatInt(int64(o), 10)
	}
}

func (o BalanceUpdateOrigin) MarshalText() (text []byte, err error) {
	return []byte(o.String()), nil
}

// not present in the binary protocol
type BalanceUpdateKind int

func (k BalanceUpdateKind) String() string {
	switch k {
	case BalanceUpdateContract:
		return "contract"
	case BalanceUpdateAccumulator:
		return "accumulator"
	case BalanceUpdateFreezer:
		return "freezer"
	case BalanceUpdateMinted:
		return "minted"
	case BalanceUpdateBurned:
		return "burned"
	case BalanceUpdateCommitment:
		return "commitment"
	default:
		return "<unknown>"
	}
}

func (k BalanceUpdateKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

const (
	BalanceUpdateContract BalanceUpdateKind = iota
	BalanceUpdateAccumulator
	BalanceUpdateFreezer
	BalanceUpdateMinted
	BalanceUpdateBurned
	BalanceUpdateCommitment
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

type BalanceUpdates interface {
	GetBalanceUpdates() []BalanceUpdate
}
