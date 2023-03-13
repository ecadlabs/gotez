package gotez_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/b58"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	type testCase struct {
		title  string
		genKey func() crypto.Signer
	}
	cases := []testCase{
		{
			title:  "Ed25519",
			genKey: func() crypto.Signer { _, k, _ := ed25519.GenerateKey(rand.Reader); return k },
		},
		{
			title:  "Secp256k1",
			genKey: func() crypto.Signer { k, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader); return k },
		},
		{
			title:  "P256",
			genKey: func() crypto.Signer { k, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader); return k },
		},
		{
			title:  "BLS",
			genKey: func() crypto.Signer { k, _ := minpk.GenerateKey(rand.Reader); return k },
		},
	}

	type privateKey interface {
		crypto.Signer
		Equal(x crypto.Signer) bool
	}

	type publicKey interface {
		Equal(x crypto.PublicKey) bool
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// generate key
			priv := c.genKey().(privateKey)
			// encode to internal roundtrip
			tzPriv, err := tz.NewPrivateKey(priv)
			require.NoError(t, err)
			tmp, err := tzPriv.PrivateKey()
			require.NoError(t, err)
			require.True(t, priv.Equal(tmp))
			require.Equal(t, priv, tmp)

			// encode to base58 roundtrip
			tmp2, err := b58.ParsePrivateKey(tzPriv.ToBase58())
			require.NoError(t, err)
			require.Equal(t, tzPriv, tmp2)

			// encode to base58 roundtrip using encrypted type
			tmp3, err := b58.ParsePrivateKey(tzPriv.ToBase58())
			require.NoError(t, err)
			decrypted, err := tmp3.Decrypt(nil)
			require.NoError(t, err)
			require.Equal(t, tzPriv, decrypted)

			// get public
			pub := priv.Public().(publicKey)
			// encode to internal roundtrip
			tzPub, err := tz.NewPublicKey(pub)
			require.NoError(t, err)
			tmp4, err := tzPub.PublicKey()
			require.NoError(t, err)
			require.True(t, pub.Equal(tmp4))
			require.Equal(t, pub, tmp4)

			// encode to base58 roundtrip
			tmp5, err := b58.ParsePublicKey(tzPub.ToBase58())
			require.NoError(t, err)
			require.Equal(t, tzPub, tmp5)
		})
	}
}
