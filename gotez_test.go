package gotez

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

type args struct {
	data []byte
	v    any
	opt  []DecodeOption
}

type testCase struct {
	name       string
	args       args
	wantRest   []byte
	wantErr    bool
	wantResult any
}

func TestDecode(t *testing.T) {
	type withOmit struct {
		//lint:ignore U1000 test skipping
		priv uint32
		X    uint32
		Omit uint32 `tz:"omit"`
		Y    uint32
	}

	type withDyn struct {
		X uint32
		S []byte `tz:"dyn"`
		Y uint32
	}

	type withOpt struct {
		X *uint32 `tz:"opt"`
	}

	type withOptDyn struct {
		X *uint32 `tz:"opt,dyn"`
		Y uint32
	}

	type withDynOpt struct {
		X *uint32 `tz:"dyn,opt"`
		Y uint32
	}

	type withConst struct {
		X uint32
		C uint32 `tz:"const=0xabcd"`
		Y uint32
	}

	tests := []testCase{
		{
			name: "bool",
			args: args{
				data: []byte{0xff},
				v:    new(bool),
			},
			wantRest:   []byte{},
			wantResult: func() *bool { x := true; return &x }(),
		},
		{
			name: "uint8",
			args: args{
				data: []byte{0xab},
				v:    new(uint8),
			},
			wantRest:   []byte{},
			wantResult: func() *uint8 { x := uint8(0xab); return &x }(),
		},
		{
			name: "int8",
			args: args{
				data: []byte{123},
				v:    new(int8),
			},
			wantRest:   []byte{},
			wantResult: func() *int8 { x := int8(123); return &x }(),
		},
		{
			name: "*uint8",
			args: args{
				data: []byte{0xab},
				v:    func() **uint8 { var x *uint8; return &x }(),
			},
			wantRest:   []byte{},
			wantResult: func() **uint8 { x := uint8(0xab); p := &x; return &p }(),
		},
		{
			name: "uint16",
			args: args{
				data: []byte{0xab, 0xcd},
				v:    new(uint16),
			},
			wantRest:   []byte{},
			wantResult: func() *uint16 { x := uint16(0xabcd); return &x }(),
		},
		{
			name: "int16",
			args: args{
				data: []byte{0xab, 0xcd},
				v:    new(int16),
			},
			wantRest:   []byte{},
			wantResult: func() *int16 { x := int16(-21555); return &x }(),
		},
		{
			name: "*uint16",
			args: args{
				data: []byte{0xab, 0xcd},
				v:    func() **uint16 { var x *uint16; return &x }(),
			},
			wantRest:   []byte{},
			wantResult: func() **uint16 { x := uint16(0xabcd); p := &x; return &p }(),
		},
		{
			name: "uint32",
			args: args{
				data: []byte{0x01, 0x23, 0xab, 0xcd},
				v:    new(uint32),
			},
			wantRest:   []byte{},
			wantResult: func() *uint32 { x := uint32(0x0123abcd); return &x }(),
		},
		{
			name: "int32",
			args: args{
				data: []byte{0x01, 0x23, 0xab, 0xcd},
				v:    new(int32),
			},
			wantRest:   []byte{},
			wantResult: func() *int32 { x := int32(0x0123abcd); return &x }(),
		},
		{
			name: "*uint32",
			args: args{
				data: []byte{0x01, 0x23, 0xab, 0xcd},
				v:    func() **uint32 { var x *uint32; return &x }(),
			},
			wantRest:   []byte{},
			wantResult: func() **uint32 { x := uint32(0x0123abcd); p := &x; return &p }(),
		},
		{
			name: "uint64",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    new(uint64),
			},
			wantRest:   []byte{},
			wantResult: func() *uint64 { x := uint64(0x0123456789abcdef); return &x }(),
		},
		{
			name: "int64",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    new(int64),
			},
			wantRest:   []byte{},
			wantResult: func() *int64 { x := int64(0x0123456789abcdef); return &x }(),
		},
		{
			name: "*uint64",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    func() **uint64 { var x *uint64; return &x }(),
			},
			wantRest:   []byte{},
			wantResult: func() **uint64 { x := uint64(0x0123456789abcdef); p := &x; return &p }(),
		},
		{
			name: "[n]byte",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    new([8]byte),
			},
			wantRest:   []byte{},
			wantResult: &[8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		},
		{
			name: "*[n]byte",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    func() **[8]byte { var x *[8]byte; return &x }(),
			},
			wantRest:   []byte{},
			wantResult: func() **[8]byte { x := &[8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}; return &x }(),
		},
		{
			name: "[]byte",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    new([]byte),
			},
			wantRest:   []byte{},
			wantResult: &[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
		},
		{
			name: "*[]byte",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    func() **[]byte { var x *[]byte; return &x }(),
			},
			wantRest:   []byte{},
			wantResult: func() **[]byte { x := &[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}; return &x }(),
		},
		{
			name: "[]uint32",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    new([]uint32),
			},
			wantRest:   []byte{},
			wantResult: &[]uint32{0x01234567, 0x89abcdef},
		},
		{
			name: "struct",
			args: args{
				data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
				v:    new(withOmit),
			},
			wantRest: []byte{},
			wantResult: &withOmit{
				X: 0x01234567,
				Y: 0x89abcdef,
			},
		},
		{
			name: "dynamic struct",
			args: args{
				data: []byte{
					0x01, 0x23, 0x45, 0x67, // X
					0x00, 0x00, 0x00, 0x08, // array length
					0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, // array
					0x89, 0xab, 0xcd, 0xef, // Y
				},
				v: new(withDyn),
			},
			wantRest: []byte{},
			wantResult: &withDyn{
				X: 0x01234567,
				S: []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77},
				Y: 0x89abcdef,
			},
		},
		{
			name: "option some",
			args: args{
				data: []byte{
					0xff,                   // tag
					0x01, 0x23, 0x45, 0x67, // X
				},
				v: new(Option[uint32]),
			},
			wantRest:   []byte{},
			wantResult: func() *Option[uint32] { v := Some(uint32(0x01234567)); return &v }(),
		},
		{
			name: "option none",
			args: args{
				data: []byte{
					0x00, // tag
				},
				v: new(Option[uint32]),
			},
			wantRest:   []byte{},
			wantResult: func() *Option[uint32] { v := None[uint32](); return &v }(),
		},
		{
			name: "optional attribute some",
			args: args{
				data: []byte{
					0xff,                   // tag
					0x01, 0x23, 0x45, 0x67, // X
				},
				v: new(withOpt),
			},
			wantRest:   []byte{},
			wantResult: func() *withOpt { x := uint32(0x01234567); return &withOpt{X: &x} }(),
		},
		{
			name: "optional attribute none",
			args: args{
				data: []byte{
					0x00, // tag
				},
				v: new(withOpt),
			},
			wantRest:   []byte{},
			wantResult: &withOpt{X: nil},
		},
		{
			name: "optional and dynamic attributes",
			args: args{
				data: []byte{
					0xff,                   // tag
					0x00, 0x00, 0x00, 0x05, // X length
					0x01, 0x23, 0x45, 0x67, // X
					0x00,                   // extra byte
					0x89, 0xab, 0xcd, 0xef, // Y
				},
				v: new(withOptDyn),
			},
			wantRest:   []byte{},
			wantResult: func() *withOptDyn { x := uint32(0x01234567); return &withOptDyn{X: &x, Y: 0x89abcdef} }(),
		},
		{
			name: "dynamic and optional attributes",
			args: args{
				data: []byte{
					0x00, 0x00, 0x00, 0x06, // X length
					0xff,                   // tag
					0x01, 0x23, 0x45, 0x67, // X
					0x00,                   // extra byte
					0x89, 0xab, 0xcd, 0xef, // Y
				},
				v: new(withDynOpt),
			},
			wantRest:   []byte{},
			wantResult: func() *withDynOpt { x := uint32(0x01234567); return &withDynOpt{X: &x, Y: 0x89abcdef} }(),
		},
		{
			name: "const",
			args: args{
				data: []byte{
					0x01, 0x23, 0x45, 0x67, // X
					0x00, 0x00, 0xab, 0xcd, // C
					0x89, 0xab, 0xcd, 0xef, // Y
				},
				v: new(withConst),
			},
			wantRest: []byte{},
			wantResult: &withConst{
				X: 0x01234567,
				C: 0xabcd,
				Y: 0x89abcdef,
			},
		},
		{
			name: "const err",
			args: args{
				data: []byte{
					0x01, 0x23, 0x45, 0x67, // X
					0x00, 0x00, 0xaa, 0xaa, // C
					0x89, 0xab, 0xcd, 0xef, // Y
				},
				v: new(withConst),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRest, err := Decode(tt.args.data, tt.args.v, tt.args.opt...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResult, tt.args.v)
				require.Equal(t, tt.wantRest, gotRest)
			}
		})
	}
}

type testEnum interface {
	testEnum()
}

type variant1 uint32

func (variant1) testEnum() {}

type variant2 struct {
	X uint32
	Y uint32
}

func (*variant2) testEnum() {}

type variant3 uint32

func (variant3) testEnum() {}

type testEnum1 interface {
	testEnum1()
}

type variant11 uint32

func (variant11) testEnum1() {}

func TestEnum(t *testing.T) {
	enums := NewEnumRegistry()
	enums.RegisterEnum(Variants[testEnum]{
		0: variant1(0),
		1: (*variant2)(nil),
	}, variant3(0))
	enums.RegisterEnum(Variants[testEnum1]{
		0: variant11(0),
	}, nil)

	tests := []testCase{
		{
			name: "var1",
			args: args{
				data: []byte{
					0x00,                   // tag
					0x01, 0x23, 0xab, 0xcd, // Var1
				},
				v:   new(testEnum),
				opt: []DecodeOption{Enums(enums)},
			},
			wantRest: []byte{},
			wantResult: func() *testEnum {
				var e testEnum = variant1(0x0123abcd)
				return &e
			}(),
		},
		{
			name: "var2",
			args: args{
				data: []byte{
					0x01,                   // tag
					0x01, 0x23, 0x45, 0x67, // X
					0x89, 0xab, 0xcd, 0xef, // Y
				},
				v:   new(testEnum),
				opt: []DecodeOption{Enums(enums)},
			},
			wantRest: []byte{},
			wantResult: func() *testEnum {
				var e testEnum = &variant2{
					X: 0x01234567,
					Y: 0x89abcdef,
				}
				return &e
			}(),
		},
		{
			name: "var3",
			args: args{
				data: []byte{
					0xff,                   // tag
					0x01, 0x23, 0xab, 0xcd, // Var3
				},
				v:   new(testEnum),
				opt: []DecodeOption{Enums(enums)},
			},
			wantRest: []byte{},
			wantResult: func() *testEnum {
				var e testEnum = variant3(0x0123abcd)
				return &e
			}(),
		},
		{
			name: "unknown tag",
			args: args{
				data: []byte{
					0xff, // tag
				},
				v:   new(testEnum1),
				opt: []DecodeOption{Enums(enums)},
			},
			wantErr:    true,
			wantRest:   nil,
			wantResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRest, err := Decode(tt.args.data, tt.args.v, tt.args.opt...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResult, tt.args.v)
				require.Equal(t, tt.wantRest, gotRest)
			}
		})
	}
}

func TestBigInt(t *testing.T) {
	type args struct {
		data []byte
		v    *BigInt
		opt  []DecodeOption
	}

	type testCase struct {
		name       string
		args       args
		wantRest   []byte
		wantErr    bool
		wantResult *big.Int
	}

	tests := []testCase{
		{
			name: "zero",
			args: args{
				data: []byte{0x00},
				v:    new(BigInt),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(0),
		},
		{
			name: "long zero",
			args: args{
				data: []byte{0x80, 0x00},
				v:    new(BigInt),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(0),
		},
		{
			name: "positive",
			args: args{
				data: []byte{0x95, 0x84, 0xcc, 0xde, 0x8f, 0xbd, 0x88, 0xa2, 0x22},
				v:    new(BigInt),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(1234567890123456789),
		},
		{
			name: "negative",
			args: args{
				data: []byte{0xd5, 0x84, 0xcc, 0xde, 0x8f, 0xbd, 0x88, 0xa2, 0x22},
				v:    new(BigInt),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(-1234567890123456789),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRest, err := Decode(tt.args.data, tt.args.v, tt.args.opt...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResult, tt.args.v.Int())
				require.Equal(t, tt.wantRest, gotRest)
			}
		})
	}
}

func TestBigUint(t *testing.T) {
	type args struct {
		data []byte
		v    *BigUint
		opt  []DecodeOption
	}

	type testCase struct {
		name       string
		args       args
		wantRest   []byte
		wantErr    bool
		wantResult *big.Int
	}

	tests := []testCase{
		{
			name: "zero",
			args: args{
				data: []byte{0x00},
				v:    new(BigUint),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(0),
		},
		{
			name: "long zero",
			args: args{
				data: []byte{0x80, 0x00},
				v:    new(BigUint),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(0),
		},
		{
			name: "one",
			args: args{
				data: []byte{0x01},
				v:    new(BigUint),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(1),
		},
		{
			name: "short",
			args: args{
				data: []byte{0x95, 0x9a, 0xef, 0x3a},
				v:    new(BigUint),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(123456789),
		},
		{
			name: "long",
			args: args{
				data: []byte{0x95, 0x82, 0xa6, 0xef, 0xc7, 0x9e, 0x84, 0x91, 0x11},
				v:    new(BigUint),
			},
			wantRest:   []byte{},
			wantResult: big.NewInt(1234567890123456789),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRest, err := Decode(tt.args.data, tt.args.v, tt.args.opt...)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantResult, tt.args.v.Int())
				require.Equal(t, tt.wantRest, gotRest)
			}
		})
	}
}
