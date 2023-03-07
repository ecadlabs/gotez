package gotez

import (
	"errors"

	"github.com/ecadlabs/gotez/encoding"
)

const (
	Ed25519PublicKeyBytesLen             = 32
	Secp256k1PublicKeyBytesLen           = 33
	P256PublicKeyBytesLen                = 33
	BLSPublicKeyBytesLen                 = 48
	Ed25519PrivateKeyBytesLen            = 64
	Secp256k1PrivateKeyBytesLen          = 32
	P256PrivateKeyBytesLen               = 32
	BLSPrivateKeyBytesLen                = 32
	Ed25519EncryptedSeedBytesLen         = 56
	Secp256k1EncryptedPrivateKeyBytesLen = 56
	P256EncryptedPrivateKeyBytesLen      = 56
	BLSEncryptedPrivateKeyBytesLen       = 56
)

var (
	ErrPrivateKey          = errors.New("gotez: unknown private key type")
	ErrPublicKey           = errors.New("gotez: unknown public key type")
	ErrInvalidDecryptedLen = errors.New("gotez: invalid decrypted key length")
)

func (pkh *Ed25519PublicKeyHash) PublicKeyHash() []byte   { return pkh[:] }
func (pkh *Secp256k1PublicKeyHash) PublicKeyHash() []byte { return pkh[:] }
func (pkh *P256PublicKeyHash) PublicKeyHash() []byte      { return pkh[:] }
func (pkh *BLSPublicKeyHash) PublicKeyHash() []byte       { return pkh[:] }

func (*Ed25519PublicKey) PublicKey()   {}
func (*Secp256k1PublicKey) PublicKey() {}
func (*P256PublicKey) PublicKey()      {}
func (*BLSPublicKey) PublicKey()       {}

func (*Ed25519PrivateKey) PrivateKey()   {}
func (*Secp256k1PrivateKey) PrivateKey() {}
func (*P256PrivateKey) PrivateKey()      {}
func (*BLSPrivateKey) PrivateKey()       {}

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

type PublicKeyHash interface {
	Base58Encoder
	PublicKeyHash() []byte
}

type PublicKey interface {
	Base58Encoder
	PublicKey()
}

type PrivateKey interface {
	EncryptedPrivateKey
	PrivateKey()
}

type EncryptedPrivateKey interface {
	Base58Encoder
	Decrypt(passCb func() ([]byte, error)) (PrivateKey, error)
}

func NewPublicKeyFromBase58(src []byte) (PublicKey, error) {
	prefix, payload, err := DecodeTZBase58(src)
	if err != nil {
		return nil, err
	}
	switch prefix {
	case &PfxEd25519PublicKey:
		var out Ed25519PublicKey
		copy(out[:], payload)
		return &out, nil

	case &PfxSecp256k1PublicKey:
		var out Secp256k1PublicKey
		copy(out[:], payload)
		return &out, nil

	case &PfxP256PublicKey:
		var out P256PublicKey
		copy(out[:], payload)
		return &out, nil

	case &PfxBLS12_381PublicKey:
		var out BLSPublicKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, ErrPublicKey
	}
}

func NewPrivateKeyFromBase58(src []byte) (PrivateKey, error) {
	prefix, payload, err := DecodeTZBase58(src)
	if err != nil {
		return nil, err
	}
	switch prefix {
	case &PfxEd25519SecretKey:
		var out Ed25519PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxSecp256k1SecretKey:
		var out Secp256k1PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxP256SecretKey:
		var out P256PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxBLS12_381SecretKey:
		var out BLSPrivateKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, ErrPrivateKey
	}
}

func NewEncryptedPrivateKeyFromBase58(src []byte) (EncryptedPrivateKey, error) {
	prefix, payload, err := DecodeTZBase58(src)
	if err != nil {
		return nil, err
	}
	switch prefix {
	case &PfxEd25519SecretKey:
		var out Ed25519PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxSecp256k1SecretKey:
		var out Secp256k1PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxP256SecretKey:
		var out P256PrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxBLS12_381SecretKey:
		var out BLSPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxEd25519EncryptedSeed:
		var out Ed25519EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxSecp256k1EncryptedSecretKey:
		var out Secp256k1EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxP256EncryptedSecretKey:
		var out P256EncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	case &PfxBLS12_381EncryptedSecretKey:
		var out BLSEncryptedPrivateKey
		copy(out[:], payload)
		return &out, nil

	default:
		return nil, ErrPrivateKey
	}
}

func (pk *Ed25519EncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	if len(decrypted) != Ed25519PrivateKeyBytesLen {
		return nil, ErrInvalidDecryptedLen
	}
	var out Ed25519PrivateKey
	copy(out[:], decrypted)
	return &out, nil
}

func (pk *Secp256k1EncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	if len(decrypted) != Secp256k1PrivateKeyBytesLen {
		return nil, ErrInvalidDecryptedLen
	}
	var out Secp256k1PrivateKey
	copy(out[:], decrypted)
	return &out, nil
}

func (pk *P256EncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	if len(decrypted) != P256PrivateKeyBytesLen {
		return nil, ErrInvalidDecryptedLen
	}
	var out P256PrivateKey
	copy(out[:], decrypted)
	return &out, nil
}

func (pk *BLSEncryptedPrivateKey) Decrypt(passCb func() ([]byte, error)) (PrivateKey, error) {
	decrypted, err := decryptPrivateKey(pk[:], passCb)
	if err != nil {
		return nil, err
	}
	if len(decrypted) != BLSPrivateKeyBytesLen {
		return nil, ErrInvalidDecryptedLen
	}
	var out BLSPrivateKey
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
