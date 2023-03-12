package gotez

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"errors"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	"github.com/ecadlabs/gotez/encoding"
	"golang.org/x/crypto/blake2b"
)

const (
	Ed25519PublicKeyBytesLen             = 32
	Secp256k1PublicKeyBytesLen           = 33
	P256PublicKeyBytesLen                = 33
	BLSPublicKeyBytesLen                 = 48
	Ed25519SeedBytesLen                  = 32
	Secp256k1PrivateKeyBytesLen          = 32
	P256PrivateKeyBytesLen               = 32
	BLSPrivateKeyBytesLen                = 32
	Ed25519EncryptedSeedBytesLen         = 56
	Secp256k1EncryptedPrivateKeyBytesLen = 56
	P256EncryptedPrivateKeyBytesLen      = 56
	BLSEncryptedPrivateKeyBytesLen       = 56
)

var (
	ErrPrivateKeyType      = errors.New("gotez: unknown private key type")
	ErrPublicKeyType       = errors.New("gotez: unknown public key type")
	ErrInvalidDecryptedLen = errors.New("gotez: invalid decrypted key length")
)

type PublicKeyHash interface {
	Base58Encoder
	PublicKeyHash()
}

type PublicKey interface {
	Base58Encoder
	PublicKey() (crypto.PublicKey, error)
	Hash() PublicKeyHash
}

type PrivateKey interface {
	EncryptedPrivateKey
	PrivateKey() (crypto.PrivateKey, error)
}

type EncryptedPrivateKey interface {
	Base58Encoder
	Decrypt(passCb func() ([]byte, error)) (PrivateKey, error)
}

func (pkh *Ed25519PublicKeyHash) PublicKeyHash()   {}
func (pkh *Secp256k1PublicKeyHash) PublicKeyHash() {}
func (pkh *P256PublicKeyHash) PublicKeyHash()      {}
func (pkh *BLSPublicKeyHash) PublicKeyHash()       {}

func (pk *Ed25519PublicKey) Hash() PublicKeyHash {
	digest, err := blake2b.New(20, nil)
	if err != nil {
		panic(err)
	}
	digest.Write(pk[:])
	var out Ed25519PublicKeyHash
	copy(out[:], digest.Sum(nil))
	return &out
}

func (pk *Secp256k1PublicKey) Hash() PublicKeyHash {
	digest, err := blake2b.New(20, nil)
	if err != nil {
		panic(err)
	}
	digest.Write(pk[:])
	var out Secp256k1PublicKeyHash
	copy(out[:], digest.Sum(nil))
	return &out
}

func (pk *P256PublicKey) Hash() PublicKeyHash {
	digest, err := blake2b.New(20, nil)
	if err != nil {
		panic(err)
	}
	digest.Write(pk[:])
	var out P256PublicKeyHash
	copy(out[:], digest.Sum(nil))
	return &out
}

func (pk *BLSPublicKey) Hash() PublicKeyHash {
	digest, err := blake2b.New(20, nil)
	if err != nil {
		panic(err)
	}
	digest.Write(pk[:])
	var out BLSPublicKeyHash
	copy(out[:], digest.Sum(nil))
	return &out
}

func (pk *Ed25519PublicKey) PublicKey() (crypto.PublicKey, error) {
	if len(pk) != ed25519.PublicKeySize {
		panic("gotez: invalid ed25519 public key length") // unlikely
	}
	out := make(ed25519.PublicKey, ed25519.PublicKeySize)
	copy(out, pk[:])
	return out, nil
}

func (pk *Secp256k1PublicKey) PublicKey() (crypto.PublicKey, error) {
	x, y, err := unmarshalCompressed(pk[:], secp256k1.S256())
	if err != nil {
		return nil, err
	}
	return &ecdsa.PublicKey{
		Curve: secp256k1.S256(),
		X:     x,
		Y:     y,
	}, nil
}

func (pk *P256PublicKey) PublicKey() (crypto.PublicKey, error) {
	x, y, err := unmarshalCompressed(pk[:], elliptic.P256())
	if err != nil {
		return nil, err
	}
	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, nil
}

func (pk *BLSPublicKey) PublicKey() (crypto.PublicKey, error) {
	return minpk.PublicKeyFromBytes(pk[:])
}

func (priv *Ed25519PrivateKey) PrivateKey() (crypto.PrivateKey, error) {
	if len(priv) != ed25519.SeedSize {
		panic("gotez: invalid ed25519 private key length") // unlikely
	}
	return ed25519.NewKeyFromSeed(priv[:]), nil
}

