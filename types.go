package gotez

//go:generate go run generate.go

import (
	"fmt"
	"math/big"
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

func (*Ed25519PublicKeyHash) PublicKeyHash()   {}
func (*Ed25519PublicKey) PublicKey()           {}
func (*Secp256k1PublicKeyHash) PublicKeyHash() {}
func (*Secp256k1PublicKey) PublicKey()         {}
func (*P256PublicKeyHash) PublicKeyHash()      {}
func (*P256PublicKey) PublicKey()              {}
func (*BLSPublicKeyHash) PublicKeyHash()       {}
func (*BLSPublicKey) PublicKey()               {}

type PublicKeyHash interface {
	Base58Encoder
	PublicKeyHash()
}

type PublicKey interface {
	Base58Encoder
	PublicKey()
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
	Base58() []byte
	String() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[PublicKeyHash]{
		Variants: encoding.Variants[PublicKeyHash]{
			0: (*Ed25519PublicKeyHash)(nil),
			1: (*Secp256k1PublicKeyHash)(nil),
			2: (*P256PublicKeyHash)(nil),
			3: (*BLSPublicKeyHash)(nil),
		},
	})
	encoding.RegisterEnum(&encoding.Enum[PublicKey]{
		Variants: encoding.Variants[PublicKey]{
			0: (*Ed25519PublicKey)(nil),
			1: (*Secp256k1PublicKey)(nil),
			2: (*P256PublicKey)(nil),
			3: (*BLSPublicKey)(nil),
		},
	})
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

type BigInt []byte

func getLen(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, encoding.ErrBuffer
	}
	i := 0
	for i < len(data) && data[i]&0x80 != 0 {
		i += 1
	}
	if i == len(data) {
		return 0, encoding.ErrBuffer
	}
	return i + 1, nil
}

func (b *BigInt) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	ln, err := getLen(data)
	if err != nil {
		return nil, err
	}
	*b = BigInt(data[:ln])
	return data[ln:], nil
}

func (b BigInt) Int() *big.Int {
	res := big.NewInt(0)
	if len(b) == 0 {
		return res
	}
	neg := b[0]&0x40 != 0
	shift := uint(0)
	var tmp big.Int
	for i, x := range b {
		var (
			mask uint8
			s    uint
		)
		if i == 0 {
			mask = 0x3f
			s = 6
		} else {
			mask = 0x7f
			s = 7
		}
		tmp.SetInt64(int64(x & mask))
		tmp.Lsh(&tmp, shift)
		res.Or(res, &tmp)
		shift += s
	}
	if neg {
		res.Neg(res)
	}
	return res
}

func (b BigInt) String() string {
	return b.Int().String()
}

type BigUint []byte

func (b *BigUint) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	ln, err := getLen(data)
	if err != nil {
		return nil, err
	}
	*b = BigUint(data[:ln])
	return data[ln:], nil
}

func (b BigUint) Int() *big.Int {
	res := big.NewInt(0)
	if len(b) == 0 {
		return res
	}
	shift := uint(0)
	var tmp big.Int
	for _, x := range b {
		tmp.SetInt64(int64(x & 0x7f))
		tmp.Lsh(&tmp, shift)
		res.Or(res, &tmp)
		shift += 7
	}
	return res
}

func (b BigUint) String() string {
	return b.Int().String()
}

type Timestamp int64

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0).UTC()
}

func (t Timestamp) String() string {
	return t.Time().String()
}

type Option[T any] struct {
	some  bool
	value T
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		some:  true,
		value: val,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

func (op Option[T]) Unwrap() T {
	if !op.some {
		panic(fmt.Sprintf("called `Unwrap()` on a `None` value of type %T", op))
	}
	return op.value
}

func (op Option[T]) UnwrapUnchecked() T {
	return op.value
}

func (op Option[T]) IsSome() bool { return op.some }
func (op Option[T]) IsNone() bool { return !op.some }

func (op Option[T]) Or(val Option[T]) Option[T] {
	if op.some {
		return op
	}
	return val
}

func (op Option[T]) OrElse(f func() Option[T]) Option[T] {
	if op.some {
		return op
	}
	return f()
}

func (op Option[T]) UnwrapOr(def T) T {
	if op.some {
		return op.value
	}
	return def
}

func (op Option[T]) UnwrapOrElse(f func() T) T {
	if op.some {
		return op.value
	}
	return f()
}

func (op Option[T]) UnwrapOrZero() T {
	if op.some {
		return op.value
	}
	var t T
	return t
}

func (op *Option[T]) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if len(data) < 1 {
		return nil, encoding.ErrBuffer
	}
	out := Option[T]{
		some: data[0] != 0,
	}
	data = data[1:]

	if out.some {
		data, err = encoding.Decode(data, &out.value, encoding.Ctx(ctx))
		if err != nil {
			return nil, err
		}
	}
	*op = out
	return data, nil
}

func (op Option[T]) String() string {
	if op.some {
		return fmt.Sprintf("Some(%v)", op.value)
	}
	return "None"
}
