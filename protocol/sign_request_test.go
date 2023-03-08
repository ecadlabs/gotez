package protocol

import (
	"encoding/hex"
	"testing"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/stretchr/testify/require"
)

func TestSignRequest(t *testing.T) {
	type testCase struct {
		title     string
		src       string
		expect    SignRequest
		watermark *Watermark
	}

	testCases := []testCase{
		{
			title: "preendorsement",
			src:   "12ed9d217c2f50673bab6b20dfb0a88ca93b4a0c72a34c807af5dffbece2cba3d2b509835f14006000000002000000041f1ebb39759cc957216f88fb4d005abc206fb00a53f8d57ac01be00c084cba97",
			expect: &PreendorsementRequest{
				Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
				Branch: &tz.BlockHash{
					0x2f, 0x50, 0x67, 0x3b, 0xab, 0x6b, 0x20, 0xdf, 0xb0, 0xa8, 0x8c, 0xa9, 0x3b, 0x4a,
					0xc, 0x72, 0xa3, 0x4c, 0x80, 0x7a, 0xf5, 0xdf, 0xfb, 0xec, 0xe2, 0xcb, 0xa3, 0xd2,
					0xb5, 0x9, 0x83, 0x5f,
				},
				Operation: &Preendorsement{
					Slot:  96,
					Level: 2,
					Round: 4,
					BlockPayloadHash: &tz.BlockPayloadHash{
						0x1f, 0x1e, 0xbb, 0x39, 0x75, 0x9c, 0xc9, 0x57, 0x21, 0x6f, 0x88, 0xfb, 0x4d, 0x0,
						0x5a, 0xbc, 0x20, 0x6f, 0xb0, 0xa, 0x53, 0xf8, 0xd5, 0x7a, 0xc0, 0x1b, 0xe0, 0xc,
						0x8, 0x4c, 0xba, 0x97,
					},
				},
			},
			watermark: &Watermark{
				Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
				Level: Level{
					Level: 2,
					Round: tz.Some(int32(4)),
				},
				Order: WmOrderPreendorsement,
			},
		},
		{
			title: "endorsement",
			src:   "13ed9d217cfc81eee810737b04018acef4db74d056b79edc43e6be46cae7e4c217c22a82f01500120000518d0000000003e7ea1f67dbb0bb6cfa372cb092cd9cf786b4f1b5e5139da95b915fb95e698d",
			expect: &EndorsementRequest{
				Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
				Branch: &tz.BlockHash{
					0xfc, 0x81, 0xee, 0xe8, 0x10, 0x73, 0x7b, 0x4, 0x1, 0x8a, 0xce, 0xf4, 0xdb, 0x74, 0xd0,
					0x56, 0xb7, 0x9e, 0xdc, 0x43, 0xe6, 0xbe, 0x46, 0xca, 0xe7, 0xe4, 0xc2, 0x17, 0xc2,
					0x2a, 0x82, 0xf0,
				},
				Operation: &Endorsement{
					Slot:  18,
					Level: 20877,
					Round: 0,
					BlockPayloadHash: &tz.BlockPayloadHash{
						0x3, 0xe7, 0xea, 0x1f, 0x67, 0xdb, 0xb0, 0xbb, 0x6c, 0xfa, 0x37, 0x2c, 0xb0, 0x92,
						0xcd, 0x9c, 0xf7, 0x86, 0xb4, 0xf1, 0xb5, 0xe5, 0x13, 0x9d, 0xa9, 0x5b, 0x91, 0x5f,
						0xb9, 0x5e, 0x69, 0x8d,
					},
				},
			},
			watermark: &Watermark{
				Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
				Level: Level{
					Level: 20877,
					Round: tz.Some(int32(0)),
				},
				Order: WmOrderEndorsement,
			},
		},
		{
			title: "block",
			src:   "11ed9d217c0000518e0118425847ac255b6d7c30ce8fec23b8eaf13b741de7d18509ac2ef83c741209630000000061947af504805682ea5d089837764b3efcc90b91db24294ff9ddb66019f332ccba17cc4741000000210000000102000000040000518e0000000000000004ffffffff0000000400000000eb1320a71e8bf8b0162a3ec315461e9153a38b70d00d5dde2df85eb92748f8d068d776e356683a9e23c186ccfb72ddc6c9857bb1704487972922e7c89a7121f800000000a8e1dd3c000000000000",
			expect: &TenderbakeBlockRequest{
				Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
				BlockHeader: TenderbakeBlockHeader{
					ShellHeader: ShellHeader{
						Level: 20878,
						Proto: 1,
						Predecessor: &tz.BlockHash{
							0x18, 0x42, 0x58, 0x47, 0xac, 0x25, 0x5b, 0x6d, 0x7c, 0x30, 0xce, 0x8f, 0xec,
							0x23, 0xb8, 0xea, 0xf1, 0x3b, 0x74, 0x1d, 0xe7, 0xd1, 0x85, 0x9, 0xac, 0x2e,
							0xf8, 0x3c, 0x74, 0x12, 0x9, 0x63,
						},
						Timestamp:      1637120757,
						ValidationPass: 4,
						OperationsHash: &tz.OperationsHash{
							0x80, 0x56, 0x82, 0xea, 0x5d, 0x8, 0x98, 0x37, 0x76, 0x4b, 0x3e, 0xfc, 0xc9,
							0xb, 0x91, 0xdb, 0x24, 0x29, 0x4f, 0xf9, 0xdd, 0xb6, 0x60, 0x19, 0xf3, 0x32,
							0xcc, 0xba, 0x17, 0xcc, 0x47, 0x41,
						},
						Fitness: []byte{
							0x0, 0x0, 0x0, 0x1, 0x2, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x51, 0x8e, 0x0, 0x0,
							0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x4, 0x0,
							0x0, 0x0, 0x0,
						},
						Context: &tz.ContextHash{
							0xeb, 0x13, 0x20, 0xa7, 0x1e, 0x8b, 0xf8, 0xb0, 0x16, 0x2a, 0x3e, 0xc3, 0x15,
							0x46, 0x1e, 0x91, 0x53, 0xa3, 0x8b, 0x70, 0xd0, 0xd, 0x5d, 0xde, 0x2d, 0xf8,
							0x5e, 0xb9, 0x27, 0x48, 0xf8, 0xd0,
						},
					},
					PayloadHash: &tz.BlockPayloadHash{
						0x68, 0xd7, 0x76, 0xe3, 0x56, 0x68, 0x3a, 0x9e, 0x23, 0xc1, 0x86, 0xcc, 0xfb, 0x72,
						0xdd, 0xc6, 0xc9, 0x85, 0x7b, 0xb1, 0x70, 0x44, 0x87, 0x97, 0x29, 0x22, 0xe7, 0xc8,
						0x9a, 0x71, 0x21, 0xf8,
					},
					PayloadRound:              0,
					ProofOfWorkNonce:          &[8]byte{0xa8, 0xe1, 0xdd, 0x3c, 0x0, 0x0, 0x0, 0x0},
					SeedNonceHash:             tz.None[*tz.CycleNonceHash](),
					LiquidityBakingToggleVote: 0,
				},
			},
			watermark: &Watermark{
				Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
				Level: Level{
					Level: 20878,
					Round: tz.Some(int32(0)),
				},
				Order: WmOrderDefault,
			},
		},
		{
			title: "operation",
			src:   "03a60703a9567bf69ec66b368c3d8562eba4cbf29278c2c10447a684e3aa1436856c00a0c7a9b0bcd6a48ee0c13094327f215ba2adeaa7d40dabc1af25e36fde02c096b10201f525eabd8b0eeace1494233ea0230d2c9ad6619b00ffff0b66756c66696c6c5f61736b0000000907070088f0f6010306",
			expect: &GenericOperationRequest{
				Branch: &tz.BlockHash{
					0xa6, 0x7, 0x3, 0xa9, 0x56, 0x7b, 0xf6, 0x9e, 0xc6, 0x6b, 0x36, 0x8c, 0x3d, 0x85, 0x62,
					0xeb, 0xa4, 0xcb, 0xf2, 0x92, 0x78, 0xc2, 0xc1, 0x4, 0x47, 0xa6, 0x84, 0xe3, 0xaa,
					0x14, 0x36, 0x85,
				},
				Operations: []OperationContents{
					&Transaction{
						ManagerOperation: ManagerOperation{
							Source: &tz.Ed25519PublicKeyHash{
								0xa0, 0xc7, 0xa9, 0xb0, 0xbc, 0xd6, 0xa4, 0x8e, 0xe0, 0xc1, 0x30, 0x94, 0x32, 0x7f, 0x21,
								0x5b, 0xa2, 0xad, 0xea, 0xa7,
							},
							Fee:          tz.BigUint{0xd4, 0xd},
							Counter:      tz.BigUint{0xab, 0xc1, 0xaf, 0x25},
							GasLimit:     tz.BigUint{0xe3, 0x6f},
							StorageLimit: tz.BigUint{0xde, 0x2},
						},
						Amount: tz.BigUint{0xc0, 0x96, 0xb1, 0x2},
						Destination: &tz.OriginatedContract{
							ContractHash: &tz.ContractHash{
								0xf5, 0x25, 0xea, 0xbd, 0x8b, 0xe, 0xea, 0xce, 0x14, 0x94, 0x23, 0x3e, 0xa0, 0x23, 0xd,
								0x2c, 0x9a, 0xd6, 0x61, 0x9b,
							},
							Padding: 0x0,
						},
						Parameters: tz.Some(Parameters{
							Entrypoint: EpNamed{String: "fulfill_ask"},
							Value:      []byte{0x7, 0x7, 0x0, 0x88, 0xF0, 0xF6, 0x1, 0x3, 0x6},
						}),
					},
				},
			},
		},
		{
			title: "multiple operations",
			src:   "03a60703a9567bf69ec66b368c3d8562eba4cbf29278c2c10447a684e3aa1436856c00a0c7a9b0bcd6a48ee0c13094327f215ba2adeaa7d40dabc1af25e36fde02c096b10201f525eabd8b0eeace1494233ea0230d2c9ad6619b00ffff0b66756c66696c6c5f61736b0000000907070088f0f60103066e00e1ba5449f2938568ace14b5dd54f31936dc86722ba08e0eaa917f53800ff0002298c03ed7d454a101eb7022bc95f7e5f41ac78",
			expect: &GenericOperationRequest{
				Branch: &tz.BlockHash{
					0xa6, 0x7, 0x3, 0xa9, 0x56, 0x7b, 0xf6, 0x9e, 0xc6, 0x6b, 0x36, 0x8c, 0x3d, 0x85, 0x62,
					0xeb, 0xa4, 0xcb, 0xf2, 0x92, 0x78, 0xc2, 0xc1, 0x4, 0x47, 0xa6, 0x84, 0xe3, 0xaa,
					0x14, 0x36, 0x85,
				},
				Operations: []OperationContents{
					&Transaction{
						ManagerOperation: ManagerOperation{
							Source: &tz.Ed25519PublicKeyHash{
								0xa0, 0xc7, 0xa9, 0xb0, 0xbc, 0xd6, 0xa4, 0x8e, 0xe0, 0xc1, 0x30, 0x94, 0x32, 0x7f, 0x21,
								0x5b, 0xa2, 0xad, 0xea, 0xa7,
							},
							Fee:          tz.BigUint{0xd4, 0xd},
							Counter:      tz.BigUint{0xab, 0xc1, 0xaf, 0x25},
							GasLimit:     tz.BigUint{0xe3, 0x6f},
							StorageLimit: tz.BigUint{0xde, 0x2},
						},
						Amount: tz.BigUint{0xc0, 0x96, 0xb1, 0x2},
						Destination: &tz.OriginatedContract{
							ContractHash: &tz.ContractHash{
								0xf5, 0x25, 0xea, 0xbd, 0x8b, 0xe, 0xea, 0xce, 0x14, 0x94, 0x23, 0x3e, 0xa0, 0x23, 0xd,
								0x2c, 0x9a, 0xd6, 0x61, 0x9b,
							},
							Padding: 0x0,
						},
						Parameters: tz.Some(Parameters{
							Entrypoint: EpNamed{String: "fulfill_ask"},
							Value:      []byte{0x7, 0x7, 0x0, 0x88, 0xF0, 0xF6, 0x1, 0x3, 0x6},
						}),
					},
					&Delegation{
						ManagerOperation: ManagerOperation{
							Source:       &tz.Ed25519PublicKeyHash{0xe1, 0xba, 0x54, 0x49, 0xf2, 0x93, 0x85, 0x68, 0xac, 0xe1, 0x4b, 0x5d, 0xd5, 0x4f, 0x31, 0x93, 0x6d, 0xc8, 0x67, 0x22},
							Fee:          tz.BigUint{0xba, 0x8},
							Counter:      tz.BigUint{0xe0, 0xea, 0xa9, 0x17},
							GasLimit:     tz.BigUint{0xf5, 0x38},
							StorageLimit: tz.BigUint{0x0},
						},
						Delegate: tz.Some[tz.PublicKeyHash](&tz.Ed25519PublicKeyHash{0x2, 0x29, 0x8c, 0x3, 0xed, 0x7d, 0x45, 0x4a, 0x10, 0x1e, 0xb7, 0x2, 0x2b, 0xc9, 0x5f, 0x7e, 0x5f, 0x41, 0xac, 0x78}),
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		test := tc
		t.Run(test.title, func(t *testing.T) {
			buf, err := hex.DecodeString(test.src)
			require.NoError(t, err)
			var req SignRequest
			_, err = encoding.Decode(buf, &req)
			require.NoError(t, err)
			require.Equal(t, test.expect, req)
			if test.watermark != nil {
				require.Equal(t, test.watermark, req.(WithWatermark).Watermark())
			}
		})
	}
}

func TestWatermark(t *testing.T) {
	type expect struct {
		wm     Watermark
		expect bool
	}

	type testCase struct {
		stored Watermark
		expect []expect
	}

	testCases := []testCase{
		{
			stored: Watermark{
				Chain: &tz.ChainID{},
				Level: Level{
					Level: 1,
					Round: tz.Some(int32(1)),
				},
				Order: WmOrderDefault,
			},
			expect: []expect{
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(0)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // level above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.None[int32](),
						},
						Order: WmOrderDefault,
					},
					expect: false, // round is set above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // level and round above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // round above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 0,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: false, // level below
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(1)),
						},
						Order: WmOrderDefault,
					},
					expect: false, // level and round below
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(1)),
						},
						Order: WmOrderEndorsement,
					},
					expect: true, // don't have endorsement
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{0xed, 0x9d, 0x21, 0x7c},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(1)),
						},
						Order: WmOrderEndorsement,
					},
					expect: false, // wrong chain
				},
			},
		},
		{
			stored: Watermark{
				Chain: &tz.ChainID{},
				Level: Level{
					Level: 1,
					Round: tz.Some(int32(1)),
				},
				Order: WmOrderEndorsement,
			},
			expect: []expect{
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(0)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // level above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(0)),
						},
						Order: WmOrderEndorsement,
					},
					expect: true, // level above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // level and round above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderEndorsement,
					},
					expect: true, // level and round above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // round above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderEndorsement,
					},
					expect: true, // order above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 0,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: false, // level below
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 0,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderEndorsement,
					},
					expect: false, // level below
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(1)),
						},
						Order: WmOrderDefault,
					},
					expect: false, // level and round below
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 1,
							Round: tz.Some(int32(1)),
						},
						Order: WmOrderEndorsement,
					},
					expect: false, // have endorsement
				},
			},
		},
		{
			stored: Watermark{
				Chain: &tz.ChainID{},
				Level: Level{
					Level: 1,
					Round: tz.None[int32](),
				},
				Order: WmOrderDefault,
			},
			expect: []expect{
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.None[int32](),
						},
						Order: WmOrderDefault,
					},
					expect: true, // level above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 2,
							Round: tz.Some(int32(2)),
						},
						Order: WmOrderDefault,
					},
					expect: true, // level above
				},
				{
					wm: Watermark{
						Chain: &tz.ChainID{},
						Level: Level{
							Level: 0,
							Round: tz.None[int32](),
						},
						Order: WmOrderDefault,
					},
					expect: false, // level below
				},
			},
		},
	}

	for _, c := range testCases {
		for _, ex := range c.expect {
			require.Equal(t, ex.expect, ex.wm.Validate(&c.stored))
		}
	}
}
