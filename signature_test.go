package gotez_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	blst "github.com/ecadlabs/goblst"
	"github.com/ecadlabs/goblst/minpk"
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/b58"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/blake2b"
)

func TestSignature(t *testing.T) {
	var message = []byte("message text")

	t.Run("Ed25519", func(t *testing.T) {
		_, priv, err := ed25519.GenerateKey(rand.Reader)
		require.NoError(t, err)
		sig := ed25519.Sign(priv, message)
		tzSig := tz.NewEd25519Signature(sig)
		tmp, err := b58.ParseSignature(tzSig.ToBase58())
		require.NoError(t, err)
		require.Equal(t, tzSig, tmp.(*tz.Ed25519Signature))
	})

	t.Run("Secp256k1", func(t *testing.T) {
		digest := blake2b.Sum256(message)
		priv, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		require.NoError(t, err)
		r, s, err := ecdsa.Sign(rand.Reader, priv, digest[:])
		require.NoError(t, err)
		tzSig := tz.NewSecp256k1Signature(r, s)
		tmp, err := b58.ParseSignature(tzSig.ToBase58())
		require.NoError(t, err)
		tmp1 := tmp.(*tz.Secp256k1Signature)
		require.Equal(t, tzSig, tmp1)
		rNew, sNew := tmp1.Get()
		require.True(t, rNew.Cmp(r) == 0)
		require.True(t, sNew.Cmp(s) == 0)
	})

	t.Run("P256", func(t *testing.T) {
		digest := blake2b.Sum256(message)
		priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)
		r, s, err := ecdsa.Sign(rand.Reader, priv, digest[:])
		require.NoError(t, err)
		tzSig := tz.NewP256Signature(r, s)
		tmp, err := b58.ParseSignature(tzSig.ToBase58())
		require.NoError(t, err)
		tmp1 := tmp.(*tz.P256Signature)
		require.Equal(t, tzSig, tmp1)
		rNew, sNew := tmp1.Get()
		require.True(t, rNew.Cmp(r) == 0)
		require.True(t, sNew.Cmp(s) == 0)
	})

	t.Run("BLS", func(t *testing.T) {
		priv, err := minpk.GenerateKey(rand.Reader)
		require.NoError(t, err)
		sig := minpk.Sign(priv, message, blst.Basic)
		tzSig := tz.NewBLSSignature(sig)
		tmp, err := b58.ParseSignature(tzSig.ToBase58())
		require.NoError(t, err)
		tmp1 := tmp.(*tz.BLSSignature)
		require.Equal(t, tzSig, tmp1)
		sigNew, err := tmp1.Get()
		require.NoError(t, err)
		require.Equal(t, sig, sigNew)
	})
}