func (priv *Secp256k1PrivateKey) PrivateKey() (crypto.PrivateKey, error) {
	return ecPrivateKeyFromBytes(priv[:], secp256k1.S256())

}

func (priv *P256PrivateKey) PrivateKey() (crypto.PrivateKey, error) {
	return ecPrivateKeyFromBytes(priv[:], elliptic.P256())
}

func (priv *BLSPrivateKey) PrivateKey() (crypto.PrivateKey, error) {
	return minpk.PrivateKeyFromBytes(priv[:])
}

// stub
func (pk *Ed25519PrivateKey) Decrypt(func() ([]byte, error)) (PrivateKey, error) {
	return pk, nil
}

// stub
func (pk *Secp256k1PrivateKey) Decrypt(func() ([]byte, error)) (PrivateKey, error) {
	return pk, nil
}

// stub
func (pk *P256PrivateKey) Decrypt(func() ([]byte, error)) (PrivateKey, error) { return pk, nil }

// stub
func (pk *BLSPrivateKey) Decrypt(func() ([]byte, error)) (PrivateKey, error) { return pk, nil }

func NewPublicKey(src crypto.PublicKey) (PublicKey, error) {
	switch key := src.(type) {
	case *ecdsa.PublicKey:
		payload := elliptic.MarshalCompressed(key.Curve, key.X, key.Y)
		switch {
		case key.Curve == elliptic.P256():
			var out P256PublicKey
			if len(payload) != len(out) {
				panic("gotez: invalid public key length")
			}
			copy(out[:], payload)
			return &out, nil
		case curveEqual(key.Curve, secp256k1.S256()):
			var out Secp256k1PublicKey
			if len(payload) != len(out) {
				panic("gotez: invalid public key length")
			}
			copy(out[:], payload)
			return &out, nil
		default:
			return nil, ErrPublicKeyType
		}

	case ed25519.PublicKey:
		var out Ed25519PublicKey
		if len(key) != len(out) {
			panic("gotez: invalid public key length")
		}
		copy(out[:], key)
		return &out, nil

	case *minpk.PublicKey:
		payload := key.Bytes()
		var out BLSPublicKey
		if len(payload) != len(out) {
			panic("gotez: invalid public key length")
		}
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, ErrPublicKeyType
	}
}

func NewPrivateKey(src crypto.PrivateKey) (PrivateKey, error) {
	switch key := src.(type) {
	case *ecdsa.PrivateKey:
		b := key.D.Bytes()
		payload := make([]byte, (key.Params().N.BitLen()+7)>>3)
		copy(payload[len(payload)-len(b):], b)
		switch {
		case key.Curve == elliptic.P256():
			var out P256PrivateKey
			if len(out) != len(payload) {
				panic("gotez: invalid private key length")
			}
			copy(out[:], payload)
			return &out, nil
		case curveEqual(key.Curve, secp256k1.S256()):
			var out Secp256k1PrivateKey
			if len(out) != len(payload) {
				panic("gotez: invalid private key length")
			}
			copy(out[:], payload)
			return &out, nil
		default:
			return nil, ErrPrivateKeyType
		}

	case ed25519.PrivateKey:
		payload := key.Seed()
		var out Ed25519PrivateKey
		if len(out) != len(payload) {
			panic("gotez: invalid private key length")
		}
		copy(out[:], payload)
		return &out, nil

	case *minpk.PrivateKey:
		payload := key.Bytes()
		var out BLSPrivateKey
		if len(out) != len(payload) {
			panic("gotez: invalid private key length")
		}
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, ErrPrivateKeyType
	}
}

func (pk *Ed25519EncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	var out Ed25519PrivateKey
	if len(decrypted) != len(out) {
		return nil, ErrInvalidDecryptedLen
	}
	copy(out[:], decrypted)
	return &out, nil
}

func (pk *Secp256k1EncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	var out Secp256k1PrivateKey
	if len(decrypted) != len(out) {
		return nil, ErrInvalidDecryptedLen
	}
	copy(out[:], decrypted)
	return &out, nil
}

func (pk *P256EncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	var out P256PrivateKey
	if len(decrypted) != len(out) {
		return nil, ErrInvalidDecryptedLen
	}
	copy(out[:], decrypted)
	return &out, nil
}

func (pk *BLSEncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	var out BLSPrivateKey
	if len(decrypted) != len(out) {
		return nil, ErrInvalidDecryptedLen
	}
	copy(out[:], decrypted)
	return &out, nil
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
}
