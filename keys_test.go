package gotez

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	type testCase struct {
		title           string
		genKey          func() crypto.PrivateKey
		decodeBase58    func([]byte) (PrivateKey, error)
		decodePubBase58 func([]byte) (PublicKey, error)
	}
	cases := []testCase{
		{
			title:           "ed25519",
			genKey:          func() crypto.PrivateKey { _, k, _ := ed25519.GenerateKey(rand.Reader); return k },
			decodeBase58:    func(src []byte) (PrivateKey, error) { return NewEd25519PrivateKeyFromBase58(src) },
			decodePubBase58: func(src []byte) (PublicKey, error) { return NewEd25519PublicKeyFromBase58(src) },
		},
		{
			title:           "secp256k1",
			genKey:          func() crypto.PrivateKey { k, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader); return k },
			decodeBase58:    func(src []byte) (PrivateKey, error) { return NewSecp256k1PrivateKeyFromBase58(src) },
			decodePubBase58: func(src []byte) (PublicKey, error) { return NewSecp256k1PublicKeyFromBase58(src) },
		},
		{
			title:           "p256",
			genKey:          func() crypto.PrivateKey { k, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader); return k },
			decodeBase58:    func(src []byte) (PrivateKey, error) { return NewP256PrivateKeyFromBase58(src) },
			decodePubBase58: func(src []byte) (PublicKey, error) { return NewP256PublicKeyFromBase58(src) },
		},
		{
			title:           "bls",
			genKey:          func() crypto.PrivateKey { k, _ := minpk.GenerateKey(rand.Reader); return k },
			decodeBase58:    func(src []byte) (PrivateKey, error) { return NewBLSPrivateKeyFromBase58(src) },
			decodePubBase58: func(src []byte) (PublicKey, error) { return NewBLSPublicKeyFromBase58(src) },
		},
	}

	type privateKey interface {
		Public() crypto.PublicKey
		Equal(x crypto.PrivateKey) bool
	}

	type publicKey interface {
		Equal(x crypto.PublicKey) bool
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// generate key
			priv := c.genKey().(privateKey)
			// encode to internal roundtrip
			tzPriv, err := NewPrivateKeyFrom(priv)
			require.NoError(t, err)
			tmp, err := tzPriv.PrivateKey()
			require.NoError(t, err)
			require.True(t, priv.Equal(tmp))
			require.Equal(t, priv, tmp)

			// encode to base58 roundtrip
			tmp2, err := NewPrivateKeyFromBase58(tzPriv.Base58())
			require.NoError(t, err)
			require.Equal(t, tzPriv, tmp2)

			// encode to base58 roundtrip using encrypted type
			tmp3, err := NewEncryptedPrivateKeyFromBase58(tzPriv.Base58())
			require.NoError(t, err)
			decrypted, err := tmp3.Decrypt(nil)
			require.NoError(t, err)
			require.Equal(t, tzPriv, decrypted)

			// encode to base58 roundtrip using type specific call
			tmp6, err := c.decodeBase58(tzPriv.Base58())
			require.NoError(t, err)
			require.Equal(t, tzPriv, tmp6)

			// get public
			pub := priv.Public().(publicKey)
			// encode to internal roundtrip
			tzPub, err := NewPublicKeyFrom(pub)
			require.NoError(t, err)
			tmp4, err := tzPub.PublicKey()
			require.NoError(t, err)
			require.True(t, pub.Equal(tmp4))
			require.Equal(t, pub, tmp4)

			// encode to base58 roundtrip
			tmp5, err := NewPublicKeyFromBase58(tzPub.Base58())
			require.NoError(t, err)
			require.Equal(t, tzPub, tmp5)

			// encode to base58 roundtrip using type specific call
			tmp7, err := c.decodePubBase58(tzPub.Base58())
			require.NoError(t, err)
			require.Equal(t, tzPub, tmp7)
		})
	}
}
