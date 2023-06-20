package proto_012_Psithaca

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/big_map"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
)

//json:kind=OperationKind()
type Origination struct {
	ManagerOperation
	Balance  tz.BigUint                  `json:"balance"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
}

func (*Origination) OperationKind() string { return "origination" }

type OriginationResultContents struct {
	BigMapDiff tz.Option[big_map.Diff] `json:"big_map_diff"`
	BalanceUpdates
	OriginatedContracts []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedGas         tz.BigUint                  `json:"consumed_gas"`
	ConsumedMilligas    tz.BigUint                  `json:"consumed_milligas"`
	StorageSize         tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff tz.BigInt                   `json:"paid_storage_size_diff"`
	LazyStorageDiff     tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

func (r *OriginationResultContents) GetConsumedMilligas() tz.BigUint   { return r.ConsumedMilligas }
func (r *OriginationResultContents) GetStorageSize() tz.BigInt         { return r.StorageSize }
func (r *OriginationResultContents) GetPaidStorageSizeDiff() tz.BigInt { return r.PaidStorageSizeDiff }
func (r *OriginationResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	x := r.PaidStorageSizeDiff.Int()
	x.Add(x, big.NewInt(int64(constants.GetOriginationSize())))
	return x
}

//json:kind=OperationKind()
type OriginationSuccessfulManagerResult struct {
	core.OperationResultApplied[*OriginationResultContents]
}

func (*OriginationSuccessfulManagerResult) OperationKind() string { return "origination" }

type OriginationResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginationResult]{
		Variants: encoding.Variants[OriginationResult]{
			0: (*core.OperationResultApplied[*OriginationResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*OriginationResultContents])(nil),
		},
	})
}

type OriginationContentsAndResult struct {
	Origination
	Metadata ManagerMetadata[OriginationResult] `json:"metadata"`
}

func (*OriginationContentsAndResult) OperationContentsAndResult() {}
func (op *OriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type OriginationInternalOperationResult struct {
	Source   core.ContractID             `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Balance  tz.BigUint                  `json:"balance"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
	Result   OriginationResult           `json:"result"`
}

func (o *OriginationInternalOperationResult) GetSource() core.Address { return o.Source }
func (r *OriginationInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*OriginationInternalOperationResult) OperationKind() string { return "origination" }
