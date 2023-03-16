package gotez

import (
	"math/big"
)

type Signature interface {
	Base58Encoder
	Signature()
}

func (sig *GenericSignature) Signature()   {}
func (sig *Ed25519Signature) Signature()   {}
func (sig *Secp256k1Signature) Signature() {}
func (sig *P256Signature) Signature()      {}
func (sig *BLSSignature) Signature()       {}

func NewEd25519Signature(sig []byte) *Ed25519Signature {
	var out Ed25519Signature
	if len(sig) != len(out) {
		panic("gotez: invalid ed25519 signature length length")
	}
	copy(out[:], sig)
	return &out
}

func NewSecp256k1Signature(r, s *big.Int) *Secp256k1Signature {
	sr := r.Bytes()
	ss := s.Bytes()
	if len(sr) > 32 || len(ss) > 32 {
		panic("gotez: invalid ECDSA signature size")
	}
	var out Secp256k1Signature
	copy(out[32-len(sr):], sr)
	copy(out[64-len(ss):], ss)
	return &out
}

func (sig *Secp256k1Signature) Point() (r, s *big.Int) {
	return new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])
}

func NewP256Signature(r, s *big.Int) *P256Signature {
	sr := r.Bytes()
	ss := s.Bytes()
	if len(sr) > 32 || len(ss) > 32 {
		panic("gotez: invalid ECDSA signature size")
	}
	var out P256Signature
	copy(out[32-len(sr):], sr)
	copy(out[64-len(ss):], ss)
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
