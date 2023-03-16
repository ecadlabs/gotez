package gotez

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ecadlabs/gotez/b58/base58"
	"github.com/ecadlabs/gotez/b58/prefix"
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
	ErrInvalidDecryptedLen = errors.New("gotez: invalid decrypted key length")
)

type PublicKeyHash interface {
	Base58Encoder
	PublicKeyHash() []byte
	ToComparable[EncodedPublicKeyHash, PublicKeyHash]
}

type PublicKey interface {
	Base58Encoder
	PublicKey()
	Hash() PublicKeyHash
}

type PrivateKey interface {
	PrivateKey()
	EncryptedPrivateKey
}

type EncryptedPrivateKey interface {
	Base58Encoder
	Decrypt(passCb func() ([]byte, error)) (PrivateKey, error)
}

func (pkh *Ed25519PublicKeyHash) PublicKeyHash() []byte   { return pkh[:] }
func (pkh *Secp256k1PublicKeyHash) PublicKeyHash() []byte { return pkh[:] }
func (pkh *P256PublicKeyHash) PublicKeyHash() []byte      { return pkh[:] }
func (pkh *BLSPublicKeyHash) PublicKeyHash() []byte       { return pkh[:] }

func (priv *Ed25519PrivateKey) PrivateKey()   {}
func (priv *Secp256k1PrivateKey) PrivateKey() {}
func (priv *P256PrivateKey) PrivateKey()      {}
func (priv *BLSPrivateKey) PrivateKey()       {}

func (pk *Ed25519PublicKey) PublicKey()   {}
func (pk *Secp256k1PublicKey) PublicKey() {}
func (pk *P256PublicKey) PublicKey()      {}
func (pk *BLSPublicKey) PublicKey()       {}

const publicKeyHashComparableKeyLen = PKHBytesLen + 1

type EncodedPublicKeyHash [publicKeyHashComparableKeyLen]byte

func (k EncodedPublicKeyHash) ToKey() PublicKeyHash {
	var pkh PublicKeyHash
	if _, err := encoding.Decode(k[:], &pkh); err != nil {
		panic(err)
	}
	return pkh
}

func (k EncodedPublicKeyHash) MarshalText() (text []byte, err error) {
	var pkh PublicKeyHash
	if _, err := encoding.Decode(k[:], &pkh); err != nil {
		return nil, err
	}
	return pkh.ToBase58(), nil
}

func (k *EncodedPublicKeyHash) UnmarshalText(text []byte) error {
	pre, payload, err := base58.DecodeTZ(text)
	if err != nil {
		return err
	}
	var result PublicKeyHash
	switch pre {
	case &prefix.Ed25519PublicKeyHash:
		var out Ed25519PublicKeyHash
		copy(out[:], payload)
		result = &out

	case &prefix.Secp256k1PublicKeyHash:
		var out Secp256k1PublicKeyHash
		copy(out[:], payload)
		result = &out

	case &prefix.P256PublicKeyHash:
		var out P256PublicKeyHash
		copy(out[:], payload)
		result = &out

	case &prefix.BLS12_381PublicKeyHash:
		var out BLSPublicKeyHash
		copy(out[:], payload)
		result = &out

	default:
		return errors.New("gotez: unknown public key prefix")
	}
	var (
		x   PublicKeyHash = result
		buf bytes.Buffer
	)
	if err := encoding.Encode(&buf, &x); err != nil {
		panic(err)
	}
	out := buf.Bytes()
	if len(out) != publicKeyHashComparableKeyLen {
		panic("invalid public key hash length")
	}
	copy(k[:], out)
	return nil
}

func (pkh *Ed25519PublicKeyHash) ToComparable() (out EncodedPublicKeyHash) {
	var (
		x   PublicKeyHash = pkh
		buf bytes.Buffer
	)
	if err := encoding.Encode(&buf, &x); err != nil {
		panic(err)
	}
	b := buf.Bytes()
	if len(b) != publicKeyHashComparableKeyLen {
		panic("invalid public key hash length")
	}
	copy(out[:], b)
	return
}

func (pkh *Secp256k1PublicKeyHash) ToComparable() (out EncodedPublicKeyHash) {
	var (
		x   PublicKeyHash = pkh
		buf bytes.Buffer
	)
	if err := encoding.Encode(&buf, &x); err != nil {
		panic(err)
	}
	b := buf.Bytes()
	if len(b) != publicKeyHashComparableKeyLen {
		panic("invalid public key hash length")
	}
	copy(out[:], b)
	return
}

func (pkh *P256PublicKeyHash) ToComparable() (out EncodedPublicKeyHash) {
	var (
		x   PublicKeyHash = pkh
		buf bytes.Buffer
	)
	if err := encoding.Encode(&buf, &x); err != nil {
		panic(err)
	}
	b := buf.Bytes()
	if len(b) != publicKeyHashComparableKeyLen {
		panic("invalid public key hash length")
	}
	copy(out[:], b)
	return
}

func (pkh *BLSPublicKeyHash) ToComparable() (out EncodedPublicKeyHash) {
	var (
		x   PublicKeyHash = pkh
		buf bytes.Buffer
	)
	if err := encoding.Encode(&buf, &x); err != nil {
		panic(err)
	}
	b := buf.Bytes()
	if len(b) != publicKeyHashComparableKeyLen {
		panic("invalid public key hash length")
	}
	copy(out[:], b)
	return
}

func NewEd25519PublicKey(key []byte) *Ed25519PublicKey {
	var out Ed25519PublicKey
	if len(key) != len(out) {
		panic("gotez: invalid Ed25519 public key length")
	}
	copy(out[:], key)
	return &out
}

func NewSecp256k1PublicKey(compressedPoint []byte) *Secp256k1PublicKey {
	var out Secp256k1PublicKey
	if len(compressedPoint) != len(out) {
		panic("gotez: invalid Secp256k1 public key length")
	}
	copy(out[:], compressedPoint)
	return &out
}

func NewP256PublicKey(compressedPoint []byte) *P256PublicKey {
	var out P256PublicKey
	if len(compressedPoint) != len(out) {
		panic("gotez: invalid P256 public key length")
	}
	copy(out[:], compressedPoint)
	return &out
}

func NewBLSPublicKey(compressedPoint []byte) *BLSPublicKey {
	var out BLSPublicKey
	if len(compressedPoint) != len(out) {
		panic("gotez: invalid BLS public key length")
	}
	copy(out[:], compressedPoint)
	return &out
}

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

func NewEd25519PrivateKey(key []byte) *Ed25519PrivateKey {
	var out Ed25519PrivateKey
	if len(key) != len(out) {
		panic("gotez: invalid Ed25519 private key length")
	}
	copy(out[:], key)
	return &out
}

func NewSecp256k1PrivateKey(scalar *big.Int) *Secp256k1PrivateKey {
	var out Secp256k1PrivateKey
	payload := scalar.Bytes()
	if len(payload) > len(out) {
		panic("gotez: invalid Secp256k1 private key length")
	}
	copy(out[len(out)-len(payload):], payload)
	return &out
}

func NewP256PrivateKey(scalar *big.Int) *P256PrivateKey {
	var out P256PrivateKey
	payload := scalar.Bytes()
	if len(payload) > len(out) {
		panic("gotez: invalid P256 private key length")
	}
	copy(out[len(out)-len(payload):], payload)
	return &out
}

func NewBLSPrivateKey(scalar []byte) *BLSPrivateKey {
	var out BLSPrivateKey
	if len(scalar) != len(out) {
		panic("gotez: invalid BLS private key length")
	}
	copy(out[:], scalar)
	return &out
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
