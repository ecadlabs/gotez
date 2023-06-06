package proto_013_PtJakart

import (
	"encoding/json"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca/lazy"
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
func (t *TxRollupDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.TXRollupAddress)
}

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
	Storage                      tz.Option[expression.Expression] `json:"storage"`
	BalanceUpdates               []*BalanceUpdate                 `tz:"dyn" json:"balance_updates"`
	OriginatedContracts          []core.OriginatedContractID      `tz:"dyn" json:"originated_contracts"`
	ConsumedGas                  tz.BigUint                       `json:"consumed_gas"`
	ConsumedMilligas             tz.BigUint                       `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                        `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                        `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                             `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff]      `json:"lazy_storage_diff"`
}

func (*ToContract) TransactionResultDestination() {}

type ToTxRollup struct {
	BalanceUpdates      []*BalanceUpdate   `tz:"dyn" json:"balance_updates"`
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
	TransactionResultDestination
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }
func (c TransactionResultContents) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.TransactionResultDestination)
}

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult]
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
	Source      core.ContractID
	Nonce       uint16
	Amount      tz.BigUint
	Destination TransactionDestination
	Parameters  tz.Option[Parameters]
	Result      TransactionResult
}

func (r *TransactionInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }
