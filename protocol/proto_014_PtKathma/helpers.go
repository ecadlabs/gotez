package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez/v2"
)

type UnsignedOperation struct {
	Branch   *tz.BlockHash       `json:"branch"`
	Contents []OperationContents `json:"contents"`
}

type SignedOperation struct {
	UnsignedOperation
	Signature *tz.GenericSignature `json:"signature"`
}

type RunOperationRequest struct {
	Operation SignedOperation `json:"operation"`
	ChainID   *tz.ChainID     `json:"chain_id"`
}
