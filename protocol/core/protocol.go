package core

import (
	"errors"
	"strconv"

	tz "github.com/ecadlabs/gotez"
)

type Protocol uint8

const (
	Proto000Ps9mPmXa Protocol = iota
	Proto001PtCJ7pwo
	Proto002PsYLVpVv
	Proto003PsddFKi3
	Proto004Pt24m4xi
	Proto005PsBabyM1
	Proto006PsCARTHA
	Proto007PsDELPH1
	Proto008PtEdo2Zk
	Proto009PsFLoren
	Proto010PtGRANAD
	Proto011PtHangz2
	Proto012Psithaca
	Proto013PtJakart
	Proto014PtKathma
	Proto015PtLimaPt
	Proto016PtMumbai
)

var protoNames = map[Protocol]string{
	Proto000Ps9mPmXa: "000_Ps9mPmXa",
	Proto001PtCJ7pwo: "001_PtCJ7pwo",
	Proto002PsYLVpVv: "002_PsYLVpVv",
	Proto003PsddFKi3: "003_PsddFKi3",
	Proto004Pt24m4xi: "004_Pt24m4xi",
	Proto005PsBabyM1: "005_PsBabyM1",
	Proto006PsCARTHA: "006_PsCARTHA",
	Proto007PsDELPH1: "007_PsDELPH1",
	Proto008PtEdo2Zk: "008_PtEdo2Zk",
	Proto009PsFLoren: "009_PsFLoren",
	Proto010PtGRANAD: "010_PtGRANAD",
	Proto011PtHangz2: "011_PtHangz2",
	Proto012Psithaca: "012_Psithaca",
	Proto013PtJakart: "013_PtJakart",
	Proto014PtKathma: "014_PtKathma",
	Proto015PtLimaPt: "015_PtLimaPt",
	Proto016PtMumbai: "016_PtMumbai",
}

func (p Protocol) Name() string {
	if name, ok := protoNames[p]; ok {
		return name
	}
	return strconv.FormatInt(int64(p), 10)
}

func (p Protocol) String() string {
	if h, err := p.Hash(); err == nil {
		return h.String()
	}
	return strconv.FormatInt(int64(p), 10)
}

var protocolHashes = map[Protocol]*tz.ProtocolHash{
	Proto000Ps9mPmXa: {0x38, 0xec, 0xde, 0xf0, 0xcd, 0x8, 0x64, 0xf, 0x31, 0x8a, 0x9b, 0x5, 0x5f, 0x6b, 0xd, 0xc, 0x9a, 0xe0, 0x30, 0x91, 0x3a, 0x87, 0x1d, 0x9b, 0x9d, 0x86, 0xfb, 0x84, 0x63, 0x17, 0xda, 0x21},
	Proto001PtCJ7pwo: {0xc2, 0x5d, 0x8f, 0x24, 0x46, 0xae, 0x4e, 0x33, 0xcf, 0x35, 0x2b, 0xd, 0x82, 0xd2, 0xfb, 0xed, 0x1f, 0x2d, 0xf0, 0x52, 0x8e, 0xfb, 0xf2, 0x2f, 0x31, 0xf2, 0x1a, 0x5, 0x2b, 0x2a, 0x99, 0x38},
	Proto002PsYLVpVv: {0x6c, 0x2c, 0xca, 0x12, 0x99, 0x6, 0x14, 0xa7, 0xcc, 0x55, 0x11, 0x3a, 0xe6, 0x3e, 0x21, 0x36, 0x5, 0xd5, 0x9e, 0x92, 0x65, 0xaa, 0x3f, 0x2b, 0xc, 0xa6, 0x51, 0x1f, 0x1f, 0x98, 0xbc, 0x1},
	Proto003PsddFKi3: {0x78, 0x2f, 0xe, 0x56, 0xd7, 0x1e, 0x26, 0xcf, 0xe0, 0x9e, 0x2e, 0x33, 0xeb, 0x44, 0xb3, 0x2a, 0xd6, 0xd4, 0x67, 0xb7, 0x42, 0xbb, 0x79, 0xd4, 0x61, 0xdb, 0x2d, 0x86, 0xaf, 0x22, 0xad, 0xad},
	Proto004Pt24m4xi: {0xab, 0x22, 0xe4, 0x6e, 0x78, 0x72, 0xaa, 0x13, 0xe3, 0x66, 0xe4, 0x55, 0xbb, 0x4f, 0x5d, 0xbe, 0xde, 0x85, 0x6a, 0xb0, 0x86, 0x4e, 0x1d, 0xa7, 0xe1, 0x22, 0x55, 0x45, 0x79, 0xee, 0x71, 0xf8},
	Proto005PsBabyM1: {0x3d, 0xb, 0x4b, 0xac, 0xb5, 0xc3, 0xe1, 0x52, 0xa1, 0x67, 0xda, 0x26, 0xfe, 0xfc, 0x26, 0x6b, 0xd3, 0xa0, 0xe1, 0x4f, 0xc4, 0xe4, 0x1e, 0x6c, 0x53, 0x62, 0x3b, 0xf4, 0x82, 0x83, 0x3d, 0xa2},
	Proto006PsCARTHA: {0x3e, 0x5e, 0x3a, 0x60, 0x6a, 0xfa, 0xb7, 0x4a, 0x59, 0xca, 0x9, 0xe3, 0x33, 0x63, 0x3e, 0x27, 0x70, 0xb6, 0x49, 0x2c, 0x5e, 0x59, 0x44, 0x55, 0xb7, 0x1e, 0x9a, 0x2f, 0xe, 0xa9, 0x2a, 0xfb},
	Proto007PsDELPH1: {0x40, 0xca, 0xb8, 0x3d, 0x3f, 0x37, 0xa6, 0x4d, 0xa2, 0x6b, 0x57, 0xad, 0x3d, 0x4, 0x32, 0xae, 0x88, 0x12, 0x93, 0xa2, 0x51, 0x69, 0xad, 0xa3, 0x87, 0xbf, 0xc7, 0x4a, 0x1c, 0xbf, 0x9e, 0x6e},
	Proto008PtEdo2Zk: {0xc7, 0xad, 0x4f, 0x7a, 0x0, 0xe, 0x28, 0xe9, 0xee, 0xfc, 0x58, 0xde, 0x8e, 0xa1, 0x17, 0x2d, 0xe8, 0x43, 0x24, 0x2b, 0xd2, 0xe6, 0x88, 0x77, 0x99, 0x53, 0xd3, 0x41, 0x6a, 0x44, 0x64, 0xb},
	Proto009PsFLoren: {0x45, 0x96, 0x28, 0x5c, 0x68, 0x71, 0x69, 0x1e, 0x25, 0x19, 0x6c, 0x6a, 0x8d, 0x26, 0xd9, 0xa, 0x3a, 0xc9, 0x13, 0x75, 0x73, 0x1e, 0x39, 0x26, 0x10, 0x3c, 0x51, 0x7a, 0x13, 0xa0, 0xba, 0x56},
	Proto010PtGRANAD: {0xcb, 0xb9, 0x44, 0xf7, 0x42, 0x44, 0xea, 0x26, 0x81, 0x98, 0x1f, 0x25, 0x99, 0x5f, 0x8e, 0xbb, 0xa8, 0xff, 0x6c, 0xee, 0x8c, 0x3, 0x68, 0x92, 0xfe, 0x90, 0x1c, 0xb7, 0x60, 0xc4, 0xe3, 0x9e},
	Proto011PtHangz2: {0xce, 0x5f, 0x6, 0x1e, 0x34, 0xb5, 0xa2, 0x1f, 0xea, 0xb8, 0xdb, 0xdf, 0xe7, 0x55, 0xef, 0x17, 0xe7, 0xc, 0x9f, 0x56, 0x54, 0x64, 0xf0, 0x67, 0xac, 0x5e, 0x7c, 0x2, 0xbe, 0x83, 0xa, 0x48},
	Proto012Psithaca: {0x84, 0x24, 0x52, 0xc, 0xf9, 0xbb, 0xf0, 0xa4, 0x27, 0x70, 0x20, 0x4d, 0x95, 0xdc, 0xc1, 0xf1, 0x1e, 0x40, 0x4f, 0xdb, 0x3e, 0x90, 0xc8, 0x48, 0x50, 0xc4, 0xcf, 0xdb, 0x50, 0xc5, 0xc4, 0xb9},
	Proto013PtJakart: {0xd0, 0xa3, 0xf0, 0x7b, 0x8a, 0xdf, 0xcf, 0x61, 0xf5, 0xca, 0x60, 0xf2, 0x44, 0xca, 0x9a, 0x87, 0x6e, 0x76, 0xcb, 0xad, 0x91, 0x40, 0x98, 0xf, 0x6c, 0x88, 0xd0, 0xbf, 0x90, 0xa, 0xc6, 0xd8},
	Proto014PtKathma: {0xd2, 0xea, 0x9f, 0x23, 0xa1, 0xa1, 0x1, 0x10, 0x91, 0x84, 0x1b, 0x12, 0xe3, 0x2c, 0xe2, 0xf8, 0xc3, 0xfa, 0xcf, 0xf2, 0x7f, 0xee, 0xe5, 0x8b, 0xb7, 0xc9, 0xe9, 0x5, 0x67, 0xd1, 0x14, 0x25},
	Proto015PtLimaPt: {0xd5, 0x7e, 0xd8, 0x8b, 0xe5, 0xa6, 0x98, 0x15, 0xe3, 0x93, 0x86, 0xa3, 0x3f, 0x7d, 0xca, 0xd3, 0x91, 0xf5, 0xf5, 0x7, 0xe0, 0x3b, 0x37, 0x6e, 0x49, 0x92, 0x72, 0xc8, 0x6c, 0x6c, 0xf2, 0xa7},
	Proto016PtMumbai: {0xd8, 0x32, 0x5f, 0x11, 0xe4, 0x61, 0x4, 0x77, 0xaa, 0x71, 0xca, 0xca, 0x44, 0x1a, 0x50, 0x99, 0x86, 0xdc, 0xec, 0x3, 0x1b, 0xa, 0xac, 0xd8, 0x2a, 0x9e, 0x10, 0x79, 0x2b, 0x37, 0x14, 0x87},
}

func (p Protocol) Hash() (*tz.ProtocolHash, error) {
	if h, ok := protocolHashes[p]; ok {
		return h, nil
	}
	return nil, errors.New("gotez: unknown protocol version")
}

const LatestProtocol = Proto016PtMumbai
