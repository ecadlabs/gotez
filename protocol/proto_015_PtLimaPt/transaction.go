package proto_015_PtLimaPt

import (
	"encoding/json"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca/lazy"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters
type TxRollupDestination = proto_013_PtJakart.TxRollupDestination
type ScRollupDestination = proto_014_PtKathma.ScRollupDestination
type ToScRollup = proto_014_PtKathma.ToScRollup

type TransactionResultDestination interface {
	proto_013_PtJakart.TransactionResultDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract)(nil),
			1: (*ToTxRollup)(nil),
			2: (*ToScRollup)(nil),
		},
	})
}

type ToContract struct {
	Storage                      tz.Option[expression.Expression] `json:"storage"`
	BalanceUpdates               []*BalanceUpdate                 `tz:"dyn" json:"balance_updates"`
	TicketUpdates                []*TicketReceipt                 `tz:"dyn" json:"ticket_updates"`
	OriginatedContracts          []core.OriginatedContractID      `tz:"dyn" json:"originated_contracts"`
	ConsumedMilligas             tz.BigUint                       `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                        `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                        `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                             `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff]      `json:"lazy_storage_diff"`
}

func (*ToContract) TransactionResultDestination() {}

type ToTxRollup struct {
	BalanceUpdates      []*BalanceUpdate   `tz:"dyn" json:"balance_updates"`
	ConsumedMilligas    tz.BigUint         `json:"consumed_milligas"`
	TicketHash          *tz.ScriptExprHash `json:"ticket_hash"`
	PaidStorageSizeDiff tz.BigUint         `json:"paid_storage_size_diff"`
}

func (*ToTxRollup) TransactionResultDestination() {}

type TicketReceipt struct {
	TicketToken TicketToken            `json:"ticket_token"`
	Updates     []*TicketReceiptUpdate `tz:"dyn" json:"updates"`
}

type TicketToken struct {
	Ticketer    core.ContractID       `json:"ticketer"`
	ContentType expression.Expression `json:"content_type"`
	Content     expression.Expression `json:"content"`
}

type TicketReceiptUpdate struct {
	Account TransactionDestination `json:"account"`
	Amount  tz.BigInt              `json:"amount"`
}

type TransactionDestination interface {
	core.TransactionDestination
}

type ZkRollupDestination struct {
	*tz.ZkRollupAddress
	Padding uint8
}

func (*ZkRollupDestination) TransactionDestination() {}
func (z *ZkRollupDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(z.ZkRollupAddress)
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: (*core.ImplicitContract)(nil),
			1: (*core.OriginatedContract)(nil),
			2: (*TxRollupDestination)(nil),
			3: (*ScRollupDestination)(nil),
			4: (*ZkRollupDestination)(nil),
		},
	})
}

type TransactionResultContents struct {
	TransactionResultDestination
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }
func (c TransactionResultContents) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.TransactionResultDestination)
}

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult] `json:"metadata"`
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) Operation() core.Operation {
	return &op.Transaction
}

type TransactionResultApplied struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (*TransactionResultApplied) TransactionResult() {}

type TransactionResultBacktracked struct {
	core.OperationResultBacktracked[TransactionResultContents]
}

func (*TransactionResultBacktracked) TransactionResult() {}

type TransactionResultFailed struct{ core.OperationResultFailed }

func (*TransactionResultFailed) TransactionResult() {}

type TransactionResultSkipped struct{ core.OperationResultSkipped }

func (*TransactionResultSkipped) TransactionResult() {}

type TransactionResult interface {
	proto_012_Psithaca.TransactionResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*TransactionResultApplied)(nil),
			1: (*TransactionResultFailed)(nil),
			2: (*TransactionResultSkipped)(nil),
			3: (*TransactionResultBacktracked)(nil),
		},
	})
}

type TransactionInternalOperationResult struct {
	Source      core.ContractID        `json:"source"`
	Nonce       uint16                 `json:"nonce"`
	Amount      tz.BigUint             `json:"amount"`
	Destination TransactionDestination `json:"destination"`
	Parameters  tz.Option[Parameters]  `json:"parameters"`
	Result      TransactionResult      `json:"result"`
}

func (r *TransactionInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }
