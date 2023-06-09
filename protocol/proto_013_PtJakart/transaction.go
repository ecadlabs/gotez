package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters

type ScRollupInbox struct {
	Rollup                                 *tz.ScRollupAddress `tz:"dyn" json:"rollup"`
	MessageCounter                         tz.BigInt           `json:"message_counter"`
	NbMessagesInCommitmentPeriod           int64               `json:"nb_messages_in_commitment_period"`
	StartingLevelOfCurrentCommitmentPeriod int32               `json:"starting_level_of_current_commitment_period"`
	Level                                  int32               `json:"level"`
	CurrentLevelHash                       *tz.Bytes32         `json:"current_level_hash"`
	OldLevelsMessages                      OldLevelsMessages   `json:"old_levels_messages"`
}

type OldLevelsMessages struct {
	Index        int32       `json:"index"`
	Content      *tz.Bytes32 `json:"content"`
	BackPointers tz.Bytes    `tz:"dyn" json:"back_pointers"`
}

type TxRollupDestination struct {
	*tz.TXRollupAddress
	Padding uint8
}

func (*TxRollupDestination) TransactionDestination() {}

type TransactionResultDestination interface {
	TransactionResultDestination()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract)(nil),
			1: (*ToTxRollup)(nil),
		},
	})
}

type ToContract struct {
	Storage tz.Option[expression.Expression] `json:"storage"`
	BalanceUpdates
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedGas                  tz.BigUint                  `json:"consumed_gas"`
	ConsumedMilligas             tz.BigUint                  `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                   `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                        `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

func (*ToContract) TransactionResultDestination() {}

type ToTxRollup struct {
	BalanceUpdates
	ConsumedGas         tz.BigUint         `json:"consumed_gas"`
	ConsumedMilligas    tz.BigUint         `json:"consumed_milligas"`
	TicketHash          *tz.ScriptExprHash `json:"ticket_hash"`
	PaidStorageSizeDiff tz.BigUint         `json:"paid_storage_size_diff"`
}

func (*ToTxRollup) TransactionResultDestination() {}

type TransactionDestination interface {
	core.TransactionDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: (*core.ImplicitContract)(nil),
			1: (*core.OriginatedContract)(nil),
			2: (*TxRollupDestination)(nil),
		},
	})
}

type TransactionResultContents struct {
	TransactionResultDestination `json:"contents"`
}

//json:kind=OperationKind()
type TransactionSuccessfulManagerResult struct{ TransactionResultApplied }

func (*TransactionSuccessfulManagerResult) OperationKind() string { return "transaction" }

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult] `json:"metadata"`
}

func (op *TransactionContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}

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

//json:kind=OperationKind()
type TransactionInternalOperationResult struct {
	Source      core.ContractID        `json:"source"`
	Nonce       uint16                 `json:"nonce"`
	Amount      tz.BigUint             `json:"amount"`
	Destination TransactionDestination `json:"destination"`
	Parameters  tz.Option[Parameters]  `json:"parameters"`
	Result      TransactionResult      `json:"result"`
}

func (r *TransactionInternalOperationResult) GetSource() core.Address { return r.Source }
func (r *TransactionInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }
