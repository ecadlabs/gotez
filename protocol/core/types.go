package core

//go:generate go run ../../cmd/genmarshaller.go

import (
	"encoding/json"
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
)

type OperationContents interface {
	OperationKind() string
}

type OperationContentsAndResult interface {
	OperationContents
	GetMetadata() any
}

type ManagerOperationMetadata interface {
	BalanceUpdates
	GetResult() ManagerOperationResult
	GetInternalOperationResults() []InternalOperationResult
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

func (b BallotKind) String() string {
	switch b {
	case BallotYay:
		return "yay"
	case BallotNay:
		return "nay"
	case BallotPass:
		return "pass"
	default:
		return strconv.FormatInt(int64(b), 10)
	}
}

func (b BallotKind) MarshalText() (text []byte, err error) {
	return []byte(b.String()), nil
}

type ContractID interface {
	tz.Base58Encoder
	ContractID()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ContractID]{
		Variants: encoding.Variants[ContractID]{
			0: (*ImplicitContract)(nil),
			1: (*OriginatedContract)(nil),
		},
	})
}

type OriginatedContract struct {
	*tz.ContractHash
	Padding uint8
}

type ImplicitContract struct {
	tz.PublicKeyHash
}

func (c *ImplicitContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.PublicKeyHash)
}

func (*ImplicitContract) ContractID() {}

type OriginatedContractID interface {
	tz.Base58Encoder
	OriginatedContractID()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginatedContractID]{
		Variants: encoding.Variants[OriginatedContractID]{
			1: (*OriginatedContract)(nil),
		},
	})
}

type TransactionDestination interface {
	tz.Base58Encoder
	TransactionDestination()
}

func (*OriginatedContract) ContractID()             {}
func (*OriginatedContract) OriginatedContractID()   {}
func (*OriginatedContract) TransactionDestination() {}
func (c *OriginatedContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ContractHash)
}
func (*ImplicitContract) TransactionDestination() {}

type Signed interface {
	GetSignature() (tz.Signature, error)
}

type ManagerOperation interface {
	OperationContents
	GetSource() tz.PublicKeyHash
	GetFee() tz.BigUint
	GetCounter() tz.BigUint
	GetGasLimit() tz.BigUint
	GetStorageLimit() tz.BigUint
}

type Transaction interface {
	OperationContents
	ManagerOperation
	GetAmount() tz.BigUint
	GetDestination() ContractID
	GetParameters() tz.Option[Parameters]
}

type Parameters interface {
	GetEntrypoint() string
	GetValue() expression.Expression
}
