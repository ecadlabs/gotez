package rlp

import (
	"errors"
	"fmt"
	"io"
	"math"
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

func getUint(s []byte) (val uint64, err error) {
	if len(s) == 0 {
		return 0, ErrEOS
	}

	var i int
	for i < len(s) && s[i] == 0 {
		i++
	}

	if rem := len(s) - i; rem > 8 {
		return 0, ErrOverflow
	} else if rem != 0 {
		shift := (rem - 1) * 8
		for i < len(s) {
			val |= uint64(s[i]) << shift
			shift -= 8
			i++
		}
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
		ln, err := getUint(s[i : i+lnLn])
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

func (s *String) Uint64() (uint64, error) {
	n, list, val, err := getElem(*s)
	if err != nil {
		return 0, err
	} else if list {
		return 0, ErrType
	}
	*s = (*s)[n:]
	return getUint(val)
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

func Uint[T ~uint8 | ~uint16 | ~uint32 | ~uint64](s *String) (T, error) {
	v, err := s.Uint64()
	if err != nil {
		return 0, err
	}
	if v > uint64(^T(0)) {
		return 0, ErrOverflow
	}
	return T(v), nil
}

type ByteSetter[T any] interface {
	SetBytes(buf []byte) T
}

func SetBytes[T ByteSetter[R], R any](s *String, dst T) error {
	n, list, val, err := getElem(*s)
	if err != nil {
		return err
	} else if list {
		return ErrType
	}
	dst.SetBytes(val)
	*s = (*s)[n:]
	return nil
}
