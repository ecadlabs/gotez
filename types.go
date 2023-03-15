package gotez

//go:generate go run generate.go

import (
	"bytes"
	"errors"
	"time"

	"github.com/ecadlabs/gotez/encoding"
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
	BLSSignatureBytesLen          = 96
	CycleNonceBytesLen            = 32
	ProtocolHashBytesLen          = 32
	ContractHashBytesLen          = 20
	OperationHashBytesLen         = 32
	BlockPayloadHashBytesLen      = 32
	ScriptExprBytesLen            = 32
	SlotHeaderBytesLen            = 48
	ProofOfWorkNonceBytesLen      = 8
)

type Comparable[K any] interface {
	comparable
	ToKey() K
}

type ToComparable[H Comparable[K], K any] interface {
	ToComparable() H
}

type ContractID interface {
	Base58Encoder
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
	Base58Encoder
	OriginatedContractID()
}

func (*OriginatedContract) ContractID()           {}
func (*OriginatedContract) OriginatedContractID() {}
func (*ImplicitContract) ContractID()             {}

type Base58Encoder interface {
	ToBase58() []byte
	String() string
}

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

type String string

func (str *String) DecodeTZ(data []byte, ctx *encoding.Context) ([]byte, error) {
	if len(data) < 1 {
		return nil, encoding.ErrBuffer
	}
	length := int(data[0])
	if len(data) < 1+length {
		return nil, encoding.ErrBuffer
	}
	*str = String(data[1 : length+1])
	return data[length+1:], nil
}

func (str String) EncodeTZ(ctx *encoding.Context) ([]byte, error) {
	var buf bytes.Buffer
	if len(str) > 255 {
		return nil, errors.New("gotez: string is too long")
	}
	buf.WriteByte(byte(len(str)))
	buf.Write([]byte(str))
	return buf.Bytes(), nil
}

type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0).UTC()
}

func (t Timestamp) String() string {
	return t.Time().String()
}
