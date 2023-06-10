package core

//go:generate go run ../../cmd/genmarshaller.go

import (
	"errors"
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/b58/base58"
	"github.com/ecadlabs/gotez/v2/b58/prefix"
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
			0: ImplicitContract{},
			1: OriginatedContract{},
		},
	})
}

type OriginatedContract struct {
	*tz.ContractHash
	Padding uint8
}

func (OriginatedContract) ContractID()             {}
func (OriginatedContract) OriginatedContractID()   {}
func (OriginatedContract) TransactionDestination() {}

type ImplicitContract struct {
	tz.PublicKeyHash
}

func (ImplicitContract) ContractID()             {}
func (ImplicitContract) TransactionDestination() {}

type OriginatedContractID interface {
	tz.Base58Encoder
	OriginatedContractID()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginatedContractID]{
		Variants: encoding.Variants[OriginatedContractID]{
			1: OriginatedContract{},
		},
	})
}

func ParseContractID(src []byte) (ContractID, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519PublicKeyHash:
		var out tz.Ed25519PublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.Secp256k1PublicKeyHash:
		var out tz.Secp256k1PublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.P256PublicKeyHash:
		var out tz.P256PublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.BLS12_381PublicKeyHash:
		var out tz.BLSPublicKeyHash
		copy(out[:], payload)
		return ImplicitContract{&out}, nil

	case &prefix.ContractHash:
		var out tz.ContractHash
		copy(out[:], payload)
		return OriginatedContract{&out, 0}, nil

	default:
		return nil, errors.New("gotez: unknown contract id prefix")
	}
}

type TransactionDestination interface {
	Address
	TransactionDestination()
}

type Address = tz.Base58Encoder

type Signed interface {
	GetSignature() (tz.Signature, error)
}

type OperationWithSource interface {
	GetSource() Address
}

type ManagerOperation interface {
	OperationContents
	OperationWithSource
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
