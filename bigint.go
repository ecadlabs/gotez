package gotez

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ecadlabs/gotez/v2/encoding"
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

var (
	mask6 = big.NewInt(0x3f)
	mask7 = big.NewInt(0x7f)
)

func NewBigInt(value *big.Int) BigInt {
	var (
		neg bool
		val big.Int
	)
	if value.Sign() < 0 {
		neg = true
		val.Neg(value)
	} else {
		val.Set(value)
	}

	res := make(BigInt, 0, (val.BitLen()+7)/7+1)
	var tmp big.Int
	for i := 0; ; i++ {
		var (
			mask *big.Int
			s    uint
		)
		if i == 0 {
			mask = mask6
			s = 6
		} else {
			mask = mask7
			s = 7
		}
		x := uint8(tmp.And(&val, mask).Int64())
		val.Rsh(&val, s)
		if i == 0 && neg {
			x |= 0x40
		}
		if val.Sign() != 0 {
			res = append(res, x|0x80)
		} else {
			res = append(res, x)
			break
		}
	}
	return res
}

func NewBigInt64(val int64) BigInt {
	var (
		neg bool
	)
	if val < 0 {
		neg = true
		val = -val
	}

	res := make(BigInt, 0, 11)
	for i := 0; ; i++ {
		var (
			mask int64
			s    uint
		)
		if i == 0 {
			mask = 0x3f
			s = 6
		} else {
			mask = 0x7f
			s = 7
		}
		x := uint8(val & mask)
		val >>= s
		if i == 0 && neg {
			x |= 0x40
		}
		if val != 0 {
			res = append(res, x|0x80)
		} else {
			res = append(res, x)
			break
		}
	}
	return res
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
func (b BigInt) Sign() int {
	switch {
	case len(b) == 0 || b[0] == 0:
		return 0
	case b[0]&0x40 != 0:
		return -1
	default:
		return 1
	}
}

func (b BigInt) String() string {
	return b.Int().String()
}

func (b BigInt) MarshalText() (text []byte, err error) {
	return b.Int().MarshalText()
}

type BigUint []byte

func NewBigUint(value *big.Int) (BigUint, error) {
	var val big.Int
	if value.Sign() < 0 {
		return nil, errors.New("(biguint) negative value")
	} else {
		val.Set(value)
	}

	res := make(BigUint, 0, (val.BitLen()+7)/7+1)
	var tmp big.Int
	for i := 0; ; i++ {
		x := uint8(tmp.And(&val, mask7).Int64())
		val.Rsh(&val, 7)
		if val.Sign() != 0 {
			res = append(res, x|0x80)
		} else {
			res = append(res, x)
			break
		}
	}
	return res, nil
}

func NewBigUint64(val uint64) BigUint {
	res := make(BigUint, 0, 11)
	for i := 0; ; i++ {
		x := uint8(val & 0x7f)
		val >>= 7
		if val != 0 {
			res = append(res, x|0x80)
		} else {
			res = append(res, x)
			break
		}
	}
	return res
}

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

func (b BigUint) IsZero() bool {
	return len(b) == 0 || b[0] == 0
}

func (b BigUint) String() string {
	return b.Int().String()
}

func (b BigUint) MarshalText() (text []byte, err error) {
	return b.Int().MarshalText()
}

func BigZero() BigInt   { return BigInt{0} }
func BigUZero() BigUint { return BigUint{0} }
