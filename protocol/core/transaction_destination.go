package core

import (
	"bytes"
	"errors"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/b58/base58"
	"github.com/ecadlabs/gotez/v2/b58/prefix"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type TransactionDestination interface {
	tz.Base58Encoder
	TransactionDestination()
	Eq(other TransactionDestination) bool
}

type Address = TransactionDestination

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: ImplicitContract{},
			1: OriginatedContract{},
			2: TxRollupDestination{},
			3: SmartRollupDestination{},
			4: ZkRollupDestination{},
		},
	})
}

type TxRollupDestination struct {
	*tz.TXRollupAddress
	Padding uint8
}

func (TxRollupDestination) TransactionDestination() {}
func (a TxRollupDestination) Eq(b TransactionDestination) bool {
	if other, ok := b.(TxRollupDestination); ok {
		return bytes.Equal(a.TXRollupAddress[:], other.TXRollupAddress[:])
	}
	return false
}

type SmartRollupDestination struct {
	*tz.SmartRollupAddress
	Padding uint8
}

func (SmartRollupDestination) TransactionDestination() {}
func (a SmartRollupDestination) Eq(b TransactionDestination) bool {
	if other, ok := b.(SmartRollupDestination); ok {
		return bytes.Equal(a.SmartRollupAddress[:], other.SmartRollupAddress[:])
	}
	return false
}

type ZkRollupDestination struct {
	*tz.ZkRollupAddress
	Padding uint8
}

func (ZkRollupDestination) TransactionDestination() {}
func (a ZkRollupDestination) Eq(b TransactionDestination) bool {
	if other, ok := b.(ZkRollupDestination); ok {
		return bytes.Equal(a.ZkRollupAddress[:], other.ZkRollupAddress[:])
	}
	return false
}

func ParseTransactionDestination(src []byte) (TransactionDestination, error) {
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

	case &prefix.TXRollupAddress:
		var out tz.TXRollupAddress
		copy(out[:], payload)
		return TxRollupDestination{&out, 0}, nil

	case &prefix.SmartRollupHash:
		var out tz.SmartRollupAddress
		copy(out[:], payload)
		return SmartRollupDestination{&out, 0}, nil

	case &prefix.ZkRollupHash:
		var out tz.ZkRollupAddress
		copy(out[:], payload)
		return ZkRollupDestination{&out, 0}, nil

	default:
		return nil, errors.New("gotez: unknown destination prefix")
	}
}
