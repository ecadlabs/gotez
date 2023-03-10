package gotez

import (
	"errors"
	"math/big"

	"github.com/ecadlabs/goblst/minpk"
	"github.com/ecadlabs/gotez/base58"
	"github.com/ecadlabs/gotez/base58/prefix"
)

type Signature interface {
	Signature()
}

func (*GenericSignature) Signature()   {}
func (*Ed25519Signature) Signature()   {}
func (*Secp256k1Signature) Signature() {}
func (*P256Signature) Signature()      {}
func (*BLSSignature) Signature()       {}

var ErrSignatureType = errors.New("gotez: unknown signature type")

func NewSignatureFromBase58(src []byte) (Signature, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.GenericSignature:
		var out GenericSignature
		copy(out[:], payload)
		return &out, nil

	case &prefix.Ed25519Signature:
		var out Ed25519Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1Signature:
		var out Secp256k1Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256Signature:
		var out P256Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381Signature:
		var out BLSSignature
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, ErrSignatureType
	}
}

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

func NewBLSSignatureFrom(sig *minpk.Signature) *BLSSignature {
	var out BLSSignature
	s := sig.Bytes()
	if len(s) != len(out) {
		panic("gotez: invalid ed25519 signature length length")
	}
	copy(out[:], s)
	return &out
}
