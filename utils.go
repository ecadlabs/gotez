package gotez

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha512"
	"errors"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/pbkdf2"
)

var (
	ErrPrivateKeyDecrypt = errors.New("gotez: unable to decrypt the private key")
	ErrInvalidPublicKey  = errors.New("gotez: invalid public key")
	ErrInvalidPrivateKey = errors.New("gotez: invalid private key")
)

const (
	encIterations = 32768
	encKeyLen     = 32
)

func decryptPrivateKey(data []byte, passCb func() ([]byte, error)) ([]byte, error) {
	if passCb == nil {
		return nil, ErrPrivateKeyDecrypt
	}
	passphrase, err := passCb()
	if err != nil {
		return nil, err
	}

	salt, box := data[:8], data[8:]
	secretboxKey := pbkdf2.Key(passphrase, salt, encIterations, encKeyLen, sha512.New)

	var (
		tmp   [32]byte
		nonce [24]byte
	)
	copy(tmp[:], secretboxKey)
	opened, ok := secretbox.Open(nil, box, &nonce, &tmp)
	if !ok {
		return nil, ErrPrivateKeyDecrypt
	}

	return opened, nil
}

func curveEqual(a, b elliptic.Curve) bool {
	return a == b ||
		a.Params().BitSize == b.Params().BitSize &&
			a.Params().P.Cmp(b.Params().P) == 0 &&
			a.Params().N.Cmp(b.Params().N) == 0 &&
			a.Params().B.Cmp(b.Params().B) == 0 &&
			a.Params().Gx.Cmp(b.Params().Gx) == 0 &&
			a.Params().Gy.Cmp(b.Params().Gy) == 0
}

// See https://github.com/golang/go/blob/master/src/crypto/elliptic/elliptic.go
func unmarshalCompressed(data []byte, curve elliptic.Curve) (x, y *big.Int, err error) {
	byteLen := (curve.Params().BitSize + 7) / 8
	if len(data) != 1+byteLen {
		return nil, nil, ErrInvalidPublicKey
	}
	if data[0] != 2 && data[0] != 3 { // compressed form
		return nil, nil, ErrInvalidPublicKey
	}
	p := curve.Params().P
	x = new(big.Int).SetBytes(data[1:])
	if x.Cmp(p) >= 0 {
		return nil, nil, ErrInvalidPublicKey
	}

	// secp256k1 polynomial: x³ + b
	// P-* polynomial: x³ - 3x + b
	y = new(big.Int).Mul(x, x)
	y.Mul(y, x)
	if curve != secp256k1.S256() {
		x1 := new(big.Int).Lsh(x, 1)
		x1.Add(x1, x)
		y.Sub(y, x1)
	}
	y.Add(y, curve.Params().B)
	y.Mod(y, curve.Params().P)
	y.ModSqrt(y, p)

	if y == nil {
		return nil, nil, ErrInvalidPublicKey
	}
	if byte(y.Bit(0)) != data[0]&1 {
		y.Neg(y).Mod(y, p)
	}
	if !curve.IsOnCurve(x, y) {
		return nil, nil, ErrInvalidPublicKey
	}
	return
}

// see https://golang.org/src/crypto/x509/sec1.go
func ecPrivateKeyFromBytes(b []byte, curve elliptic.Curve) (key *ecdsa.PrivateKey, err error) {
	k := new(big.Int).SetBytes(b)
	curveOrder := curve.Params().N
	if k.Cmp(curveOrder) >= 0 {
		return nil, ErrInvalidPrivateKey
	}

	priv := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
		},
		D: k,
	}

	privateKey := make([]byte, (curveOrder.BitLen()+7)/8)
	// Some private keys have leading zero padding. This is invalid
	// according to [SEC1], but this code will ignore it.
	for len(b) > len(privateKey) {
		if b[0] != 0 {
			return nil, ErrInvalidPrivateKey
		}
		b = b[1:]
	}

	// Some private keys remove all leading zeros, this is also invalid
	// according to [SEC1] but since OpenSSL used to do this, we ignore
	// this too.
	copy(privateKey[len(privateKey)-len(b):], b)
	priv.X, priv.Y = curve.ScalarBaseMult(privateKey)

	return &priv, nil
}
