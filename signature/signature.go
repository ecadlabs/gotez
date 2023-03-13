// Package signature wraps native signature types

package signature

import (
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Signature any

// ECDSA is a type representing an ecdsa signature.
type ECDSA struct {
	R     *big.Int
	S     *big.Int
	Curve elliptic.Curve
}

func (e *ECDSA) String() string {
	return fmt.Sprintf("ecdsa:[c:%s,r:%x,s:%x]", e.Curve.Params().Name, e.R, e.S)
}

// ED25519 is a type representing an Ed25519 signature
type ED25519 []byte

func (e ED25519) String() string {
	return fmt.Sprintf("ed25519:[%s]", hex.EncodeToString(e))
}
