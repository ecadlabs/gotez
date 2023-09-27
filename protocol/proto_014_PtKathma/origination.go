package proto_014_PtKathma

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
)

type Origination = proto_012_Psithaca.Origination
type Script = proto_012_Psithaca.Script

type OriginationResultContents struct {
	BalanceUpdates
	OriginatedContracts []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
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

//json:kind=OperationKind()
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
	Source   core.TransactionDestination `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Balance  tz.BigUint                  `json:"balance"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
	Result   OriginationResult           `json:"result"`
}

var (
	_ core.OperationWithResult = (*OriginationInternalOperationResult)(nil)
	_ core.OperationWithSource = (*OriginationInternalOperationResult)(nil)
)

func (r *OriginationInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }
func (r *OriginationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*OriginationInternalOperationResult) OperationKind() string { return "origination" }
