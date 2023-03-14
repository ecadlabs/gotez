// Package b58 contains Base58 decoding functions of various Tezos types
package b58

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/b58/base58"
	"github.com/ecadlabs/gotez/b58/prefix"
)

//go:generate go run generate.go

func ParsePublicKey(src []byte) (tz.PublicKey, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519PublicKey:
		var out tz.Ed25519PublicKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1PublicKey:
		var out tz.Secp256k1PublicKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256PublicKey:
		var out tz.P256PublicKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381PublicKey:
		var out tz.BLSPublicKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, tz.ErrPublicKeyType
	}
}

func ParsePublicKeyHash(src []byte) (tz.PublicKeyHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519PublicKeyHash:
		var out tz.Ed25519PublicKeyHash
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1PublicKeyHash:
		var out tz.Secp256k1PublicKeyHash
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256PublicKeyHash:
		var out tz.P256PublicKeyHash
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381PublicKeyHash:
		var out tz.BLSPublicKeyHash
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, tz.ErrPublicKeyType
	}
}

func ParsePrivateKey(src []byte) (tz.PrivateKey, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519Seed:
		var out tz.Ed25519PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1SecretKey:
		var out tz.Secp256k1PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256SecretKey:
		var out tz.P256PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381SecretKey:
		var out tz.BLSPrivateKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, tz.ErrPrivateKeyType
	}
}

func ParseEncryptedPrivateKey(src []byte) (tz.EncryptedPrivateKey, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.Ed25519Seed:
		var out tz.Ed25519PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1SecretKey:
		var out tz.Secp256k1PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256SecretKey:
		var out tz.P256PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381SecretKey:
		var out tz.BLSPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Ed25519EncryptedSeed:
		var out tz.Ed25519EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1EncryptedSecretKey:
		var out tz.Secp256k1EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256EncryptedSecretKey:
		var out tz.P256EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381EncryptedSecretKey:
		var out tz.BLSEncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, tz.ErrPrivateKeyType
	}
}

func ParseSignature(src []byte) (tz.Signature, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	switch pre {
	case &prefix.GenericSignature:
		var out tz.GenericSignature
		copy(out[:], payload)
		return &out, nil

	case &prefix.Ed25519Signature:
		var out tz.Ed25519Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.Secp256k1Signature:
		var out tz.Secp256k1Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.P256Signature:
		var out tz.P256Signature
		copy(out[:], payload)
		return &out, nil

	case &prefix.BLS12_381Signature:
		var out tz.BLSSignature
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, tz.ErrSignatureType
	}
}
