package lazy

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

//go:generate go run ../../../cmd/genmarshaller.go

type StorageDiff struct {
	Contents []DiffKind `tz:"dyn" json:"contents"`
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

//json:kind=big_map
type BigMap struct {
	ID   tz.BigInt `json:"id"`
	Diff BigMapOp  `json:"diff"`
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

//json:action=update
type BigMapUpdate struct {
	Updates []*KeyValue `tz:"dyn" json:"updates"`
}

func (*BigMapUpdate) LazyStorageBigMapOp() string { return "update" }

//json:action=remove
type BigMapRemove struct{}

func (*BigMapRemove) LazyStorageBigMapOp() string { return "remove" }

//json:action=copy
type BigMapCopy struct {
	Source  tz.BigInt   `json:"source"`
	Updates []*KeyValue `tz:"dyn" json:"updates"`
}

func (*BigMapCopy) LazyStorageBigMapOp() string { return "copy" }

//json:action=alloc
type BigMapAlloc struct {
	Updates   []*KeyValue           `tz:"dyn" json:"updates"`
	KeyType   expression.Expression `json:"key_type"`
	ValueType expression.Expression `json:"value_type"`
}

func (*BigMapAlloc) LazyStorageBigMapOp() string { return "alloc" }

type KeyValue struct {
	KeyHash *tz.ScriptExprHash               `json:"key_hash"`
	Key     expression.Expression            `json:"key"`
	Value   tz.Option[expression.Expression] `json:"value"`
}

//json:kind=sapling_state
type SaplingState struct {
	ID   tz.BigInt      `json:"id"`
	Diff SaplingStateOp `json:"diff"`
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

//json:action=update
type SaplingStateUpdate struct {
	Updates SaplingStateUpdates `json:"updates"`
}

func (*SaplingStateUpdate) LazyStorageSaplingStateOp() string { return "update" }

//json:action=remove
type SaplingStateRemove struct{}

func (*SaplingStateRemove) LazyStorageSaplingStateOp() string { return "remove" }

//json:action=copy
type SaplingStateCopy struct {
	Source  tz.BigInt           `json:"source"`
	Updates SaplingStateUpdates `json:"updates"`
}

func (*SaplingStateCopy) LazyStorageSaplingStateOp() string { return "copy" }

//json:action=alloc
type SaplingStateAlloc struct {
	Updates  SaplingStateUpdates `json:"updates"`
	MemoSize uint16              `json:"memo_size"`
}

func (*SaplingStateAlloc) LazyStorageSaplingStateOp() string { return "alloc" }

type SaplingStateUpdates struct {
	CommitmentsAndCiphertexts []*CommitmentAndCiphertext `tz:"dyn" json:"commitments_and_ciphertexts"`
	Nullifiers                tz.Bytes                   `tz:"dyn" json:"nullifiers"`
}

type CommitmentAndCiphertext struct {
	Commitment *[32]byte         `json:"commitment"`
	Ciphertext SaplingCiphertext `json:"ciphertext"`
}

type SaplingCiphertext struct {
	Cv         *[32]byte `json:"cv"`
	Epk        *[32]byte `json:"epk"`
	PayloadEnc []byte    `tz:"dyn" json:"payload_enc"`
	NonceEnc   *[24]byte `json:"nonce_enc"`
	PayloadOut *[80]byte `json:"payload_out"`
	NonceOut   *[24]byte `json:"nonce_out"`
}
