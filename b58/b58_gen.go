package b58

import (
	"fmt"

	"github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/b58/base58"
	"github.com/ecadlabs/gotez/b58/prefix"
)

// Code generated by generate.go DO NOT EDIT.

func ParseBlockHash(src []byte) (*gotez.BlockHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.BlockHash {
		return nil, fmt.Errorf("gotez: invalid BlockHash encoding")
	}
	var out gotez.BlockHash
	copy(out[:], payload)
	return &out, nil
}

func ParseOperationsHash(src []byte) (*gotez.OperationsHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.OperationListListHash {
		return nil, fmt.Errorf("gotez: invalid OperationsHash encoding")
	}
	var out gotez.OperationsHash
	copy(out[:], payload)
	return &out, nil
}

func ParseContextHash(src []byte) (*gotez.ContextHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.ContextHash {
		return nil, fmt.Errorf("gotez: invalid ContextHash encoding")
	}
	var out gotez.ContextHash
	copy(out[:], payload)
	return &out, nil
}

func ParseChainID(src []byte) (*gotez.ChainID, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.ChainID {
		return nil, fmt.Errorf("gotez: invalid ChainID encoding")
	}
	var out gotez.ChainID
	copy(out[:], payload)
	return &out, nil
}

func ParseBlockPayloadHash(src []byte) (*gotez.BlockPayloadHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.ValueHash {
		return nil, fmt.Errorf("gotez: invalid BlockPayloadHash encoding")
	}
	var out gotez.BlockPayloadHash
	copy(out[:], payload)
	return &out, nil
}

func ParseCycleNonceHash(src []byte) (*gotez.CycleNonceHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.CycleNonce {
		return nil, fmt.Errorf("gotez: invalid CycleNonceHash encoding")
	}
	var out gotez.CycleNonceHash
	copy(out[:], payload)
	return &out, nil
}

func ParseEd25519PublicKeyHash(src []byte) (*gotez.Ed25519PublicKeyHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.Ed25519PublicKeyHash {
		return nil, fmt.Errorf("gotez: invalid Ed25519PublicKeyHash encoding")
	}
	var out gotez.Ed25519PublicKeyHash
	copy(out[:], payload)
	return &out, nil
}

func ParseSecp256k1PublicKeyHash(src []byte) (*gotez.Secp256k1PublicKeyHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.Secp256k1PublicKeyHash {
		return nil, fmt.Errorf("gotez: invalid Secp256k1PublicKeyHash encoding")
	}
	var out gotez.Secp256k1PublicKeyHash
	copy(out[:], payload)
	return &out, nil
}

func ParseP256PublicKeyHash(src []byte) (*gotez.P256PublicKeyHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.P256PublicKeyHash {
		return nil, fmt.Errorf("gotez: invalid P256PublicKeyHash encoding")
	}
	var out gotez.P256PublicKeyHash
	copy(out[:], payload)
	return &out, nil
}

func ParseBLSPublicKeyHash(src []byte) (*gotez.BLSPublicKeyHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.BLS12_381PublicKeyHash {
		return nil, fmt.Errorf("gotez: invalid BLSPublicKeyHash encoding")
	}
	var out gotez.BLSPublicKeyHash
	copy(out[:], payload)
	return &out, nil
}

func ParseProtocolHash(src []byte) (*gotez.ProtocolHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.ProtocolHash {
		return nil, fmt.Errorf("gotez: invalid ProtocolHash encoding")
	}
	var out gotez.ProtocolHash
	copy(out[:], payload)
	return &out, nil
}

func ParseContractHash(src []byte) (*gotez.ContractHash, error) {
	pre, payload, err := base58.DecodeTZ(src)
	if err != nil {
		return nil, err
	}
	if pre != &prefix.ContractHash {
		return nil, fmt.Errorf("gotez: invalid ContractHash encoding")
	}
	var out gotez.ContractHash
	copy(out[:], payload)
	return &out, nil
}
