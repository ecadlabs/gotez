package proto_005_PsBABY5H

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
)

type Transaction struct {
	ManagerOperation
	Amount      tz.BigUint
	Destination tz.ContractID
	Parameters  tz.Option[Parameters]
}

type Parameters struct {
	Entrypoint Entrypoint
	Value      expression.Expression `tz:"dyn"`
}

type Entrypoint interface {
	Entrypoint()
}

type EpDefault struct{}
type EpRoot struct{}
type EpDo struct{}
type EpSetDelegate struct{}
type EpRemoveDelegate struct{}
type EpNamed struct {
	tz.String
}

func (EpDefault) Entrypoint()        {}
func (EpRoot) Entrypoint()           {}
func (EpDo) Entrypoint()             {}
func (EpSetDelegate) Entrypoint()    {}
func (EpRemoveDelegate) Entrypoint() {}
func (EpNamed) Entrypoint()          {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Entrypoint]{
		Variants: encoding.Variants[Entrypoint]{
			0:   EpDefault{},
			1:   EpRoot{},
			2:   EpDo{},
			3:   EpSetDelegate{},
			4:   EpRemoveDelegate{},
			255: EpNamed{},
		},
	})
}

type ManagerOperation struct {
	Source       tz.PublicKeyHash
	Fee          tz.BigUint
	Counter      tz.BigUint
	GasLimit     tz.BigUint
	StorageLimit tz.BigUint
}

func (*Transaction) OperationKind() string { return "transaction" }

type Script struct {
	Code    expression.Expression `tz:"dyn"`
	Storage expression.Expression `tz:"dyn"`
}

type Origination struct {
	ManagerOperation
	Balance  tz.BigUint
	Delegate tz.Option[tz.PublicKeyHash]
	Script   Script
}

func (*Origination) OperationKind() string { return "origination" }

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash]
}

func (*Delegation) OperationKind() string { return "delegation" }

type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*Reveal) OperationKind() string { return "reveal" }

type SeedNonceRevelation struct {
	Level int32
	Nonce *[tz.SeedNonceBytesLen]byte
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }
