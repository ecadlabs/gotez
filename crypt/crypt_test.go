package crypt

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"sync"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ecadlabs/goblst/minpk"
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/b58"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	title  string
	genKey func() PrivateKey
}

var cases = []testCase{
	{
		title: "Ed25519",
		genKey: func() PrivateKey {
			_, k, _ := ed25519.GenerateKey(rand.Reader)
			return Ed25519PrivateKey(k)
		},
	},
	{
		title: "Secp256k1",
		genKey: func() PrivateKey {
			k, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
			return (*ECDSAPrivateKey)(k)
		},
	},
	{
		title: "P256",
		genKey: func() PrivateKey {
			k, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			return (*ECDSAPrivateKey)(k)
		},
	},
	{
		title: "BLS",
		genKey: func() PrivateKey {
			k, _ := minpk.GenerateKey(rand.Reader)
			return (*BLSPrivateKey)(k)
		},
	},
}

func TestKey(t *testing.T) {
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			// generate key
			priv := c.genKey()
			// encode to internal roundtrip
			tzPriv := priv.ToProtocol()
			tmp, err := NewPrivateKey(tzPriv)
			require.NoError(t, err)
			require.True(t, priv.Equal(tmp))
			require.Equal(t, priv, tmp)

			// encode to base58 roundtrip
			tmp2, err := ParsePrivateKey(priv.ToBase58())
			require.NoError(t, err)
			require.True(t, priv.Equal(tmp2))

			// encode to base58 roundtrip using encrypted type
			tmp3, err := b58.ParsePrivateKey(tzPriv.ToBase58())
			require.NoError(t, err)
			decrypted, err := tmp3.Decrypt(nil)
			require.NoError(t, err)
			require.Equal(t, tzPriv, decrypted)

			// get public
			pub := priv.Public()
			// encode to internal roundtrip
			tzPub := pub.ToProtocol()
			tmp4, err := NewPublicKey(tzPub)
			require.NoError(t, err)
			require.True(t, pub.Equal(tmp4))
			require.Equal(t, pub, tmp4)

			// encode to base58 roundtrip
			tmp5, err := ParsePublicKey(pub.ToBase58())
			require.NoError(t, err)
			require.True(t, pub.Equal(tmp5))
		})
	}
}

func asGeneric(sig tz.Signature) tz.Signature {
	switch sig := sig.(type) {
	case *tz.Ed25519Signature:
		return (*tz.GenericSignature)(sig)
	case *tz.Secp256k1Signature:
		return (*tz.GenericSignature)(sig)
	case *tz.P256Signature:
		return (*tz.GenericSignature)(sig)
	case *tz.BLSSignature:
		return nil
	default:
		panic("unknown")
	}
}

func TestSignature(t *testing.T) {
	var message = []byte("message text")

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			priv := c.genKey()
			sig, err := priv.Sign(message)
			require.NoError(t, err)

			sig1, err := NewSignature(sig.ToProtocol())
			require.NoError(t, err)
			require.Equal(t, sig, sig1)

			sig2, err := ParseSignature(sig.ToBase58())
			require.NoError(t, err)
			require.Equal(t, sig, sig2)

			require.True(t, sig.Verify(priv.Public(), message))

			if priv, ok := priv.(*BLSPrivateKey); ok {
				sig3, err := priv.SignAugmented(message)
				require.NoError(t, err)
				require.True(t, sig3.(*BLSSignature).VerifyAugmented(priv.Public(), message))
			}

			// via generic
			if genSig := asGeneric(sig.ToProtocol()); genSig != nil {
				sig, err := NewSignature(genSig)
				require.NoError(t, err)
				require.True(t, sig.Verify(priv.Public(), message))
			}
		})
	}
}

var wg sync.WaitGroup

func MinpkGenkey(t *testing.T) {
	defer wg.Done()
	k, err := minpk.GenerateKey(rand.Reader)
	require.NoError(t, err)
	require.NotNil(t, k)
}

func TestMinpkGenkey1kSerial(t *testing.T) {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		MinpkGenkey(t)
	}
}

func TestMinpkGenkey1kParallel(t *testing.T) {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go MinpkGenkey(t)
	}
	wg.Wait()
}

// TestLowSCanonization verifies that ECDSA signatures are normalized to low-S form.
// This is required for secp256k1 (tz2) compatibility with Tezos/libsecp256k1.
// P-256 (tz3) does not require canonization and is returned unchanged for performance.
// See issue #364: https://github.com/ecadlabs/signatory/issues/364
func TestLowSCanonization(t *testing.T) {
	curves := []struct {
		name           string
		curve          elliptic.Curve
		shouldCanonize bool
	}{
		{"secp256k1", secp256k1.S256(), true},
		{"P256", elliptic.P256(), false},
	}

	for _, c := range curves {
		t.Run(c.name, func(t *testing.T) {
			order := c.curve.Params().N
			halfOrder := new(big.Int).Quo(order, big.NewInt(2))

			// Create a high-S value (S > N/2)
			highS := new(big.Int).Add(halfOrder, big.NewInt(1000))
			r := big.NewInt(12345)

			t.Run("NewECDSASignature_canonizes_highS", func(t *testing.T) {
				sig := NewECDSASignature(r, highS, c.curve)
				if c.shouldCanonize {
					// secp256k1: S should now be <= N/2
					require.LessOrEqual(t, sig.S.Cmp(halfOrder), 0,
						"NewECDSASignature should produce low-S signature for secp256k1")
					// Verify S was actually changed
					require.NotEqual(t, 0, highS.Cmp(sig.S),
						"high-S should be different from canonized S")
				} else {
					// P-256: S should remain unchanged (no canonization)
					require.Equal(t, 0, highS.Cmp(sig.S),
						"P-256 signatures should not be canonized, S should remain unchanged")
				}
			})

			t.Run("NewECDSASignature_preserves_lowS", func(t *testing.T) {
				lowS := big.NewInt(1000) // definitely < N/2
				sig := NewECDSASignature(r, lowS, c.curve)
				require.Equal(t, 0, lowS.Cmp(sig.S),
					"low-S should remain unchanged")
			})

			t.Run("canonizeSignature_math_correctness", func(t *testing.T) {
				sig := &ECDSASignature{R: r, S: new(big.Int).Set(highS), Curve: c.curve}
				canonical := canonizeSignature(sig)

				if c.shouldCanonize {
					// secp256k1: When S > N/2, canonical S should be N - S
					expectedS := new(big.Int).Sub(order, highS)
					require.Equal(t, 0, expectedS.Cmp(canonical.S),
						"canonical S should equal N - S for high-S values")
				} else {
					// P-256: Signature should be returned unchanged
					require.Equal(t, 0, highS.Cmp(canonical.S),
						"P-256 signatures should not be canonized, S should remain unchanged")
					require.Equal(t, sig, canonical,
						"P-256 signature should be returned unchanged (same pointer)")
				}
			})
		})
	}
}

// TestNewSignatureFromBytesCanonization verifies that signatures parsed from
// ASN.1 (as received from cloud KMS providers) are canonized to low-S form for secp256k1.
// P-256 signatures are returned unchanged (no canonization) for performance.
func TestNewSignatureFromBytesCanonization(t *testing.T) {
	curves := []struct {
		name           string
		curve          elliptic.Curve
		shouldCanonize bool
	}{
		{"secp256k1", secp256k1.S256(), true},
		{"P256", elliptic.P256(), false},
	}

	for _, c := range curves {
		t.Run(c.name, func(t *testing.T) {
			// Generate a real key and sign to get valid signature bytes
			priv, err := ecdsa.GenerateKey(c.curve, rand.Reader)
			require.NoError(t, err)

			message := []byte("test message for signature canonization")
			digest := DigestFunc(message)

			// Sign multiple times - statistically ~50% should have high-S originally
			halfOrder := new(big.Int).Quo(c.curve.Params().N, big.NewInt(2))

			for i := 0; i < 20; i++ {
				r, s, err := ecdsa.Sign(rand.Reader, priv, digest[:])
				require.NoError(t, err)

				// Create ASN.1 DER encoded signature (as cloud KMS would return)
				asn1Sig := encodeASN1Signature(r, s)

				pub := (*ECDSAPublicKey)(&priv.PublicKey)
				sig, err := NewSignatureFromBytes(asn1Sig, pub)
				require.NoError(t, err)

				ecdsaSig, ok := sig.(*ECDSASignature)
				require.True(t, ok)

				if c.shouldCanonize {
					// secp256k1: Verify S is in low-S form
					require.LessOrEqual(t, ecdsaSig.S.Cmp(halfOrder), 0,
						"NewSignatureFromBytes must produce low-S signature for secp256k1 (iteration %d)", i)
				} else {
					// P-256: S can be high or low (no canonization required)
					// Just verify the signature is valid
				}

				// Verify signature still validates
				require.True(t, sig.Verify(pub, message),
					"signature must still verify (iteration %d)", i)
			}
		})
	}
}

// encodeASN1Signature creates ASN.1 DER encoded signature from R, S values
// (simulating what cloud KMS providers return)
func encodeASN1Signature(r, s *big.Int) []byte {
	rBytes := r.Bytes()
	sBytes := s.Bytes()

	// Add leading zero if high bit is set (ASN.1 integer is signed)
	if len(rBytes) > 0 && rBytes[0]&0x80 != 0 {
		rBytes = append([]byte{0}, rBytes...)
	}
	if len(sBytes) > 0 && sBytes[0]&0x80 != 0 {
		sBytes = append([]byte{0}, sBytes...)
	}

	// Build ASN.1 SEQUENCE
	innerLen := 2 + len(rBytes) + 2 + len(sBytes)
	result := make([]byte, 0, 2+innerLen)

	// SEQUENCE tag and length
	result = append(result, 0x30)
	result = append(result, byte(innerLen))

	// INTEGER for R
	result = append(result, 0x02, byte(len(rBytes)))
	result = append(result, rBytes...)

	// INTEGER for S
	result = append(result, 0x02, byte(len(sBytes)))
	result = append(result, sBytes...)

	return result
}
