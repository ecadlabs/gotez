package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type LazyStorageDiff struct {
	Contents []LazyStorageDiffKind `tz:"dyn"`
}

type LazyStorageDiffKind interface {
	LazyStorageDiffKind() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[LazyStorageDiffKind]{
		Variants: encoding.Variants[LazyStorageDiffKind]{
			0: (*LazyStorageDiffBigMap)(nil),
			1: (*LazyStorageDiffSaplingState)(nil),
		},
	})
}

type LazyStorageDiffBigMap struct {
	ID   tz.BigUint
	Diff BigMapDiff
}

func (*LazyStorageDiffBigMap) LazyStorageDiffKind() string { return "big_map" }

type BigMapDiff interface {
	BigMapDiffKind() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BigMapDiff]{
		Variants: encoding.Variants[BigMapDiff]{
			0: (*BigMapDiffUpdate)(nil),
			1: (*BigMapDiffRemove)(nil),
			2: (*BigMapDiffCopy)(nil),
			3: (*BigMapDiffAlloc)(nil),
		},
	})
}

type BigMapDiffUpdate struct {
	Updates []*BigMapKeyValue `tz:"dyn"`
}

func (*BigMapDiffUpdate) BigMapDiffKind() string { return "update" }

type BigMapDiffRemove struct{}

func (*BigMapDiffRemove) BigMapDiffKind() string { return "remove" }

type BigMapDiffCopy struct {
	Source  tz.BigUint
	Updates []*BigMapKeyValue `tz:"dyn"`
}

func (*BigMapDiffCopy) BigMapDiffKind() string { return "copy" }

type BigMapDiffAlloc struct {
	Updates   []*BigMapKeyValue `tz:"dyn"`
	KeyType   expression.Expression
	ValueType expression.Expression
}

func (*BigMapDiffAlloc) BigMapDiffKind() string { return "alloc" }

type BigMapKeyValue struct {
	KeyHash *tz.ScriptExprHash
	Key     expression.Expression
	Value   tz.Option[expression.Expression]
}

type LazyStorageDiffSaplingState struct {
	ID   tz.BigUint
	Diff SaplingStateDiff
}

func (*LazyStorageDiffSaplingState) LazyStorageDiffKind() string { return "sapling_state" }

type SaplingStateDiff interface {
	SaplingStateDiffKind() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SaplingStateDiff]{
		Variants: encoding.Variants[SaplingStateDiff]{
			0: (*SaplingStateDiffUpdate)(nil),
			1: (*SaplingStateDiffRemove)(nil),
			2: (*SaplingStateDiffCopy)(nil),
			3: (*SaplingStateDiffAlloc)(nil),
		},
	})
}

type SaplingStateDiffUpdate struct {
	Updates SaplingStateDiffUpdateUpdates
}

func (*SaplingStateDiffUpdate) SaplingStateDiffKind() string { return "update" }

type SaplingStateDiffRemove struct{}

func (*SaplingStateDiffRemove) SaplingStateDiffKind() string { return "remove" }

type SaplingStateDiffCopy struct {
	Source  tz.BigUint
	Updates SaplingStateDiffUpdateUpdates
}

func (*SaplingStateDiffCopy) SaplingStateDiffKind() string { return "copy" }

type SaplingStateDiffAlloc struct {
	Updates  SaplingStateDiffUpdateUpdates
	MemoSize uint16
}

func (*SaplingStateDiffAlloc) SaplingStateDiffKind() string { return "alloc" }

type SaplingStateDiffUpdateUpdates struct {
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
