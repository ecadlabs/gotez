package core

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type OperationContents interface {
	OperationKind() string
}

type OperationContentsAndResult interface {
	OperationContentsAndResult()
	OperationContents() OperationContents
}

type InternalOperationResult interface {
	OperationContents
	InternalOperationResult()
}

type BalanceUpdateKind interface {
	BalanceUpdateKind() string
}

type Bytes struct {
	Bytes []byte `tz:"dyn"`
}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

type ContractID interface {
	tz.Base58Encoder
	ContractID()
}

type OriginatedContract struct {
	*tz.ContractHash
	Padding uint8
}

type ImplicitContract struct {
	tz.PublicKeyHash
}

func (*ImplicitContract) ContractID() {}

type OriginatedContractID interface {
	tz.Base58Encoder
	OriginatedContractID()
}

func (*OriginatedContract) ContractID()             {}
func (*OriginatedContract) OriginatedContractID()   {}
func (*OriginatedContract) TransactionDestination() {}
func (*ImplicitContract) TransactionDestination()   {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ContractID]{
		Variants: encoding.Variants[ContractID]{
			0: (*ImplicitContract)(nil),
			1: (*OriginatedContract)(nil),
		},
	})
	encoding.RegisterEnum(&encoding.Enum[OriginatedContractID]{
		Variants: encoding.Variants[OriginatedContractID]{
			1: (*OriginatedContract)(nil),
		},
	})
}

type BalanceUpdate interface {
	BalanceUpdate()
}

type Signed interface {
	GetSignature() (tz.Signature, error)
}

type GroupContents interface {
	Signed
	GroupContents()
	Operations() []OperationContents
}

type OperationWithOptionalMetadataContents interface {
	Signed
	OperationWithOptionalMetadataContents()
	Operations() []OperationContents
}
