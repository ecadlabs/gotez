package proto_001_PtCJ7pwo

import (
	tz "github.com/ecadlabs/gotez"
)

type ActivateAccount struct {
	PKH    *tz.Ed25519PublicKeyHash
	Secret *[tz.SecretBytesLen]byte
}

func (*ActivateAccount) OperationKind() string { return "activate_account" }

type Proposals struct {
	Source    tz.PublicKeyHash
	Period    int32
	Proposals []*tz.ProtocolHash `tz:"dyn"`
}

func (*Proposals) OperationKind() string       { return "proposals" }
func (*Proposals) OperationContentsAndResult() {}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

type Ballot struct {
	Source   tz.PublicKeyHash
	Period   int32
	Proposal *tz.ProtocolHash
	Ballot   BallotKind
}

func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}
