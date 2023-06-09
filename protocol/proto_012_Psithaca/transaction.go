package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/big_map"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
)

//json:kind=OperationKind()
type Transaction struct {
	ManagerOperation
	Amount      tz.BigUint            `json:"amount"`
	Destination core.ContractID       `json:"destination"`
	Parameters  tz.Option[Parameters] `json:"parameters"`
}

func (*Transaction) OperationKind() string             { return "transaction" }
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
	Storage    tz.Option[expression.Expression] `json:"storage"`
	BigMapDiff tz.Option[big_map.Diff]          `json:"big_map_diff"`
	BalanceUpdates
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedGas                  tz.BigUint                  `json:"consumed_gas"`
	ConsumedMilligas             tz.BigUint                  `json:"consumed_milligas"`
	StorageSize                  tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff          tz.BigInt                   `json:"paid_storage_size_diff"`
	AllocatedDestinationContract bool                        `json:"allocated_destination_contract"`
	LazyStorageDiff              tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

//json:kind=OperationKind()
type TransactionSuccessfulManagerResult struct {
	core.OperationResultApplied[*TransactionResultContents]
}

func (*TransactionSuccessfulManagerResult) OperationKind() string { return "transaction" }

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult] `json:"metadata"`
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type TransactionResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*core.OperationResultApplied[*TransactionResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*TransactionResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type TransactionInternalOperationResult struct {
	Source      core.ContractID       `json:"source"`
	Nonce       uint16                `json:"nonce"`
	Amount      tz.BigUint            `json:"amount"`
	Destination core.ContractID       `json:"destination"`
	Parameters  tz.Option[Parameters] `json:"parameters"`
	Result      TransactionResult     `json:"result"`
}

func (r *TransactionInternalOperationResult) GetSource() core.Address { return r.Source }
func (r *TransactionInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*TransactionInternalOperationResult) OperationKind() string { return "transaction" }
