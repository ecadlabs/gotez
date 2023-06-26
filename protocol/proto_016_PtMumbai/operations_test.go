package proto_016_PtMumbai

import (
	"bytes"
	"encoding/hex"
	"testing"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	expr "github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/stretchr/testify/require"
)

func TestDecodeOperations(t *testing.T) {
	type testCase struct {
		title      string
		kind       string
		src        string
		expect     OperationContents
		skipEncode bool
	}

	testCases := []testCase{
		{
			title: "activate_account",
			src:   "04c55cf02dbeecc978d9c84625dcae72bb77ea4fbd41f98b15efc63fa893d61d7d6eee4a2ce9427ac4",
			kind:  "activate_account",
			expect: &ActivateAccount{
				PKH:    &tz.Ed25519PublicKeyHash{0xc5, 0x5c, 0xf0, 0x2d, 0xbe, 0xec, 0xc9, 0x78, 0xd9, 0xc8, 0x46, 0x25, 0xdc, 0xae, 0x72, 0xbb, 0x77, 0xea, 0x4f, 0xbd},
				Secret: &tz.Bytes20{0x41, 0xf9, 0x8b, 0x15, 0xef, 0xc6, 0x3f, 0xa8, 0x93, 0xd6, 0x1d, 0x7d, 0x6e, 0xee, 0x4a, 0x2c, 0xe9, 0x42, 0x7a, 0xc4},
			},
		},
		{
			title: "seed_nonce_revelation",
			src:   "01000004d2000102030405060708090a0b0c0d0e0f000102030405060708090a0b0c0d0e0f",
			kind:  "seed_nonce_revelation",
			expect: &SeedNonceRevelation{
				Level: 1234,
				Nonce: &tz.SeedNonce{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d,
					0x0e, 0x0f, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b,
					0x0c, 0x0d, 0x0e, 0x0f,
				},
			},
		},
		{
			title: "ballot",
			src:   "060002298c03ed7d454a101eb7022bc95f7e5f41ac78000002cf7663cf120f3dc8189d5dc7d4d7a0483bcc53f3f18e700f5a2f5076aa8b9dc55c00",
			kind:  "ballot",
			expect: &Ballot{
				Source: &tz.Ed25519PublicKeyHash{0x2, 0x29, 0x8c, 0x3, 0xed, 0x7d, 0x45, 0x4a, 0x10, 0x1e, 0xb7, 0x2, 0x2b, 0xc9, 0x5f, 0x7e, 0x5f, 0x41, 0xac, 0x78},
				Period: 719,
				Proposal: &tz.ProtocolHash{
					0x76, 0x63, 0xcf, 0x12, 0xf, 0x3d, 0xc8, 0x18, 0x9d, 0x5d, 0xc7, 0xd4, 0xd7, 0xa0, 0x48,
					0x3b, 0xcc, 0x53, 0xf3, 0xf1, 0x8e, 0x70, 0xf, 0x5a, 0x2f, 0x50, 0x76, 0xaa, 0x8b, 0x9d,
					0xc5, 0x5c,
				},
				Ballot: core.BallotYay,
			},
		},
		{
			title: "delegation",
			src:   "6e00e1ba5449f2938568ace14b5dd54f31936dc86722ba08e0eaa917f53800ff0002298c03ed7d454a101eb7022bc95f7e5f41ac78",
			kind:  "delegation",
			expect: &Delegation{
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
		{
			title: "proposals",
			src:   "05009425086a67fde5a00facd42fba4f2d2763a3d5e50000005200000020d57ed88be5a69815e39386a33f7dcad391f5f507e03b376e499272c86c6cf2a7",
			kind:  "proposals",
			expect: &Proposals{
				Source: &tz.Ed25519PublicKeyHash{0x94, 0x25, 0x8, 0x6a, 0x67, 0xfd, 0xe5, 0xa0, 0xf, 0xac, 0xd4, 0x2f, 0xba, 0x4f, 0x2d, 0x27, 0x63, 0xa3, 0xd5, 0xe5},
				Period: 82,
				Proposals: []*tz.ProtocolHash{{
					0xd5, 0x7e, 0xd8, 0x8b, 0xe5, 0xa6, 0x98, 0x15, 0xe3, 0x93, 0x86, 0xa3, 0x3f, 0x7d, 0xca,
					0xd3, 0x91, 0xf5, 0xf5, 0x7, 0xe0, 0x3b, 0x37, 0x6e, 0x49, 0x92, 0x72, 0xc8, 0x6c, 0x6c,
					0xf2, 0xa7,
				}},
			},
		},
		{
			title: "endorsement",
			src:   "1519a5002792c30000000015b4d26b90b5a56f1333bb2b8f1fce2f121474c7a1e088235a9a24e2bfda5bdd",
			kind:  "endorsement",
			expect: &Endorsement{
				Slot:  6565,
				Level: 2593475,
				Round: 0,
				BlockPayloadHash: &tz.BlockPayloadHash{
					0x15, 0xb4, 0xd2, 0x6b, 0x90, 0xb5, 0xa5, 0x6f, 0x13, 0x33, 0xbb, 0x2b, 0x8f, 0x1f,
					0xce, 0x2f, 0x12, 0x14, 0x74, 0xc7, 0xa1, 0xe0, 0x88, 0x23, 0x5a, 0x9a, 0x24, 0xe2,
					0xbf, 0xda, 0x5b, 0xdd,
				},
			},
		},
		{
			title: "preendorsement",
			src:   "1419e60027926000000000c02203b1c970f9894c2d555e87e7b12c86e343fc0127c2da846fd8ee6dcc5a4c",
			kind:  "preendorsement",
			expect: &Preendorsement{
				Slot:  6630,
				Level: 2593376,
				Round: 0,
				BlockPayloadHash: &tz.BlockPayloadHash{
					0xc0, 0x22, 0x03, 0xb1, 0xc9, 0x70, 0xf9, 0x89, 0x4c, 0x2d, 0x55, 0x5e, 0x87, 0xe7,
					0xb1, 0x2c, 0x86, 0xe3, 0x43, 0xfc, 0x01, 0x27, 0xc2, 0xda, 0x84, 0x6f, 0xd8, 0xee,
					0x6d, 0xcc, 0x5a, 0x4c,
				},
			},
		},
		{
			title: "reveal",
			src:   "6b00150da49bcb649acc43f2434549433d90f3d1f7f1e702cb960be8070000f1ccb8d80895b314a26739b5713b8934e34f96f8f3c9137de5c8edeb1f925da1",
			kind:  "reveal",
			expect: &Reveal{
				ManagerOperation: ManagerOperation{
					Source: &tz.Ed25519PublicKeyHash{
						0x15, 0xd, 0xa4, 0x9b, 0xcb, 0x64, 0x9a, 0xcc, 0x43, 0xf2, 0x43, 0x45, 0x49, 0x43, 0x3d,
						0x90, 0xf3, 0xd1, 0xf7, 0xf1,
					},
					Fee:          tz.BigUint{0xe7, 0x2},
					Counter:      tz.BigUint{0xcb, 0x96, 0xb},
					GasLimit:     tz.BigUint{0xe8, 0x7},
					StorageLimit: tz.BigUint{0x0},
				},
				PublicKey: &tz.Ed25519PublicKey{
					0xf1, 0xcc, 0xb8, 0xd8, 0x8, 0x95, 0xb3, 0x14, 0xa2, 0x67, 0x39, 0xb5, 0x71, 0x3b, 0x89,
					0x34, 0xe3, 0x4f, 0x96, 0xf8, 0xf3, 0xc9, 0x13, 0x7d, 0xe5, 0xc8, 0xed, 0xeb, 0x1f, 0x92,
					0x5d, 0xa1,
				},
			},
		},
		{
			title: "double_baking_evidence",
			src:   "030000010100335adf10135116e45ce4fb0efe9e2265144342f655aae0e3692c55e0cd1f182f7abab41b00000000643af41c04d0f701621067ac4fdc6a9f28f548c942ae2fcbea2ad02fc47ff1e963f64400d50000002100000001020000000400335adf0000000000000004ffffffff00000004000000005f3812e527055f6056d39de55665a1083b810cdfb4a049b1e8da79c3c706fdfd9fbeca0acac785d87bb7d177553db3043aff6759beb9db3b6619fc129162ac75000000005977cc15eb1100000002dafcfb71f587519ad20fdff8de4ec3d6aa2bfad815ac12961e4a4b0736f00193da09e9d7b0a1f87837ccd187a2606467385b26b854b4375c7937dc49a22f53f20000010100335adf10135116e45ce4fb0efe9e2265144342f655aae0e3692c55e0cd1f182f7abab41b00000000643af41c04d0f701621067ac4fdc6a9f28f548c942ae2fcbea2ad02fc47ff1e963f64400d50000002100000001020000000400335adf0000000000000004ffffffff00000004000000005f3812e527055f6056d39de55665a1083b810cdfb4a049b1e8da79c3c706fdfd9fbeca0acac785d87bb7d177553db3043aff6759beb9db3b6619fc129162ac75000000005977cc15eb1100000002dafcfb71f587519ad20fdff8de4ec3d6aa2bfad815ac12961e4a4b0736f00193da09e9d7b0a1f87837ccd187a2606467385b26b854b4375c7937dc49a22f53f2",
			kind:  "double_baking_evidence",
			expect: &DoubleBakingEvidence{
				Block1: BlockHeader{
					UnsignedBlockHeader: UnsignedBlockHeader{
						ShellHeader: core.ShellHeader{
							Level:          3365599,
							Proto:          0x10,
							Predecessor:    &tz.BlockHash{0x13, 0x51, 0x16, 0xe4, 0x5c, 0xe4, 0xfb, 0xe, 0xfe, 0x9e, 0x22, 0x65, 0x14, 0x43, 0x42, 0xf6, 0x55, 0xaa, 0xe0, 0xe3, 0x69, 0x2c, 0x55, 0xe0, 0xcd, 0x1f, 0x18, 0x2f, 0x7a, 0xba, 0xb4, 0x1b},
							Timestamp:      1681585180,
							ValidationPass: 0x4,
							OperationsHash: &tz.OperationsHash{0xd0, 0xf7, 0x1, 0x62, 0x10, 0x67, 0xac, 0x4f, 0xdc, 0x6a, 0x9f, 0x28, 0xf5, 0x48, 0xc9, 0x42, 0xae, 0x2f, 0xcb, 0xea, 0x2a, 0xd0, 0x2f, 0xc4, 0x7f, 0xf1, 0xe9, 0x63, 0xf6, 0x44, 0x0, 0xd5},
							Fitness:        []byte{0x0, 0x0, 0x0, 0x1, 0x2, 0x0, 0x0, 0x0, 0x4, 0x0, 0x33, 0x5a, 0xdf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0},
							Context:        &tz.ContextHash{0x5f, 0x38, 0x12, 0xe5, 0x27, 0x5, 0x5f, 0x60, 0x56, 0xd3, 0x9d, 0xe5, 0x56, 0x65, 0xa1, 0x8, 0x3b, 0x81, 0xc, 0xdf, 0xb4, 0xa0, 0x49, 0xb1, 0xe8, 0xda, 0x79, 0xc3, 0xc7, 0x6, 0xfd, 0xfd},
						},
						UnsignedProtocolBlockHeader: UnsignedProtocolBlockHeader{
							PayloadHash:               &tz.BlockPayloadHash{0x9f, 0xbe, 0xca, 0xa, 0xca, 0xc7, 0x85, 0xd8, 0x7b, 0xb7, 0xd1, 0x77, 0x55, 0x3d, 0xb3, 0x4, 0x3a, 0xff, 0x67, 0x59, 0xbe, 0xb9, 0xdb, 0x3b, 0x66, 0x19, 0xfc, 0x12, 0x91, 0x62, 0xac, 0x75},
							PayloadRound:              0,
							ProofOfWorkNonce:          &tz.Bytes8{0x59, 0x77, 0xcc, 0x15, 0xeb, 0x11, 0x0, 0x0},
							SeedNonceHash:             tz.None[*tz.CycleNonceHash](),
							LiquidityBakingToggleVote: core.LiquidityBakingPass,
						},
					},
					Signature: tz.AnySignature{0xda, 0xfc, 0xfb, 0x71, 0xf5, 0x87, 0x51, 0x9a, 0xd2, 0xf, 0xdf, 0xf8, 0xde, 0x4e, 0xc3, 0xd6, 0xaa, 0x2b, 0xfa, 0xd8, 0x15, 0xac, 0x12, 0x96, 0x1e, 0x4a, 0x4b, 0x7, 0x36, 0xf0, 0x1, 0x93, 0xda, 0x9, 0xe9, 0xd7, 0xb0, 0xa1, 0xf8, 0x78, 0x37, 0xcc, 0xd1, 0x87, 0xa2, 0x60, 0x64, 0x67, 0x38, 0x5b, 0x26, 0xb8, 0x54, 0xb4, 0x37, 0x5c, 0x79, 0x37, 0xdc, 0x49, 0xa2, 0x2f, 0x53, 0xf2},
				},
				Block2: BlockHeader{
					UnsignedBlockHeader: UnsignedBlockHeader{
						ShellHeader: core.ShellHeader{
							Level:          3365599,
							Proto:          0x10,
							Predecessor:    &tz.BlockHash{0x13, 0x51, 0x16, 0xe4, 0x5c, 0xe4, 0xfb, 0xe, 0xfe, 0x9e, 0x22, 0x65, 0x14, 0x43, 0x42, 0xf6, 0x55, 0xaa, 0xe0, 0xe3, 0x69, 0x2c, 0x55, 0xe0, 0xcd, 0x1f, 0x18, 0x2f, 0x7a, 0xba, 0xb4, 0x1b},
							Timestamp:      1681585180,
							ValidationPass: 0x4,
							OperationsHash: &tz.OperationsHash{0xd0, 0xf7, 0x1, 0x62, 0x10, 0x67, 0xac, 0x4f, 0xdc, 0x6a, 0x9f, 0x28, 0xf5, 0x48, 0xc9, 0x42, 0xae, 0x2f, 0xcb, 0xea, 0x2a, 0xd0, 0x2f, 0xc4, 0x7f, 0xf1, 0xe9, 0x63, 0xf6, 0x44, 0x0, 0xd5},
							Fitness:        []byte{0x0, 0x0, 0x0, 0x1, 0x2, 0x0, 0x0, 0x0, 0x4, 0x0, 0x33, 0x5a, 0xdf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0},
							Context:        &tz.ContextHash{0x5f, 0x38, 0x12, 0xe5, 0x27, 0x5, 0x5f, 0x60, 0x56, 0xd3, 0x9d, 0xe5, 0x56, 0x65, 0xa1, 0x8, 0x3b, 0x81, 0xc, 0xdf, 0xb4, 0xa0, 0x49, 0xb1, 0xe8, 0xda, 0x79, 0xc3, 0xc7, 0x6, 0xfd, 0xfd},
						},
						UnsignedProtocolBlockHeader: UnsignedProtocolBlockHeader{
							PayloadHash:               &tz.BlockPayloadHash{0x9f, 0xbe, 0xca, 0xa, 0xca, 0xc7, 0x85, 0xd8, 0x7b, 0xb7, 0xd1, 0x77, 0x55, 0x3d, 0xb3, 0x4, 0x3a, 0xff, 0x67, 0x59, 0xbe, 0xb9, 0xdb, 0x3b, 0x66, 0x19, 0xfc, 0x12, 0x91, 0x62, 0xac, 0x75},
							PayloadRound:              0,
							ProofOfWorkNonce:          &tz.Bytes8{0x59, 0x77, 0xcc, 0x15, 0xeb, 0x11, 0x0, 0x0},
							SeedNonceHash:             tz.None[*tz.CycleNonceHash](),
							LiquidityBakingToggleVote: core.LiquidityBakingPass,
						},
					},
					Signature: tz.AnySignature{0xda, 0xfc, 0xfb, 0x71, 0xf5, 0x87, 0x51, 0x9a, 0xd2, 0xf, 0xdf, 0xf8, 0xde, 0x4e, 0xc3, 0xd6, 0xaa, 0x2b, 0xfa, 0xd8, 0x15, 0xac, 0x12, 0x96, 0x1e, 0x4a, 0x4b, 0x7, 0x36, 0xf0, 0x1, 0x93, 0xda, 0x9, 0xe9, 0xd7, 0xb0, 0xa1, 0xf8, 0x78, 0x37, 0xcc, 0xd1, 0x87, 0xa2, 0x60, 0x64, 0x67, 0x38, 0x5b, 0x26, 0xb8, 0x54, 0xb4, 0x37, 0x5c, 0x79, 0x37, 0xdc, 0x49, 0xa2, 0x2f, 0x53, 0xf2},
				},
			},
			skipEncode: true, // header decoding is incomplete
		},
		{
			title: "double_preendorsement_evidence",
			src:   "070000008b8b87b048db84d61c6d8ceaf13538ccf2bbaf2017fb5d804b77aa58ebe088520d1400000004a2e80000000009a56da0405f15df6064b4d704eb9fe6fdaf885a513ed7ba189eb5321d97386f9734c876e5b4df19fc457c6cb308bcad79ca806ad4950f4e7c3118703ecbdd67c77d63c7040fb923c78ee86300b3040bb3d2c6865e69253c57674161cfe261690000008b8b87b048db84d61c6d8ceaf13538ccf2bbaf2017fb5d804b77aa58ebe088520d1400000004a2e800000000ca774e93ab507f6781b8c0312895cfcbb5f5b7df9c261602c7ec71f9e531ae972e36bcd26702165dec701fe072bd07d0bd776da1e9658d0b62806ce17daa06e593eec7d781f9ab4f6b3e8d8c8531072c6f262187144cf2c97c3aa8710ee2304d",
			kind:  "double_preendorsement_evidence",
			expect: &DoublePreendorsementEvidence{
				Op1: InlinedPreendorsement{
					Branch: &tz.BlockHash{
						0x8b, 0x87, 0xb0, 0x48, 0xdb, 0x84, 0xd6, 0x1c, 0x6d, 0x8c, 0xea, 0xf1, 0x35, 0x38,
						0xcc, 0xf2, 0xbb, 0xaf, 0x20, 0x17, 0xfb, 0x5d, 0x80, 0x4b, 0x77, 0xaa, 0x58, 0xeb,
						0xe0, 0x88, 0x52, 0xd,
					},
					Signature: &tz.GenericSignature{
						0x97, 0x34, 0xc8, 0x76, 0xe5, 0xb4, 0xdf, 0x19, 0xfc, 0x45, 0x7c, 0x6c, 0xb3, 0x8,
						0xbc, 0xad, 0x79, 0xca, 0x80, 0x6a, 0xd4, 0x95, 0xf, 0x4e, 0x7c, 0x31, 0x18, 0x70,
						0x3e, 0xcb, 0xdd, 0x67, 0xc7, 0x7d, 0x63, 0xc7, 0x4, 0xf, 0xb9, 0x23, 0xc7, 0x8e,
						0xe8, 0x63, 0x0, 0xb3, 0x4, 0xb, 0xb3, 0xd2, 0xc6, 0x86, 0x5e, 0x69, 0x25, 0x3c,
						0x57, 0x67, 0x41, 0x61, 0xcf, 0xe2, 0x61, 0x69,
					},
					Contents: &Preendorsement{
						Slot:  0,
						Level: 303848,
						Round: 0,
						BlockPayloadHash: &tz.BlockPayloadHash{
							0x9, 0xa5, 0x6d, 0xa0, 0x40, 0x5f, 0x15, 0xdf, 0x60, 0x64, 0xb4, 0xd7, 0x4,
							0xeb, 0x9f, 0xe6, 0xfd, 0xaf, 0x88, 0x5a, 0x51, 0x3e, 0xd7, 0xba, 0x18, 0x9e,
							0xb5, 0x32, 0x1d, 0x97, 0x38, 0x6f,
						},
					},
				},
				Op2: InlinedPreendorsement{
					Branch: &tz.BlockHash{
						0x8b, 0x87, 0xb0, 0x48, 0xdb, 0x84, 0xd6, 0x1c, 0x6d, 0x8c, 0xea, 0xf1, 0x35, 0x38,
						0xcc, 0xf2, 0xbb, 0xaf, 0x20, 0x17, 0xfb, 0x5d, 0x80, 0x4b, 0x77, 0xaa, 0x58, 0xeb,
						0xe0, 0x88, 0x52, 0xd,
					},
					Signature: &tz.GenericSignature{
						0x2e, 0x36, 0xbc, 0xd2, 0x67, 0x2, 0x16, 0x5d, 0xec, 0x70, 0x1f, 0xe0, 0x72, 0xbd,
						0x7, 0xd0, 0xbd, 0x77, 0x6d, 0xa1, 0xe9, 0x65, 0x8d, 0xb, 0x62, 0x80, 0x6c, 0xe1,
						0x7d, 0xaa, 0x6, 0xe5, 0x93, 0xee, 0xc7, 0xd7, 0x81, 0xf9, 0xab, 0x4f, 0x6b, 0x3e,
						0x8d, 0x8c, 0x85, 0x31, 0x7, 0x2c, 0x6f, 0x26, 0x21, 0x87, 0x14, 0x4c, 0xf2, 0xc9,
						0x7c, 0x3a, 0xa8, 0x71, 0xe, 0xe2, 0x30, 0x4d,
					},
					Contents: &Preendorsement{
						Slot:  0,
						Level: 303848,
						Round: 0,
						BlockPayloadHash: &tz.BlockPayloadHash{
							0xca, 0x77, 0x4e, 0x93, 0xab, 0x50, 0x7f, 0x67, 0x81, 0xb8, 0xc0, 0x31, 0x28,
							0x95, 0xcf, 0xcb, 0xb5, 0xf5, 0xb7, 0xdf, 0x9c, 0x26, 0x16, 0x2, 0xc7, 0xec,
							0x71, 0xf9, 0xe5, 0x31, 0xae, 0x97,
						},
					},
				},
			},
		},
		{
			title: "double_endorsement_evidence",
			src:   "020000008ba60703a9567bf69ec66b368c3d8562eba4cbf29278c2c10447a684e3aa1436851500010000007b0000007bca774e93ab507f6781b8c0312895cfcbb5f5b7df9c261602c7ec71f9e531ae97d3a9e1467b32104921d4e2dd93265739c1a5faee7a7f8880842b096c0b6714200c43fd5872f82581dfe1cb3a76ccdadaa4d6361d72b4abee6884cb7ed87f0b040000008b6280d069cca0c2c8c97c172cc0530e3861cf8050d80970866a388c19bcbbf15f1500010000007b0000007bca774e93ab507f6781b8c0312895cfcbb5f5b7df9c261602c7ec71f9e531ae970ef3e51b218d04c29211b89f5b7582a7169b4810e6dbe46732b44c84331ae6cb32ced7c53ef55e7a2358ed66dedcb98daff1d8ec4f0638f74f215083526d2e03",
			kind:  "double_endorsement_evidence",
			expect: &DoubleEndorsementEvidence{
				Op1: InlinedEndorsement{
					Branch: &tz.BlockHash{
						0xa6, 0x7, 0x3, 0xa9, 0x56, 0x7b, 0xf6, 0x9e, 0xc6, 0x6b, 0x36, 0x8c, 0x3d, 0x85,
						0x62, 0xeb, 0xa4, 0xcb, 0xf2, 0x92, 0x78, 0xc2, 0xc1, 0x4, 0x47, 0xa6, 0x84, 0xe3,
						0xaa, 0x14, 0x36, 0x85,
					},
					Signature: tz.AnySignature{
						0xd3, 0xa9, 0xe1, 0x46, 0x7b, 0x32, 0x10, 0x49, 0x21, 0xd4, 0xe2, 0xdd, 0x93, 0x26,
						0x57, 0x39, 0xc1, 0xa5, 0xfa, 0xee, 0x7a, 0x7f, 0x88, 0x80, 0x84, 0x2b, 0x9, 0x6c,
						0xb, 0x67, 0x14, 0x20, 0xc, 0x43, 0xfd, 0x58, 0x72, 0xf8, 0x25, 0x81, 0xdf, 0xe1,
						0xcb, 0x3a, 0x76, 0xcc, 0xda, 0xda, 0xa4, 0xd6, 0x36, 0x1d, 0x72, 0xb4, 0xab, 0xee,
						0x68, 0x84, 0xcb, 0x7e, 0xd8, 0x7f, 0xb, 0x4,
					},
					Contents: &Endorsement{
						Slot:  1,
						Level: 123,
						Round: 123,
						BlockPayloadHash: &tz.BlockPayloadHash{
							0xca, 0x77, 0x4e, 0x93, 0xab, 0x50, 0x7f, 0x67, 0x81, 0xb8, 0xc0, 0x31, 0x28,
							0x95, 0xcf, 0xcb, 0xb5, 0xf5, 0xb7, 0xdf, 0x9c, 0x26, 0x16, 0x2, 0xc7, 0xec,
							0x71, 0xf9, 0xe5, 0x31, 0xae, 0x97,
						},
					},
				},
				Op2: InlinedEndorsement{
					Branch: &tz.BlockHash{
						0x62, 0x80, 0xd0, 0x69, 0xcc, 0xa0, 0xc2, 0xc8, 0xc9, 0x7c, 0x17, 0x2c, 0xc0, 0x53,
						0xe, 0x38, 0x61, 0xcf, 0x80, 0x50, 0xd8, 0x9, 0x70, 0x86, 0x6a, 0x38, 0x8c, 0x19,
						0xbc, 0xbb, 0xf1, 0x5f,
					},
					Signature: tz.AnySignature{
						0xe, 0xf3, 0xe5, 0x1b, 0x21, 0x8d, 0x4, 0xc2, 0x92, 0x11, 0xb8, 0x9f, 0x5b, 0x75,
						0x82, 0xa7, 0x16, 0x9b, 0x48, 0x10, 0xe6, 0xdb, 0xe4, 0x67, 0x32, 0xb4, 0x4c, 0x84,
						0x33, 0x1a, 0xe6, 0xcb, 0x32, 0xce, 0xd7, 0xc5, 0x3e, 0xf5, 0x5e, 0x7a, 0x23, 0x58,
						0xed, 0x66, 0xde, 0xdc, 0xb9, 0x8d, 0xaf, 0xf1, 0xd8, 0xec, 0x4f, 0x6, 0x38, 0xf7,
						0x4f, 0x21, 0x50, 0x83, 0x52, 0x6d, 0x2e, 0x3,
					},
					Contents: &Endorsement{
						Slot:  1,
						Level: 123,
						Round: 123,
						BlockPayloadHash: &tz.BlockPayloadHash{
							0xca, 0x77, 0x4e, 0x93, 0xab, 0x50, 0x7f, 0x67, 0x81, 0xb8, 0xc0, 0x31, 0x28,
							0x95, 0xcf, 0xcb, 0xb5, 0xf5, 0xb7, 0xdf, 0x9c, 0x26, 0x16, 0x2, 0xc7, 0xec,
							0x71, 0xf9, 0xe5, 0x31, 0xae, 0x97,
						},
					},
				},
			},
		},
		{
			title: "drain_delegate",
			src:   "090003139962165dff55ac1510d35113d8c523a3b5ad0014d2bfd6116174b6b413894ea946e1b4f39a8e0b01ffd703746c5850c9df6e35c9e92942f2f708673f",
			kind:  "drain_delegate",
			expect: &DrainDelegate{
				ConsensusKey: &tz.Ed25519PublicKeyHash{
					0x3, 0x13, 0x99, 0x62, 0x16, 0x5d, 0xff, 0x55, 0xac, 0x15, 0x10, 0xd3, 0x51, 0x13, 0xd8,
					0xc5, 0x23, 0xa3, 0xb5, 0xad,
				},
				Delegate: &tz.Ed25519PublicKeyHash{
					0x14, 0xd2, 0xbf, 0xd6, 0x11, 0x61, 0x74, 0xb6, 0xb4, 0x13, 0x89, 0x4e, 0xa9, 0x46, 0xe1,
					0xb4, 0xf3, 0x9a, 0x8e, 0xb,
				},
				Destination: &tz.Secp256k1PublicKeyHash{
					0xff, 0xd7, 0x3, 0x74, 0x6c, 0x58, 0x50, 0xc9, 0xdf, 0x6e, 0x35, 0xc9, 0xe9, 0x29, 0x42,
					0xf2, 0xf7, 0x8, 0x67, 0x3f,
				},
			},
		},
		{
			title: "transaction",
			src:   "6c00a0c7a9b0bcd6a48ee0c13094327f215ba2adeaa7d40dabc1af25e36fde02c096b10201f525eabd8b0eeace1494233ea0230d2c9ad6619b00ffff0b66756c66696c6c5f61736b0000000907070088f0f6010306",
			kind:  "transaction",
			expect: &Transaction{
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
				Destination: core.OriginatedContract{
					ContractHash: &tz.ContractHash{
						0xf5, 0x25, 0xea, 0xbd, 0x8b, 0xe, 0xea, 0xce, 0x14, 0x94, 0x23, 0x3e, 0xa0, 0x23, 0xd,
						0x2c, 0x9a, 0xd6, 0x61, 0x9b,
					},
					Padding: 0x0,
				},
				Parameters: tz.Some(Parameters{
					Entrypoint: EpNamed{String: "fulfill_ask"},
					Value: &expr.Prim20{
						Prim: 7,
						Args: [2]expr.Expression{
							expr.Int{Int: tz.BigInt{0x88, 0xf0, 0xf6, 0x01}},
							expr.Prim00(6),
						},
					},
				}),
			},
		},
		{
			title: "origination",
			src:   "6d0032bbefe9f3df408ad83bc45252d317a43a9104af990bc985ee2add0e9e0a000000000127020000012205000765055f0765076503680369036803590501076507650666076507650368036903680000000725636c61696d7304680000000e25636f6e74726163745f74797065076508610368036900000009256d65746164617461046e00000006256f776e6572050202000000b7037a034c0321057100020317031703480319033c072c020000001807430368010000000d556e617574686f72697a65642e032702000000000743036a000003130319032a072c020000001c07430368010000001154657a206e6f742061636365707465642e03270200000000037a05210003031703170521000403170316034205210004031603170570000403160316057000030552020000000c034c0521000505700002035005700003032003420342053d036d0342000002a207070707020000022207070707010000006f6b65706c65723a2f2f7a43543568746b654472644c7a4c4a516d667a664c657878444c656d59647159756d464573375037366463476f7a5550517867682f7a6233385342776b42546367315a4151654c514b6b4c414577624437713146636d744a77654358345373584651335378650a00000020c7eadbb0e91498fbc4e6392fd951eb0127fbcaf8d6bea5d297ae2f58604197be010000001456657269666961626c6543726564656e7469616c07070707010000006f6b65706c65723a2f2f7a43543568746b654472644c7a4c4a516d667a664c657878444c656d59647159756d464573375037366463476f7a5550517867682f7a62333853437159506e7955367a743476533739476e72637050424e4c357937525646466964655254624a6972364a48430a00000020b5f850226add36a33a9ac6365f2eedc89a99d3ace7b35afa3a15e67c687fbf05010000001456657269666961626c6543726564656e7469616c07070707010000006f6b65706c65723a2f2f7a43543568746b654472644c7a4c4a516d667a664c657878444c656d59647159756d464573375037366463476f7a5550517867682f7a623338534553775447436963456379426d7558477646694368526472684137356a733535656934414d7969634e4e51660a00000020e76fd1cee09b5d79a8bd99db252c864194cccbc98fcf7afebc9f9c2baa4de292010000001456657269666961626c6543726564656e7469616c010000000a747a70726f66696c657307070200000038070401000000000a0000002c68747470733a2f2f747a70726f66696c65732e636f6d2f747a69703031365f6d657461646174612e6a736f6e0100000024747a315147486370396b6943453272694a59754663614452447a78523277667364786379",
			kind:  "origination",
			expect: &Origination{
				ManagerOperation: ManagerOperation{
					Source: &tz.Ed25519PublicKeyHash{
						0x32, 0xbb, 0xef, 0xe9, 0xf3, 0xdf, 0x40, 0x8a, 0xd8, 0x3b, 0xc4, 0x52, 0x52, 0xd3, 0x17,
						0xa4, 0x3a, 0x91, 0x4, 0xaf,
					},
					Fee:          tz.BigUint{0x99, 0xb},
					Counter:      tz.BigUint{0xc9, 0x85, 0xee, 0x2a},
					GasLimit:     tz.BigUint{0xdd, 0xe},
					StorageLimit: tz.BigUint{0x9e, 0xa},
				},
				Balance:  tz.BigUint{0x0},
				Delegate: tz.None[tz.PublicKeyHash](),
				Script: Script{
					Code: expr.Seq{
						Value: []expr.Expression{
							&expr.Prim10{
								Prim: expr.Prim_parameter,
								Arg: &expr.Prim20{
									Prim: expr.Prim_pair,
									Args: [2]expr.Expression{
										&expr.Prim10{
											Prim: expr.Prim_list,
											Arg: &expr.Prim20{
												Prim: expr.Prim_pair,
												Args: [2]expr.Expression{
													&expr.Prim20{
														Prim: expr.Prim_pair,
														Args: [2]expr.Expression{
															expr.Prim00(expr.Prim_string),
															expr.Prim00(expr.Prim_bytes),
														},
													},
													expr.Prim00(expr.Prim_string),
												},
											},
										},
										expr.Prim00(expr.Prim_bool),
									},
								},
							},
							&expr.Prim10{
								Prim: expr.Prim_storage,
								Arg: &expr.Prim20{
									Prim: expr.Prim_pair,
									Args: [2]expr.Expression{
										&expr.Prim20{
											Prim: expr.Prim_pair,
											Args: [2]expr.Expression{
												&expr.Prim1X{
													Prim: expr.Prim_set,
													Arg: &expr.Prim20{
														Prim: expr.Prim_pair,
														Args: [2]expr.Expression{
															&expr.Prim20{
																Prim: expr.Prim_pair,
																Args: [2]expr.Expression{
																	expr.Prim00(expr.Prim_string),
																	expr.Prim00(expr.Prim_bytes),
																},
															},
															expr.Prim00(expr.Prim_string),
														},
													},
													Annots: "%claims",
												},
												&expr.Prim0X{
													Prim:   expr.Prim_string,
													Annots: "%contract_type",
												},
											},
										},
										&expr.Prim20{
											Prim: expr.Prim_pair,
											Args: [2]expr.Expression{
												&expr.Prim2X{
													Prim: expr.Prim_big_map,
													Args: [2]expr.Expression{
														expr.Prim00(expr.Prim_string),
														expr.Prim00(expr.Prim_bytes),
													},
													Annots: "%metadata",
												},
												&expr.Prim0X{
													Prim:   expr.Prim_address,
													Annots: "%owner",
												},
											},
										},
									},
								},
							},
							&expr.Prim10{
								Prim: expr.Prim_code,
								Arg: expr.Seq{
									Value: []expr.Expression{
										expr.Prim00(expr.Prim_UNPAIR),
										expr.Prim00(expr.Prim_SWAP),
										expr.Prim00(expr.Prim_DUP),
										&expr.Prim10{
											Prim: expr.Prim_DUG,
											Arg: expr.Int{
												Int: tz.BigInt{2},
											},
										},
										expr.Prim00(expr.Prim_CDR),
										expr.Prim00(expr.Prim_CDR),
										expr.Prim00(expr.Prim_SENDER),
										expr.Prim00(expr.Prim_COMPARE),
										expr.Prim00(expr.Prim_NEQ),
										&expr.Prim20{
											Prim: expr.Prim_IF,
											Args: [2]expr.Expression{
												expr.Seq{
													Value: []expr.Expression{
														&expr.Prim20{
															Prim: expr.Prim_PUSH,
															Args: [2]expr.Expression{
																expr.Prim00(expr.Prim_string),
																expr.String{
																	String: "Unauthorized.",
																},
															},
														},
														expr.Prim00(expr.Prim_FAILWITH),
													},
												},
												expr.Seq{
													Value: []expr.Expression{},
												},
											},
										},
										&expr.Prim20{
											Prim: expr.Prim_PUSH,
											Args: [2]expr.Expression{
												expr.Prim00(expr.Prim_mutez),
												expr.Int{
													Int: tz.BigInt{0},
												},
											},
										},
										expr.Prim00(expr.Prim_AMOUNT),
										expr.Prim00(expr.Prim_COMPARE),
										expr.Prim00(expr.Prim_GT),
										&expr.Prim20{
											Prim: expr.Prim_IF,
											Args: [2]expr.Expression{
												expr.Seq{
													Value: []expr.Expression{
														&expr.Prim20{
															Prim: expr.Prim_PUSH,
															Args: [2]expr.Expression{
																expr.Prim00(expr.Prim_string),
																expr.String{
																	String: "Tez not accepted.",
																},
															},
														},
														expr.Prim00(expr.Prim_FAILWITH),
													},
												},
												expr.Seq{
													Value: []expr.Expression{},
												},
											},
										},
										expr.Prim00(expr.Prim_UNPAIR),
										&expr.Prim10{
											Prim: expr.Prim_DUP,
											Arg: expr.Int{
												Int: tz.BigInt{3},
											},
										},
										expr.Prim00(expr.Prim_CDR),
										expr.Prim00(expr.Prim_CDR),
										&expr.Prim10{
											Prim: expr.Prim_DUP,
											Arg: expr.Int{
												Int: tz.BigInt{4},
											},
										},
										expr.Prim00(expr.Prim_CDR),
										expr.Prim00(expr.Prim_CAR),
										expr.Prim00(expr.Prim_PAIR),
										&expr.Prim10{
											Prim: expr.Prim_DUP,
											Arg: expr.Int{
												Int: tz.BigInt{4},
											},
										},
										expr.Prim00(expr.Prim_CAR),
										expr.Prim00(expr.Prim_CDR),
										&expr.Prim10{
											Prim: expr.Prim_DIG,
											Arg: expr.Int{
												Int: tz.BigInt{4},
											},
										},
										expr.Prim00(expr.Prim_CAR),
										expr.Prim00(expr.Prim_CAR),
										&expr.Prim10{
											Prim: expr.Prim_DIG,
											Arg: expr.Int{
												Int: tz.BigInt{3},
											},
										},
										&expr.Prim10{
											Prim: expr.Prim_ITER,
											Arg: expr.Seq{
												Value: []expr.Expression{
													expr.Prim00(expr.Prim_SWAP),
													&expr.Prim10{
														Prim: expr.Prim_DUP,
														Arg: expr.Int{
															Int: tz.BigInt{5},
														},
													},
													&expr.Prim10{
														Prim: expr.Prim_DIG,
														Arg: expr.Int{
															Int: tz.BigInt{2},
														},
													},
													expr.Prim00(expr.Prim_UPDATE),
												},
											},
										},
										&expr.Prim10{
											Prim: expr.Prim_DIG,
											Arg: expr.Int{
												Int: tz.BigInt{3},
											},
										},
										expr.Prim00(expr.Prim_DROP),
										expr.Prim00(expr.Prim_PAIR),
										expr.Prim00(expr.Prim_PAIR),
										&expr.Prim10{
											Prim: expr.Prim_NIL,
											Arg:  expr.Prim00(expr.Prim_operation),
										},
										expr.Prim00(expr.Prim_PAIR),
									},
								},
							},
						},
					},
					Storage: &expr.Prim20{
						Prim: expr.Prim_Pair,
						Args: [2]expr.Expression{
							&expr.Prim20{
								Prim: expr.Prim_Pair,
								Args: [2]expr.Expression{
									expr.Seq{
										Value: []expr.Expression{
											&expr.Prim20{
												Prim: expr.Prim_Pair,
												Args: [2]expr.Expression{
													&expr.Prim20{
														Prim: expr.Prim_Pair,
														Args: [2]expr.Expression{
															expr.String{
																String: "kepler://zCT5htkeDrdLzLJQmfzfLexxDLemYdqYumFEs7P76dcGozUPQxgh/zb38SBwkBTcg1ZAQeLQKkLAEwbD7q1FcmtJweCX4SsXFQ3Sxe",
															},
															expr.Bytes{
																Bytes: []uint8{
																	0xc7, 0xea, 0xdb, 0xb0, 0xe9, 0x14, 0x98, 0xfb, 0xc4, 0xe6, 0x39, 0x2f, 0xd9, 0x51, 0xeb, 0x01,
																	0x27, 0xfb, 0xca, 0xf8, 0xd6, 0xbe, 0xa5, 0xd2, 0x97, 0xae, 0x2f, 0x58, 0x60, 0x41, 0x97, 0xbe,
																},
															},
														},
													},
													expr.String{
														String: "VerifiableCredential",
													},
												},
											},
											&expr.Prim20{
												Prim: expr.Prim_Pair,
												Args: [2]expr.Expression{
													&expr.Prim20{
														Prim: expr.Prim_Pair,
														Args: [2]expr.Expression{
															expr.String{
																String: "kepler://zCT5htkeDrdLzLJQmfzfLexxDLemYdqYumFEs7P76dcGozUPQxgh/zb38SCqYPnyU6zt4vS79GnrcpPBNL5y7RVFFideRTbJir6JHC",
															},
															expr.Bytes{
																Bytes: []uint8{
																	0xb5, 0xf8, 0x50, 0x22, 0x6a, 0xdd, 0x36, 0xa3, 0x3a, 0x9a, 0xc6, 0x36, 0x5f, 0x2e, 0xed, 0xc8,
																	0x9a, 0x99, 0xd3, 0xac, 0xe7, 0xb3, 0x5a, 0xfa, 0x3a, 0x15, 0xe6, 0x7c, 0x68, 0x7f, 0xbf, 0x05,
																},
															},
														},
													},
													expr.String{
														String: "VerifiableCredential",
													},
												},
											},
											&expr.Prim20{
												Prim: expr.Prim_Pair,
												Args: [2]expr.Expression{
													&expr.Prim20{
														Prim: expr.Prim_Pair,
														Args: [2]expr.Expression{
															expr.String{
																String: "kepler://zCT5htkeDrdLzLJQmfzfLexxDLemYdqYumFEs7P76dcGozUPQxgh/zb38SESwTGCicEcyBmuXGvFiChRdrhA75js55ei4AMyicNNQf",
															},
															expr.Bytes{
																Bytes: []uint8{
																	0xe7, 0x6f, 0xd1, 0xce, 0xe0, 0x9b, 0x5d, 0x79, 0xa8, 0xbd, 0x99, 0xdb, 0x25, 0x2c, 0x86, 0x41,
																	0x94, 0xcc, 0xcb, 0xc9, 0x8f, 0xcf, 0x7a, 0xfe, 0xbc, 0x9f, 0x9c, 0x2b, 0xaa, 0x4d, 0xe2, 0x92,
																},
															},
														},
													},
													expr.String{
														String: "VerifiableCredential",
													},
												},
											},
										},
									},
									expr.String{
										String: "tzprofiles",
									},
								},
							},
							&expr.Prim20{
								Prim: expr.Prim_Pair,
								Args: [2]expr.Expression{
									expr.Seq{
										Value: []expr.Expression{
											&expr.Prim20{
												Prim: expr.Prim_Elt,
												Args: [2]expr.Expression{
													expr.String{},
													expr.Bytes{
														Bytes: []uint8("https://tzprofiles.com/tzip016_metadata.json"),
													},
												},
											},
										},
									},
									expr.String{
										String: "tz1QGHcp9kiCE2riJYuFcaDRDzxR2wfsdxcy",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			title: "register_global_constant",
			src:   "6f019965ccdba00e7ae3a73ad513cd315a1a59e57f82e303f9edae05b60a500000000f065f03620000000725633363376639",
			kind:  "register_global_constant",
			expect: &RegisterGlobalConstant{
				ManagerOperation: ManagerOperation{
					Source: &tz.Secp256k1PublicKeyHash{
						0x99, 0x65, 0xcc, 0xdb, 0xa0, 0xe, 0x7a, 0xe3, 0xa7, 0x3a, 0xd5, 0x13, 0xcd, 0x31, 0x5a,
						0x1a, 0x59, 0xe5, 0x7f, 0x82,
					},
					Fee:          tz.BigUint{0xe3, 0x3},
					Counter:      tz.BigUint{0xf9, 0xed, 0xae, 0x5},
					GasLimit:     tz.BigUint{0xb6, 0xa},
					StorageLimit: tz.BigUint{0x50},
				},
				Value: &expr.Prim1X{
					Prim:   expr.Prim_list,
					Arg:    expr.Prim00(expr.Prim_nat),
					Annots: "%c3c7f9",
				},
			},
		},
		{
			title: "set_deposits_limit",
			src:   "7000aed89ada8f3fceeabed1caeeb9f86180be29764ccd02849bea08e80700ff80b2a7ad27",
			kind:  "set_deposits_limit",
			expect: &SetDepositsLimit{
				ManagerOperation: ManagerOperation{
					Source: &tz.Ed25519PublicKeyHash{
						0xae, 0xd8, 0x9a, 0xda, 0x8f, 0x3f, 0xce, 0xea, 0xbe, 0xd1, 0xca, 0xee, 0xb9, 0xf8, 0x61,
						0x80, 0xbe, 0x29, 0x76, 0x4c,
					},
					Fee:          tz.BigUint{0xcd, 0x2},
					Counter:      tz.BigUint{0x84, 0x9b, 0xea, 0x8},
					GasLimit:     tz.BigUint{0xe8, 0x7},
					StorageLimit: tz.BigUint{0x0},
				},
				Limit: tz.Some(tz.BigUint{0x80, 0xb2, 0xa7, 0xad, 0x27}),
			},
		},
		{
			title: "increase_paid_storage",
			src:   "7101b551502ad17818169275aec7a90d3736310e673c9303f7820bcc0800010188fbff5eba69fea674fa2791b575e743664ed92100",
			kind:  "increase_paid_storage",
			expect: &IncreasePaidStorage{
				ManagerOperation: ManagerOperation{
					Source: &tz.Secp256k1PublicKeyHash{
						0xb5, 0x51, 0x50, 0x2a, 0xd1, 0x78, 0x18, 0x16, 0x92, 0x75, 0xae, 0xc7, 0xa9, 0xd, 0x37,
						0x36, 0x31, 0xe, 0x67, 0x3c,
					},
					Fee:          tz.BigUint{0x93, 0x3},
					Counter:      tz.BigUint{0xf7, 0x82, 0xb},
					GasLimit:     tz.BigUint{0xcc, 0x8},
					StorageLimit: tz.BigUint{0x0},
				},
				Amount: tz.BigInt{0x1},
				Destination: core.OriginatedContract{
					ContractHash: &tz.ContractHash{
						0x88, 0xfb, 0xff, 0x5e, 0xba, 0x69, 0xfe, 0xa6, 0x74, 0xfa, 0x27, 0x91, 0xb5, 0x75,
						0xe7, 0x43, 0x66, 0x4e, 0xd9, 0x21,
					},
					Padding: 0x0,
				},
			},
		},
		{
			title: "update_consensus_key",
			src:   "7201c1af3f80e52a44bc3d5612194a2f5e3f5b6fb3a09803f6820bcc08000202ada8d11ad108191ca48f62b7fabe5d552df0a0feed4d18dc7ae8944ad416b95d",
			kind:  "update_consensus_key",
			expect: &UpdateConsensusKey{
				ManagerOperation: ManagerOperation{
					Source: &tz.Secp256k1PublicKeyHash{
						0xc1, 0xaf, 0x3f, 0x80, 0xe5, 0x2a, 0x44, 0xbc, 0x3d, 0x56, 0x12, 0x19, 0x4a, 0x2f, 0x5e,
						0x3f, 0x5b, 0x6f, 0xb3, 0xa0,
					},
					Fee:          tz.BigUint{0x98, 0x3},
					Counter:      tz.BigUint{0xf6, 0x82, 0xb},
					GasLimit:     tz.BigUint{0xcc, 0x8},
					StorageLimit: tz.BigUint{0x0},
				},
				PublicKey: &tz.P256PublicKey{
					0x2, 0xad, 0xa8, 0xd1, 0x1a, 0xd1, 0x8, 0x19, 0x1c, 0xa4, 0x8f, 0x62, 0xb7, 0xfa, 0xbe,
					0x5d, 0x55, 0x2d, 0xf0, 0xa0, 0xfe, 0xed, 0x4d, 0x18, 0xdc, 0x7a, 0xe8, 0x94, 0x4a, 0xd4,
					0x16, 0xb9, 0x5d,
				},
			},
		},
	}

	for _, tc := range testCases {
		test := tc
		t.Run(test.title, func(t *testing.T) {
			data, err := hex.DecodeString(test.src)
			require.NoError(t, err)
			t.Run("Decode", func(t *testing.T) {
				var op OperationContents
				rest, err := encoding.Decode(data, &op)
				require.NoError(t, err)
				require.Empty(t, rest)
				//fmt.Printf("%# v\n", pretty.Formatter(op, pretty.OptStringer(true)))
				require.Equal(t, test.expect, op)
				require.Equal(t, test.kind, op.OperationKind())
			})

			if !test.skipEncode {
				t.Run("Encode", func(t *testing.T) {
					var buf bytes.Buffer
					require.NoError(t, encoding.Encode(&buf, &test.expect))
					require.Equal(t, data, buf.Bytes())
				})
			}
		})
	}
}
