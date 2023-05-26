package lazy

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type StorageDiff struct {
	Contents []DiffKind `tz:"dyn"`
}

type DiffKind interface {
	LazyStorageDiffKind() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[DiffKind]{
		Variants: encoding.Variants[DiffKind]{
			0: (*BigMap)(nil),
			1: (*SaplingState)(nil),
		},
	})
}

type BigMap struct {
	ID   tz.BigInt
	Diff BigMapOp
}

func (*BigMap) LazyStorageDiffKind() string { return "big_map" }

type BigMapOp interface {
	LazyStorageBigMapOp() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BigMapOp]{
		Variants: encoding.Variants[BigMapOp]{
			0: (*BigMapUpdate)(nil),
			1: (*BigMapRemove)(nil),
			2: (*BigMapCopy)(nil),
			3: (*BigMapAlloc)(nil),
		},
	})
}

type BigMapUpdate struct {
	Updates []*KeyValue `tz:"dyn"`
}

func (*BigMapUpdate) LazyStorageBigMapOp() string { return "update" }

type BigMapRemove struct{}

func (*BigMapRemove) LazyStorageBigMapOp() string { return "remove" }

type BigMapCopy struct {
	Source  tz.BigInt
	Updates []*KeyValue `tz:"dyn"`
}

func (*BigMapCopy) LazyStorageBigMapOp() string { return "copy" }

type BigMapAlloc struct {
	Updates   []*KeyValue `tz:"dyn"`
	KeyType   expression.Expression
	ValueType expression.Expression
}

func (*BigMapAlloc) LazyStorageBigMapOp() string { return "alloc" }

type KeyValue struct {
	KeyHash *tz.ScriptExprHash
	Key     expression.Expression
	Value   tz.Option[expression.Expression]
}

type SaplingState struct {
	ID   tz.BigInt
	Diff SaplingStateOp
}

func (*SaplingState) LazyStorageDiffKind() string { return "sapling_state" }

type SaplingStateOp interface {
	LazyStorageSaplingStateOp() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SaplingStateOp]{
		Variants: encoding.Variants[SaplingStateOp]{
			0: (*SaplingStateUpdate)(nil),
			1: (*SaplingStateRemove)(nil),
			2: (*SaplingStateCopy)(nil),
			3: (*SaplingStateAlloc)(nil),
		},
	})
}

type SaplingStateUpdate struct {
	Updates SaplingStateUpdates
}

func (*SaplingStateUpdate) LazyStorageSaplingStateOp() string { return "update" }

type SaplingStateRemove struct{}

func (*SaplingStateRemove) LazyStorageSaplingStateOp() string { return "remove" }

type SaplingStateCopy struct {
	Source  tz.BigInt
	Updates SaplingStateUpdates
}

func (*SaplingStateCopy) LazyStorageSaplingStateOp() string { return "copy" }

type SaplingStateAlloc struct {
	Updates  SaplingStateUpdates
	MemoSize uint16
}

func (*SaplingStateAlloc) LazyStorageSaplingStateOp() string { return "alloc" }

type SaplingStateUpdates struct {
	CommitmentsAndCiphertexts []*CommitmentAndCiphertext `tz:"dyn"`
	Nullifiers                []byte                     `tz:"dyn"`
}

type CommitmentAndCiphertext struct {
	Commitment *[32]byte
	Ciphertext SaplingCiphertext
}

type SaplingCiphertext struct {
	Cv         *[32]byte
	Epk        *[32]byte
	PayloadEnc []byte `tz:"dyn"`
	NonceEnc   *[24]byte
	PayloadOut *[80]byte
	NonceOut   *[24]byte
}
