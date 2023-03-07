package protocol

import (
	tz "github.com/ecadlabs/gotez"
)

const (
	SeedNonceBytesLen             = 32
	SecretBytesLen                = 20
	PKHBytesLen                   = 20
	BlockHashBytesLen             = 32
	OperationListListHashBytesLen = 32
	ContextHashBytesLen           = 32
	ChainIdBytesLen               = 4
	GenericSignatureBytesLen      = 64
	CycleNonceBytesLen            = 32
	ProtocolHashBytesLen          = 32
	ContractHashBytesLen          = 20
	OperationHashBytesLen         = 32
	BlockPayloadHashBytesLen      = 32
	ScriptExprBytesLen            = 32
	Ed25519PublicKeyBytesLen      = 32
	Secp256K1PublicKeyBytesLen    = 33
	P256PublicKeyBytesLen         = 33
	SlotHeaderBytesLen            = 48
	ProofOfWorkNonceBytesLen      = 8
	BLSPublicKeyBytesLen          = 48
)

type BlockHash [BlockHashBytesLen]byte
type OperationsHash [OperationListListHashBytesLen]byte
type ContextHash [ContextHashBytesLen]byte
type BlockPayloadHash [BlockPayloadHashBytesLen]byte
type CycleNonceHash [CycleNonceBytesLen]byte
type Signature [GenericSignatureBytesLen]byte
type Ed25519PublicKeyHash [PKHBytesLen]byte
type Secp256k1PublicKeyHash [PKHBytesLen]byte
type P256PublicKeyHash [PKHBytesLen]byte
type BLSPublicKeyHash [PKHBytesLen]byte
type ProtocolHash [ProtocolHashBytesLen]byte
type ContractHash [ContractHashBytesLen]byte

type Ed25519PublicKey [Ed25519PublicKeyBytesLen]byte
type Secp256k1PublicKey [Secp256K1PublicKeyBytesLen]byte
type P256PublicKey [P256PublicKeyBytesLen]byte
type BLSPublicKey [BLSPublicKeyBytesLen]byte

func (*Ed25519PublicKeyHash) PublicKeyHash()   {}
func (*Secp256k1PublicKeyHash) PublicKeyHash() {}
func (*P256PublicKeyHash) PublicKeyHash()      {}
func (*BLSPublicKeyHash) PublicKeyHash()       {}

func (*Ed25519PublicKey) PublicKey()   {}
func (*Secp256k1PublicKey) PublicKey() {}
func (*P256PublicKey) PublicKey()      {}
func (*BLSPublicKey) PublicKey()       {}

type PublicKeyHash interface {
	PublicKeyHash()
}

type PublicKey interface {
	PublicKey()
}

type ContractID interface {
	ContractID()
}

type OriginatedContract struct {
	*ContractHash
	Padding uint8
}

type ImplicitContract struct {
	PublicKeyHash
}

type OriginatedContractID interface {
	OriginatedContractID()
}

func (*OriginatedContract) ContractID()           {}
func (*OriginatedContract) OriginatedContractID() {}
func (*ImplicitContract) ContractID()             {}

func init() {
	tz.RegisterEnum(&tz.Enum[PublicKeyHash]{
		Variants: tz.Variants[PublicKeyHash]{
			0: (*Ed25519PublicKeyHash)(nil),
			1: (*Secp256k1PublicKeyHash)(nil),
			2: (*P256PublicKeyHash)(nil),
			3: (*BLSPublicKeyHash)(nil),
		},
	})
	tz.RegisterEnum(&tz.Enum[PublicKey]{
		Variants: tz.Variants[PublicKey]{
			0: (*Ed25519PublicKey)(nil),
			1: (*Secp256k1PublicKey)(nil),
			2: (*P256PublicKey)(nil),
			3: (*BLSPublicKey)(nil),
		},
	})
	tz.RegisterEnum(&tz.Enum[ContractID]{
		Variants: tz.Variants[ContractID]{
			0: (*ImplicitContract)(nil),
			1: (*OriginatedContract)(nil),
		},
	})
	tz.RegisterEnum(&tz.Enum[OriginatedContractID]{
		Variants: tz.Variants[OriginatedContractID]{
			1: (*OriginatedContract)(nil),
		},
	})
}

type String string

func (str *String) DecodeTZ(data []byte, ctx *tz.Context) ([]byte, error) {
	if len(data) < 1 {
		return nil, tz.ErrBuffer
	}
	length := int(data[0])
	if len(data) < 1+length {
		return nil, tz.ErrBuffer
	}
	*str = String(data[1 : length+1])
	return data[length+1:], nil
}

type ShellHeader struct {
	Level          int32
	Proto          uint8
	Predecessor    *BlockHash
	Timestamp      tz.Timestamp
	ValidationPass uint8
	OperationsHash *OperationsHash
	Fitness        []byte `tz:"dyn"`
	Context        *ContextHash
}

type TenderbakeBlockHeader struct {
	ShellHeader
	PayloadHash               *BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*CycleNonceHash]
	LiquidityBakingToggleVote uint8
}
