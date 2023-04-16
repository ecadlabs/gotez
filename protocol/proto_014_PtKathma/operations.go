package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
)

type ManagerOperation = proto_005_PsBABY5H.ManagerOperation

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination tz.OriginatedContractID
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type VDFRevelation struct {
	Solution [2]*[100]byte
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }
