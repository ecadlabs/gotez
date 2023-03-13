package gotez

import (
	"errors"
	"math/big"

	"github.com/ecadlabs/goblst/minpk"
)

type Signature interface {
	Base58Encoder
	Signature()
}

func (*GenericSignature) Signature()   {}
func (*Ed25519Signature) Signature()   {}
func (*Secp256k1Signature) Signature() {}
func (*P256Signature) Signature()      {}
func (*BLSSignature) Signature()       {}

var ErrSignatureType = errors.New("gotez: unknown signature type")

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

func (sig *Secp256k1Signature) Get() (r, s *big.Int) {
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

func (sig *P256Signature) Get() (r, s *big.Int) {
	return new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])
}

func NewBLSSignature(sig *minpk.Signature) *BLSSignature {
	var out BLSSignature
	s := sig.Bytes()
	if len(s) != len(out) {
		panic("gotez: invalid ed25519 signature length length")
	}
	copy(out[:], s)
	return &out
}

func (sig *BLSSignature) Get() (*minpk.Signature, error) {
	return minpk.SignatureFromBytes(sig[:])
}
