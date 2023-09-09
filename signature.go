package gotez

import (
	"fmt"
	"math/big"
	"math/bits"
	"unsafe"
)

type Signature interface {
	Base58Encoder
	Signature()
}

type ConventionalSignature interface {
	Signature
	Generic() *GenericSignature
}

func (sig *GenericSignature) Signature()                 {}
func (sig *GenericSignature) Generic() *GenericSignature { return sig }

func (sig *Ed25519Signature) Signature()                 {}
func (sig *Ed25519Signature) Generic() *GenericSignature { return (*GenericSignature)(sig) }

func (sig *Secp256k1Signature) Signature()                 {}
func (sig *Secp256k1Signature) Generic() *GenericSignature { return (*GenericSignature)(sig) }

func (sig *P256Signature) Signature()                 {}
func (sig *P256Signature) Generic() *GenericSignature { return (*GenericSignature)(sig) }

func (sig *BLSSignature) Signature() {}

func NewEd25519Signature(sig []byte) *Ed25519Signature {
	var out Ed25519Signature
	if len(sig) != len(out) {
		panic("gotez: invalid ed25519 signature length length")
	}
	copy(out[:], sig)
	return &out
}

func bigByteLen(x *big.Int) int {
	return len(x.Bits()) * (bits.UintSize / 8)
}

func NewSecp256k1Signature(r, s *big.Int) *Secp256k1Signature {
	if bigByteLen(r) > 32 || bigByteLen(s) > 32 {
		panic("gotez: invalid ECDSA signature size")
	}
	var out Secp256k1Signature
	r.FillBytes(out[:32])
	s.FillBytes(out[32:])
	return &out
}

func (sig *Secp256k1Signature) Point() (r, s *big.Int) {
	return new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])
}

func NewP256Signature(r, s *big.Int) *P256Signature {
	if bigByteLen(r) > 32 || bigByteLen(s) > 32 {
		panic("gotez: invalid ECDSA signature size")
	}
	var out P256Signature
	r.FillBytes(out[:32])
	s.FillBytes(out[32:])
	return &out
}

func (sig *P256Signature) Point() (r, s *big.Int) {
	return new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])
}

func NewBLSSignature(compressedPoint []byte) *BLSSignature {
	var out BLSSignature
	if len(compressedPoint) != len(out) {
		panic("gotez: invalid ed25519 signature length length")
	}
	copy(out[:], compressedPoint)
	return &out
}

type AnySignature []byte

func (sig AnySignature) Signature() (Signature, error) {
	switch len(sig) {
	case GenericSignatureBytesLen:
		return (*GenericSignature)(unsafe.Pointer(&sig[0])), nil
	case BLSSignatureBytesLen:
		return (*BLSSignature)(unsafe.Pointer(&sig[0])), nil
	default:
		return nil, fmt.Errorf("gotez: invalid signature length: %d", len(sig))
	}
}

func (sig AnySignature) String() string {
	s, err := sig.Signature()
	if err != nil {
		return fmt.Sprintf("<%v>", err)
	}
	return string(s.ToBase58())
}

func (sig AnySignature) MarshalText() ([]byte, error) {
	s, err := sig.Signature()
	if err != nil {
		return nil, err
	}
	return s.ToBase58(), nil
}
