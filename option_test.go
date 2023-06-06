package gotez

import (
	"encoding/json"
	"testing"

	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/stretchr/testify/require"
)

type args struct {
	data []byte
	v    any
	opt  []encoding.Option
}

type testCase struct {
	name       string
	args       args
	wantRest   []byte
	wantErr    bool
	wantResult any
}

func TestOption(t *testing.T) {
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

func TestOptionJSON(t *testing.T) {
	cases := []struct {
		v      Option[int32]
		expect string
	}{
		{
			v:      Some[int32](123),
			expect: "123",
		},
		{
			v:      None[int32](),
			expect: "null",
		},
	}
	for _, tt := range cases {
		j, err := json.Marshal(&tt.v)
		require.NoError(t, err)
		require.Equal(t, []byte(tt.expect), j)

		var val Option[int32]
		require.NoError(t, json.Unmarshal([]byte(tt.expect), &val))
		require.Equal(t, &tt.v, &val)
	}
}
