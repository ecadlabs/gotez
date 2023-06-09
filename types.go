package gotez

//go:generate go run generate.go

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/ecadlabs/gotez/v2/encoding"
)

const (
	ProofOfWorkNonceBytesLen = 8
	ChainIdBytesLen          = 4
	SecretBytesLen           = 20
	AddressBytesLen          = 20
	ContractHashBytesLen     = 20
	SeedNonceBytesLen        = 32
	CycleNonceBytesLen       = 32
	HashBytesLen             = 32
	SlotHeaderBytesLen       = 48
	GenericSignatureBytesLen = 64
	BLSSignatureBytesLen     = 96
)

type Base58Encoder interface {
	ToBase58() []byte
	String() string
}

type Comparable[K any] interface {
	comparable
	ToKey() K
}

type ToComparable[H Comparable[K], K any] interface {
	ToComparable() H
}

type String string

func (str *String) DecodeTZ(data []byte, ctx *encoding.Context) ([]byte, error) {
	if len(data) < 1 {
		return nil, fmt.Errorf("(string) %w", encoding.ErrBuffer(1))
	}
	length := int(data[0])
	if len(data) < 1+length {
		return nil, fmt.Errorf("(string) %w", encoding.ErrBuffer(1))
	}
	*str = String(data[1 : length+1])
	return data[length+1:], nil
}

func (str String) EncodeTZ(ctx *encoding.Context) ([]byte, error) {
	var buf bytes.Buffer
	if len(str) > 255 {
		return nil, errors.New("string is too long")
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

type Bytes []byte

func (b Bytes) MarshalText() (text []byte, err error) {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return dst, nil
}

func (b *Bytes) UnmarshalText(text []byte) error {
	*b = make(Bytes, hex.DecodedLen(len(text)))
	_, err := hex.Decode(*b, text)
	return err
}
