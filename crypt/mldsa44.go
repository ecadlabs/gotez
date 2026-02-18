package crypt

import (
	"crypto"

	"github.com/cloudflare/circl/sign/mldsa/mldsa44"
	"github.com/ecadlabs/gotez/v2"
	tz "github.com/ecadlabs/gotez/v2"
)

type MLDSA44PrivateKey mldsa44.PrivateKey

func (priv *MLDSA44PrivateKey) ToBase58() []byte {
	return priv.ToProtocol().ToBase58()
}

func (priv *MLDSA44PrivateKey) Public() PublicKey {
	return (*MLDSA44PublicKey)((*mldsa44.PrivateKey)(priv).Public().(*mldsa44.PublicKey))
}

func (priv *MLDSA44PrivateKey) Equal(other PrivateKey) bool {
	x, ok := other.(*MLDSA44PrivateKey)
	return ok && (*mldsa44.PrivateKey)(priv).Equal((*mldsa44.PrivateKey)(x))
}

func (priv *MLDSA44PrivateKey) MarshalText() (text []byte, err error) {
	return priv.ToProtocol().ToBase58(), nil
}

func (priv *MLDSA44PrivateKey) Unwrap() crypto.PrivateKey {
	return (*mldsa44.PrivateKey)(priv)
}

func (priv *MLDSA44PrivateKey) ToProtocol() tz.PrivateKey {
	pk := (*mldsa44.PrivateKey)(priv)
	pub := pk.Public().(*mldsa44.PublicKey)
	out, err := gotez.NewMLDSA44PrivateKey(pk.Bytes(), pub.Bytes())
	if err != nil {
		panic(err)
	}
	return out
}

func (priv *MLDSA44PrivateKey) Sign(message []byte) (signature Signature, err error) {
	sig, err := (*mldsa44.PrivateKey)(priv).Sign(nil, message, nil)
	return MLDSA44Signature(sig), err
}

type MLDSA44PublicKey mldsa44.PublicKey

func (pub *MLDSA44PublicKey) Hash() PublicKeyHash {
	return pub.ToProtocol().Hash()
}

func (pub *MLDSA44PublicKey) ToBase58() []byte {
	return pub.ToProtocol().ToBase58()
}

func (pub *MLDSA44PublicKey) String() string {
	return string(pub.ToBase58())
}

func (pub *MLDSA44PublicKey) VerifySignature(sig Signature, message []byte) bool {
	if sig, ok := sig.(MLDSA44Signature); ok {
		return mldsa44.Verify((*mldsa44.PublicKey)(pub), message, nil, sig)
	}
	return false
}

func (pub *MLDSA44PublicKey) ToProtocol() tz.PublicKey {
	out, err := tz.NewMLDSA44PublicKey((*mldsa44.PublicKey)(pub).Bytes())
	if err != nil {
		panic(err)
	}
	return out
}

func (pub *MLDSA44PublicKey) Equal(other PublicKey) bool {
	x, ok := other.(*MLDSA44PublicKey)
	return ok && (*mldsa44.PublicKey)(pub).Equal((*mldsa44.PublicKey)(x))
}

func (pub *MLDSA44PublicKey) MarshalText() (text []byte, err error) {
	return pub.ToProtocol().ToBase58(), nil
}

func (pub *MLDSA44PublicKey) Unwrap() crypto.PublicKey {
	return (*mldsa44.PublicKey)(pub)
}

type MLDSA44Signature []byte

func (sig MLDSA44Signature) ToBase58() []byte {
	return sig.ToProtocol().ToBase58()
}

func (sig MLDSA44Signature) String() string {
	return string(sig.ToBase58())
}

func (sig MLDSA44Signature) Verify(pub PublicKey, message []byte) bool {
	return pub.VerifySignature(sig, message)
}

func (sig MLDSA44Signature) ToProtocol() tz.Signature {
	return tz.NewMLDSA44Signature(sig)
}

func (sig MLDSA44Signature) MarshalText() (text []byte, err error) {
	return sig.ToProtocol().ToBase58(), nil
}
