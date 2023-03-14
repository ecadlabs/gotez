package gotez

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"errors"
	"fmt"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	"github.com/ecadlabs/gotez/signature"
)

type Signature interface {
	Base58Encoder
	Signature(pub crypto.PublicKey) (signature.Signature, error)
}

func (sig *GenericSignature) Signature(pub crypto.PublicKey) (signature.Signature, error) {
	switch pk := pub.(type) {
	case *ecdsa.PublicKey:
		r, s := (*P256Signature)(sig).Get()
		return &signature.ECDSA{
			R:     r,
			S:     s,
			Curve: pk.Curve,
		}, nil

	case ed25519.PublicKey:
		return signature.ED25519(sig[:]), nil

	default:
		return nil, errors.New("gotez: generic signature encoding can't be used with BLS")
	}
}

func (sig *Ed25519Signature) Signature(crypto.PublicKey) (signature.Signature, error) {
	return signature.ED25519(sig[:]), nil
}

func (sig *Secp256k1Signature) Signature(crypto.PublicKey) (signature.Signature, error) {
	r, s := new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])
	return &signature.ECDSA{
		R:     r,
		S:     s,
		Curve: secp256k1.S256(),
	}, nil
}

func (sig *P256Signature) Signature(crypto.PublicKey) (signature.Signature, error) {
	r, s := new(big.Int).SetBytes(sig[:32]), new(big.Int).SetBytes(sig[32:])
	return &signature.ECDSA{
		R:     r,
		S:     s,
		Curve: elliptic.P256(),
	}, nil
}

func (sig *BLSSignature) Signature(crypto.PublicKey) (signature.Signature, error) {
	return minpk.SignatureFromBytes(sig[:])
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

func NewSignature(sig signature.Signature) (Signature, error) {
	switch s := sig.(type) {
	case signature.ED25519:
		return NewEd25519Signature(s), nil

	case *signature.ECDSA:
		switch {
		case s.Curve == elliptic.P256():
			return NewP256Signature(s.R, s.S), nil
		case curveEqual(s.Curve, secp256k1.S256()):
			return NewSecp256k1Signature(s.R, s.S), nil
		default:
			return nil, fmt.Errorf("gotez: unknown curve `%s`", s.Curve.Params().Name)

		}

	case *minpk.Signature:
		return NewBLSSignature(s), nil

	default:
		return nil, fmt.Errorf("gotez: unknown signature type %T", sig)
	}
}
