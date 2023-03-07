package gotez

import (
	"math/big"
	"testing"

	"github.com/ecadlabs/gotez/encoding"
	"github.com/stretchr/testify/require"
)

type args struct {
	data []byte
	v    any
	opt  []encoding.DecodeOption
}

type testCase struct {
	name       string
	args       args
	wantRest   []byte
	wantErr    bool
	wantResult any
}

func TestDecode(t *testing.T) {
	tests := []testCase{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRest, err := encoding.Decode(tt.args.data, tt.args.v, tt.args.opt...)
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
		opt  []encoding.DecodeOption
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
			gotRest, err := encoding.Decode(tt.args.data, tt.args.v, tt.args.opt...)
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
		opt  []encoding.DecodeOption
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
			gotRest, err := encoding.Decode(tt.args.data, tt.args.v, tt.args.opt...)
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
