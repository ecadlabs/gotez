package big_map

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
)

//go:generate go run ../../../cmd/genmarshaller.go

type Diff struct {
	Contents []Op `tz:"dyn" json:"contents"`
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

//json:action=BigMapDiffOp()
type Update struct {
	BigMap  tz.BigInt                        `json:"big_map"`
	KeyHash *tz.ScriptExprHash               `json:"key_hash"`
	Key     expression.Expression            `json:"key"`
	Value   tz.Option[expression.Expression] `json:"value"`
}

func (*Update) BigMapDiffOp() string { return "update" }

//json:action=BigMapDiffOp()
type Remove struct {
	BigMap tz.BigInt `json:"big_map"`
}

func (*Remove) BigMapDiffOp() string { return "remove" }

//json:action=BigMapDiffOp()
type Copy struct {
	SourceBigMap      tz.BigInt `json:"source_big_map"`
	DestinationBigMap tz.BigInt `json:"destination_big_map"`
}

func (*Copy) BigMapDiffOp() string { return "copy" }

//json:action=BigMapDiffOp()
type Alloc struct {
	BigMap    tz.BigInt             `json:"big_map"`
	KeyType   expression.Expression `json:"key_type"`
	ValueType expression.Expression `json:"value_type"`
}

func (*Alloc) BigMapDiffOp() string { return "alloc" }
