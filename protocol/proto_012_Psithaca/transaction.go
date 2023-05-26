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
	Amount      tz.BigUint
	Destination core.ContractID
	Parameters  tz.Option[Parameters]
}

func (*Transaction) OperationKind() string { return "transaction" }

type Parameters struct {
	Entrypoint Entrypoint
	Value      expression.Expression `tz:"dyn"`
}

type Entrypoint interface {
	Entrypoint()
}

type EpDefault struct{}
type EpRoot struct{}
type EpDo struct{}
type EpSetDelegate struct{}
type EpRemoveDelegate struct{}
type EpNamed struct {
	tz.String
}

func (EpDefault) Entrypoint()        {}
func (EpRoot) Entrypoint()           {}
func (EpDo) Entrypoint()             {}
func (EpSetDelegate) Entrypoint()    {}
func (EpRemoveDelegate) Entrypoint() {}
func (EpNamed) Entrypoint()          {}

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
	Storage                      tz.Option[expression.Expression]
	BigMapDiff                   tz.Option[big_map.Diff]
	BalanceUpdates               []*BalanceUpdate            `tz:"dyn"`
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn"`
	ConsumedGas                  tz.BigUint
	ConsumedMilligas             tz.BigUint
	StorageSize                  tz.BigInt
	PaidStorageSizeDiff          tz.BigInt
	AllocatedDestinationContract bool
	LazyStorageDiff              tz.Option[lazy.StorageDiff]
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult]
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) OperationContents() core.OperationContents {
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
	core.OperationResult
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
	Destination core.ContractID
	Parameters  tz.Option[Parameters]
	Result      TransactionResult
}

func (*TransactionInternalOperationResult) InternalOperationResult() {}
func (*TransactionInternalOperationResult) OperationKind() string    { return "transaction" }
