package proto_024_PtTALLiN

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
	"github.com/ecadlabs/gotez/v2/protocol/proto_021_PsQuebec"
	"github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"
)

type Transaction = proto_023_PtSeouLo.Transaction
type Parameters = proto_023_PtSeouLo.Parameters
type EpDefault = proto_023_PtSeouLo.EpDefault
type EpRoot = proto_023_PtSeouLo.EpRoot
type EpDo = proto_023_PtSeouLo.EpDo
type EpSetDelegate = proto_023_PtSeouLo.EpSetDelegate
type EpRemoveDelegate = proto_023_PtSeouLo.EpRemoveDelegate
type EpDeposit = proto_023_PtSeouLo.EpDeposit
type EpStake = proto_023_PtSeouLo.EpStake
type EpUnstake = proto_023_PtSeouLo.EpUnstake
type EpFinalizeUnstake = proto_023_PtSeouLo.EpFinalizeUnstake
type EpSetDelegateParameters = proto_023_PtSeouLo.EpSetDelegateParameters
type EpNamed = proto_023_PtSeouLo.EpNamed
type ToSmartRollup = proto_023_PtSeouLo.ToSmartRollup
type PseudoOperation = proto_023_PtSeouLo.PseudoOperation
type TicketReceipt = proto_023_PtSeouLo.TicketReceipt

type AddressRegistryElem struct {
	Address core.TransactionDestination `json:"address"`
	Index   tz.BigInt                   `json:"index"`
}

type ToContract struct {
	Storage tz.Option[expression.Expression] `json:"storage"`
	BalanceUpdates
	TicketUpdates                []*TicketReceipt            `tz:"dyn" json:"ticket_updates"`
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedMilligas             tz.BigUint                  `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                   `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                        `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
	AddressRegistryDiff          []*AddressRegistryElem      `tz:"dyn" json:"address_registry_diff"`
}

func (*ToContract) TransactionResultDestination()       {}
func (r *ToContract) GetConsumedMilligas() tz.BigUint   { return r.ConsumedMilligas }
func (r *ToContract) GetStorageSize() tz.BigInt         { return r.StorageSize }
func (r *ToContract) GetPaidStorageSizeDiff() tz.BigInt { return r.PaidStorageSizeDiff }
func (r *ToContract) EstimateStorageSize(constants core.Constants) *big.Int {
	x := r.PaidStorageSizeDiff.Int()
	if r.AllocatedDestinationContract {
		x.Add(x, big.NewInt(int64(constants.GetOriginationSize())))
	}
	return x
}

type TransactionResultDestination interface {
	proto_021_PsQuebec.TransactionResultDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract)(nil),
			2: (*ToSmartRollup)(nil),
		},
	})
}

type TransactionResultContents = TransactionResultDestination

type TransactionResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*core.OperationResultApplied[TransactionResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[TransactionResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult] `json:"metadata"`
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type TransactionInternalOperationResult struct {
	Source      core.TransactionDestination `json:"source"`
	Nonce       uint16                      `json:"nonce"`
	Amount      tz.BigUint                  `json:"amount"`
	Destination core.TransactionDestination `json:"destination"`
	Parameters  tz.Option[Parameters]       `json:"parameters"`
	Result      TransactionResult           `json:"result"`
}

var _ core.TransactionInternalOperationResult = (*TransactionInternalOperationResult)(nil)

func (r *TransactionInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }
func (r *TransactionInternalOperationResult) GetNonce() uint16                       { return r.Nonce }
func (t *TransactionInternalOperationResult) GetAmount() tz.BigUint                  { return t.Amount }
func (t *TransactionInternalOperationResult) GetDestination() core.TransactionDestination {
	return t.Destination
}
func (t *TransactionInternalOperationResult) GetParameters() tz.Option[core.Parameters] {
	if p, ok := t.Parameters.CheckUnwrapPtr(); ok {
		return tz.Some[core.Parameters](p)
	}
	return tz.None[core.Parameters]()
}
func (r *TransactionInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }

//json:kind=OperationKind()
type TransactionSuccessfulManagerResult struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (TransactionSuccessfulManagerResult) OperationKind() string { return "transaction" }

var ListPseudoOperations = proto_023_PtSeouLo.ListPseudoOperations
