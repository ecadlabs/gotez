package big_map

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type Diff struct {
	Contents []Op `tz:"dyn"`
}

type Op interface {
	BigMapDiffOp() string
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Op]{
		Variants: encoding.Variants[Op]{
			0: (*Update)(nil),
			1: (*Remove)(nil),
			2: (*Copy)(nil),
			3: (*Alloc)(nil),
		},
	})
}

type Update struct {
	BigMap  tz.BigInt
	KeyHash *tz.ScriptExprHash
	Key     expression.Expression
	Value   tz.Option[expression.Expression]
}

func (*Update) BigMapDiffOp() string { return "update" }

type Remove struct {
	BigMap tz.BigInt
}

func (*Remove) BigMapDiffOp() string { return "remove" }

type Copy struct {
	SourceBigMap      tz.BigInt
	DestinationBigMap tz.BigInt
}

func (*Copy) BigMapDiffOp() string { return "copy" }

type Alloc struct {
	BigMap    tz.BigInt
	KeyType   expression.Expression
	ValueType expression.Expression
}

func (*Alloc) BigMapDiffOp() string { return "alloc" }
