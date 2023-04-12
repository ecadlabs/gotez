package gotez

import (
	"fmt"
	"math/big"

	"github.com/ecadlabs/gotez/encoding"
)

type BigInt []byte

func getLen(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, fmt.Errorf("(bigint) %w", encoding.ErrBuffer(1))
	}
	i := 0
	for i < len(data) && data[i]&0x80 != 0 {
		i += 1
	}
	if i == len(data) {
		return 0, fmt.Errorf("(bigint) %w", encoding.ErrBuffer(i))
	}
	return i + 1, nil
}

func (b *BigInt) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	ln, err := getLen(data)
	if err != nil {
		return nil, err
	}
	*b = BigInt(data[:ln])
	return data[ln:], nil
}

func (b BigInt) Int() *big.Int {
	res := big.NewInt(0)
	if len(b) == 0 {
		return res
	}
	neg := b[0]&0x40 != 0
	shift := uint(0)
	var tmp big.Int
	for i, x := range b {
		var (
			mask uint8
			s    uint
		)
		if i == 0 {
			mask = 0x3f
			s = 6
		} else {
			mask = 0x7f
			s = 7
		}
		tmp.SetInt64(int64(x & mask))
		tmp.Lsh(&tmp, shift)
		res.Or(res, &tmp)
		shift += s
	}
	if neg {
		res.Neg(res)
	}
	return res
}

func (b BigInt) String() string {
	return b.Int().String()
}

func (b BigInt) MarshalText() (text []byte, err error) {
	return b.Int().MarshalText()
}

type BigUint []byte

func (b *BigUint) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	ln, err := getLen(data)
	if err != nil {
		return nil, err
	}
	*b = BigUint(data[:ln])
	return data[ln:], nil
}

func (b BigUint) Int() *big.Int {
	res := big.NewInt(0)
	if len(b) == 0 {
		return res
	}
	shift := uint(0)
	var tmp big.Int
	for _, x := range b {
		tmp.SetInt64(int64(x & 0x7f))
		tmp.Lsh(&tmp, shift)
		res.Or(res, &tmp)
		shift += 7
	}
	return res
}

func (b BigUint) String() string {
	return b.Int().String()
}

func (b BigUint) MarshalText() (text []byte, err error) {
	return b.Int().MarshalText()
}
