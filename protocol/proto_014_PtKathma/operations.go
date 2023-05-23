package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination core.OriginatedContractID
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	core.OperationResult
}

type VDFRevelation struct {
	Solution [2]*[100]byte
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }

type DALSlotAvailability struct {
	Endorser    tz.PublicKeyHash
	Endorsement tz.BigUint
}

func (*DALSlotAvailability) OperationKind() string { return "dal_slot_availability" }

type DALSlotAvailabilityContentsAndResult struct {
	DALSlotAvailability
	Metadata DALSlotAvailabilityMetadata
}

func (*DALSlotAvailabilityContentsAndResult) OperationContentsAndResult() {}
func (op *DALSlotAvailabilityContentsAndResult) OperationContents() core.OperationContents {
	return &op.DALSlotAvailability
}

type DALSlotAvailabilityMetadata struct {
	Delegate tz.PublicKeyHash
}
