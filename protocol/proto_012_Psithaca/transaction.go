package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca/big_map"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca/lazy"
)

type Transaction struct {
	ManagerOperation
	Amount      tz.BigUint            `json:"amount"`
	Destination core.ContractID       `json:"destination"`
	Parameters  tz.Option[Parameters] `json:"parameters"`
}

func (*Transaction) OperationKind() string             { return "transaction" }
func (op *Transaction) Operation() core.Operation      { return op }
func (t *Transaction) GetAmount() tz.BigUint           { return t.Amount }
func (t *Transaction) GetDestination() core.ContractID { return t.Destination }
func (t *Transaction) GetParameters() tz.Option[core.Parameters] {
	if p, ok := t.Parameters.CheckUnwrapPtr(); ok {
		return tz.Some[core.Parameters](p)
	}
	return tz.None[core.Parameters]()
}

var _ core.Transaction = (*Transaction)(nil)

type Parameters struct {
	Entrypoint Entrypoint            `json:"entrypoint"`
	Value      expression.Expression `tz:"dyn" json:"value"`
}

func (p *Parameters) GetEntrypoint() string           { return p.Entrypoint.Entrypoint() }
func (p *Parameters) GetValue() expression.Expression { return p.Value }

type Entrypoint interface {
	Entrypoint() string
}

type EpDefault struct{}
type EpRoot struct{}
type EpDo struct{}
type EpSetDelegate struct{}
type EpRemoveDelegate struct{}
type EpNamed struct {
	tz.String
}

func (EpDefault) Entrypoint() string                           { return "default" }
func (ep EpDefault) MarshalText() (text []byte, err error)     { return []byte(ep.Entrypoint()), nil }
func (EpRoot) Entrypoint() string                              { return "root" }
func (ep EpRoot) MarshalText() (text []byte, err error)        { return []byte(ep.Entrypoint()), nil }
func (EpDo) Entrypoint() string                                { return "do" }
func (ep EpDo) MarshalText() (text []byte, err error)          { return []byte(ep.Entrypoint()), nil }
func (EpSetDelegate) Entrypoint() string                       { return "set_delegate" }
func (ep EpSetDelegate) MarshalText() (text []byte, err error) { return []byte(ep.Entrypoint()), nil }
func (EpRemoveDelegate) Entrypoint() string                    { return "remove_delegate" }
func (ep EpRemoveDelegate) MarshalText() (text []byte, err error) {
	return []byte(ep.Entrypoint()), nil
}
func (e EpNamed) Entrypoint() string                     { return string(e.String) }
func (ep EpNamed) MarshalText() (text []byte, err error) { return []byte(ep.Entrypoint()), nil }

func init() {
	encoding.RegisterEnum(&encoding.Enum[Entrypoint]{
		Variants: encoding.Variants[Entrypoint]{
			0:   EpDefault{},
			1:   EpRoot{},
			2:   EpDo{},
			3:   EpSetDelegate{},
			4:   EpRemoveDelegate{},
			255: EpNamed{},
		},
	})
}

type TransactionResultContents struct {
	Storage                      tz.Option[expression.Expression] `json:"storage"`
	BigMapDiff                   tz.Option[big_map.Diff]          `json:"big_map_diff"`
	BalanceUpdates               []*BalanceUpdate                 `tz:"dyn" json:"balance_updates"`
	OriginatedContracts          []core.OriginatedContractID      `tz:"dyn" json:"originated_contracts"`
	ConsumedGas                  tz.BigUint                       `json:"consumed_gas"`
	ConsumedMilligas             tz.BigUint                       `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                        `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                        `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                             `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff]      `json:"lazy_storage_diff"`
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }

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
	TransactionResult()
	core.ManagerOperationResult
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
	Source      core.ContractID       `json:"source"`
	Nonce       uint16                `json:"nonce"`
	Amount      tz.BigUint            `json:"amount"`
	Destination core.ContractID       `json:"destination"`
	Parameters  tz.Option[Parameters] `json:"parameters"`
	Result      TransactionResult     `json:"result"`
}

func (r *TransactionInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }
