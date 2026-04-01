package crypt

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/cloudflare/circl/sign/mldsa/mldsa44"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	"github.com/ecadlabs/gotez/v2"
	"github.com/stretchr/testify/require"
)

func TestPKHType(t *testing.T) {
	t.Run("Ed25519", func(t *testing.T) {
		p, _, _ := ed25519.GenerateKey(rand.Reader)
		pub := Ed25519PublicKey(p)
		_, ok := pub.Hash().(*gotez.Ed25519PublicKeyHash)
		require.True(t, ok, "public key hash has type *gotez.Ed25519PublicKeyHash")
	})
	t.Run("Secp256k1", func(t *testing.T) {
		k, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		pub := (*ECDSAPublicKey)(&k.PublicKey)
		_, ok := pub.Hash().(*gotez.Secp256k1PublicKeyHash)
		require.True(t, ok, "public key hash has type *gotez.Secp256k1PublicKeyHash")
	})
	t.Run("P256", func(t *testing.T) {
		k, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		pub := (*ECDSAPublicKey)(&k.PublicKey)
		_, ok := pub.Hash().(*gotez.P256PublicKeyHash)
		require.True(t, ok, "public key hash has type *gotez.P256PublicKeyHash")
	})
	t.Run("BLS", func(t *testing.T) {
		k, _ := minpk.GenerateKey(rand.Reader)
		pub := (*BLSPublicKey)(k.PublicKey())
		_, ok := pub.Hash().(*gotez.BLSPublicKeyHash)
		require.True(t, ok, "public key hash has type *gotez.BLSPublicKeyHash")
	})
	t.Run("MLDSA44", func(t *testing.T) {
		p, _, _ := mldsa44.GenerateKey(rand.Reader)
		pub := (*MLDSA44PublicKey)(p)
		_, ok := pub.Hash().(*gotez.MLDSA44PublicKeyHash)
		require.True(t, ok, "public key hash has type *gotez.MLDSA44PublicKeyHash")
	})
}
