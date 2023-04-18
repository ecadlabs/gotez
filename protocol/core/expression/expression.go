package expression

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type Expression interface {
	Expression()
}

type Int struct {
	Value tz.BigInt
}

func (Int) Expression() {}

type String struct {
	Value string `tz:"dyn"`
}

func (String) Expression() {}

type Seq struct {
	Value []Expression `tz:"dyn"`
}

func (Seq) Expression() {}

type Prim00 = Prim

type Prim0X struct {
	Prim   Prim
	Annots string `tz:"dyn"`
}

func (*Prim0X) Expression() {}

type Prim10 struct {
	Prim Prim
	Arg  Expression
}

func (*Prim10) Expression() {}

type Prim1X struct {
	Prim   Prim
	Arg    Expression
	Annots string `tz:"dyn"`
}

func (*Prim1X) Expression() {}

type Prim20 struct {
	Prim Prim
	Args [2]Expression
}

func (*Prim20) Expression() {}

type Prim2X struct {
	Prim   Prim
	Args   [2]Expression
	Annots string `tz:"dyn"`
}

func (*Prim2X) Expression() {}

type PrimXX struct {
	Prim   Prim
	Args   []Expression `tz:"dyn"`
	Annots string       `tz:"dyn"`
}

func (*PrimXX) Expression() {}

type Bytes struct {
	Value []byte `tz:"dyn"`
}

func (Bytes) Expression() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Expression]{
		Variants: encoding.Variants[Expression]{
			0:  Int{},
			1:  String{},
			2:  Seq{},
			3:  Prim00(0),
			4:  (*Prim0X)(nil),
			5:  (*Prim10)(nil),
			6:  (*Prim1X)(nil),
			7:  (*Prim20)(nil),
			8:  (*Prim2X)(nil),
			9:  (*PrimXX)(nil),
			10: Bytes{},
		},
	})
}
