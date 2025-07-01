package rlp

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
)

type String []byte

var (
	ErrEOS           = errors.New("rlp: end of byte string")
	ErrUnexpectedEOS = errors.New("rlp: unexpected end of byte string")
	ErrType          = errors.New("rlp: invalid field type")
	ErrOverflow      = errors.New("rlp: integer overflow")
)

type errEOS struct {
	ln int
	n  int
}

func (e errEOS) Error() string {
	if e.n != 0 {
		return fmt.Sprintf("rlp: unexpected end of byte string: expected %d bytes, have %d left", e.n, e.ln)
	} else {
		return fmt.Sprintf("rlp: unexpected end of byte string: have %d bytes left", e.ln)
	}
}

func (errEOS) Is(target error) bool {
	return target == ErrUnexpectedEOS || target == io.ErrUnexpectedEOF
}

func getUintBE(s []byte) (val uint64, err error) {
	if len(s) == 0 {
		return 0, ErrEOS
	}
	for _, x := range s {
		if val >= 1<<56 {
			return 0, ErrOverflow
		}
		val = (val << 8) | uint64(x)
	}
	return val, nil
}

func getElem(s []byte) (consumed int, list bool, val []byte, err error) {
	if len(s) == 0 {
		return 0, false, nil, ErrEOS
	}

	var i int
	prefix := s[0]
	switch {
	case prefix < 0x80:
		val = s[:1]
		i++

	case prefix < 0xb8 || prefix >= 0xc0 && prefix < 0xf8:
		i++
		list = prefix >= 0xc0
		var base byte
		if list {
			base = 0xc0
		} else {
			base = 0x80
		}

		ln := int(prefix - base)
		if len(s)-i < ln {
			return 0, false, nil, errEOS{ln: len(s) - i, n: ln}
		}

		val = s[i : i+ln]
		i += ln

	default:
		i++
		list = prefix >= 0xf8
		var base byte
		if list {
			base = 0xf7
		} else {
			base = 0xb7
		}

		lnLn := int(prefix - base)
		ln, err := getUintBE(s[i : i+lnLn])
		if err != nil {
			return 0, false, nil, err
		}
		i += lnLn

		if ln > math.MaxInt {
			return 0, false, nil, ErrOverflow
		}
		if len(s)-i < int(ln) {
			return 0, false, nil, errEOS{ln: len(s) - i, n: int(ln)}
		}
		val = s[i : i+int(ln)]
		i += int(ln)
	}
	return i, list, val, nil
}

func (s *String) Uint64BE() (uint64, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return 0, err
	} else if list {
		return 0, ErrType
	}
	res, err := getUintBE(val)
	if err != nil {
		return 0, err
	}
	*s = (*s)[n:]
	return res, nil
}

func (s *String) Uint64LE() (uint64, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return 0, err
	} else if list {
		return 0, ErrType
	}
	var (
		res   uint64
		shift uint
	)
	for _, x := range val {
		if x != 0 && shift > 56 {
			return 0, ErrOverflow
		}
		res = res | (uint64(x) << shift)
		shift += 8
	}
	*s = (*s)[n:]
	return res, nil
}

func (s *String) BigIntBE() (*big.Int, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return nil, err
	} else if list {
		return nil, ErrType
	}
	res := new(big.Int).SetBytes(val)
	*s = (*s)[n:]
	return res, nil
}

const WS = bits.UintSize / 8

func (s *String) BigIntLE() (*big.Int, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return nil, err
	} else if list {
		return nil, ErrType
	}
	res := new(big.Int)
	var (
		shift uint
		t     big.Int
	)
	for _, x := range val {
		t.SetInt64(int64(x))
		t.Lsh(&t, shift)
		res.Add(res, &t)
		shift += 8
	}
	*s = (*s)[n:]
	return res, nil
}

func (s *String) Bytes() ([]byte, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return nil, err
	} else if list {
		return nil, ErrType
	}
	*s = (*s)[n:]
	return val, nil
}

func (s *String) List() (String, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return nil, err
	} else if !list {
		return nil, ErrType
	}
	*s = (*s)[n:]
	return String(val), nil
}

func (s *String) Raw() (String, error) {
	n, _, _, err := getElem(*s)
	if err != nil {
		return nil, err
	}
	raw := (*s)[:n]
	*s = (*s)[n:]
	return String(raw), nil
}

func (s *String) RawList() ([]String, error) {
	list, err := s.List()
	if err != nil {
		return nil, err
	}
	res := make([]String, 0)
	for {
		elem, err := list.Raw()
		if err != nil {
			if err == ErrEOS {
				break
			}
			return nil, err
		}
		res = append(res, elem)
	}
	return res, nil
}

func UintBE[T ~uint8 | ~uint16 | ~uint32 | ~uint64](s *String) (T, error) {
	v, err := s.Uint64BE()
	if err != nil {
		return 0, err
	}
	if v > uint64(^T(0)) {
		return 0, ErrOverflow
	}
	return T(v), nil
}

func UintLE[T ~uint8 | ~uint16 | ~uint32 | ~uint64](s *String) (T, error) {
	v, err := s.Uint64LE()
	if err != nil {
		return 0, err
	}
	if v > uint64(^T(0)) {
		return 0, ErrOverflow
	}
	return T(v), nil
}
