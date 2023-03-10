package base58

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/ecadlabs/gotez/b58/prefix"
)

const alphabetStart = 49

var base58alphabetF = []int8{
	0, 1, 2, 3, 4, 5, 6,
	7, 8, -1, -1, -1, -1, -1, -1,
	-1, 9, 10, 11, 12, 13, 14, 15,
	16, -1, 17, 18, 19, 20, 21, -1,
	22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, -1, -1, -1, -1, -1,
	-1, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, -1, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57,
}

var base58alphabetB = []int8{
	0, 1, 2, 3, 4, 5, 6, 7,
	8, 16, 17, 18, 19, 20, 21, 22,
	23, 25, 26, 27, 28, 29, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40,
	41, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 60, 61, 62, 63,
	64, 65, 66, 67, 68, 69, 70, 71,
	72, 73,
}

func Decode(src []byte) ([]byte, error) {
	i := 0
	// count and skip leading zeros
	for ; i < len(src); i++ {
		c := int(src[i]) - alphabetStart
		if c < 0 || c >= len(base58alphabetF) || base58alphabetF[c] == -1 {
			return nil, fmt.Errorf("gotez: base58 decoding error: unexpected character at position %d: %c", i, src[i])
		}
		if base58alphabetF[c] != 0 {
			break
		}
	}
	zeros := i
	acc := make([]byte, 0, len(src)/4)
	for ; i < len(src); i++ {
		c := int(src[i]) - alphabetStart
		if c < 0 || c >= len(base58alphabetF) || base58alphabetF[c] == -1 {
			return nil, fmt.Errorf("gotez: base58 decoding error: unexpected character at position %d: %c", i, src[i])
		}
		carry := int(base58alphabetF[c])
		// for every symbol x
		// acc = acc * 58 + x
		// where acc is a little endian arbitrary length integer
		for ii := 0; carry != 0 || ii < len(acc); ii++ {
			var a int
			if ii < len(acc) {
				a = int(acc[ii])
			}
			m := a*58 + carry
			b := m % 256
			carry = m / 256
			if ii < len(acc) {
				acc[ii] = byte(b)
			} else {
				acc = append(acc, byte(b))
			}
		}
	}
	out := make([]byte, len(acc)+zeros)
	for i := 0; i < len(acc); i++ {
		out[i+zeros] = acc[len(acc)-i-1]
	}
	return out, nil
}

func Encode(src []byte) []byte {
	i := 0
	// count and skip leading zeros
	for ; i < len(src) && src[i] == 0; i++ {
	}
	zeros := i
	acc := make([]byte, 0, len(src)*5)
	for ; i < len(src); i++ {
		carry := int(src[i])
		for ii := 0; carry != 0 || ii < len(acc); ii++ {
			var a int
			if ii < len(acc) {
				a = int(acc[ii])
			}
			m := a*256 + carry
			b := m % 58
			carry = m / 58
			if ii < len(acc) {
				acc[ii] = byte(b)
			} else {
				acc = append(acc, byte(b))
			}
		}
	}
	out := make([]byte, len(acc)+zeros)
	for i := 0; i < zeros; i++ {
		out[i] = alphabetStart
	}
	for i := 0; i < len(acc); i++ {
		out[i+zeros] = byte(base58alphabetB[acc[len(acc)-i-1]] + alphabetStart)
	}
	return out
}

func DecodeCheck(src []byte) ([]byte, error) {
	buf, err := Decode(src)
	if err != nil {
		return nil, err
	}
	if len(buf) < 4 {
		return nil, fmt.Errorf("gotez: base58Check decoding error: data is too short: %d", len(buf))
	}
	data := buf[:len(buf)-4]
	sum := buf[len(buf)-4:]
	s0 := sha256.Sum256(data)
	s1 := sha256.Sum256(s0[:])
	if !bytes.Equal(sum, s1[:4]) {
		return nil, errors.New("gotez: base58Check decoding error: invalid checksum")
	}
	return data, nil
}

func EncodeCheck(data []byte) []byte {
	s0 := sha256.Sum256(data)
	s1 := sha256.Sum256(s0[:])
	return Encode(append(data, s1[:4]...))
}

// ErrPrefix is returned in case of unknown Tezos base58 prefix
var ErrPrefix = errors.New("gotez: unknown Tezos base58 prefix")

func DecodeTZ(data []byte) (pre *prefix.Prefix, payload []byte, err error) {
	buf, err := DecodeCheck(data)
	if err != nil {
		return
	}
	for _, p := range prefix.List {
		if bytes.HasPrefix(buf, p.Prefix) {
			plLen := len(buf) - len(p.Prefix)
			if p.Len != 0 && plLen != p.Len {
				return p, nil, fmt.Errorf("gotez: invalid base58 message length: expected %d, got %d", p.Len, plLen)
			}
			return p, buf[len(p.Prefix):], nil
		}
	}
	err = ErrPrefix
	return
}

func EncodeTZ(pre *prefix.Prefix, payload []byte) ([]byte, error) {
	if pre.Len != 0 && len(payload) != pre.Len {
		return nil, fmt.Errorf("gotez: invalid base58 message length: expected %d, got %d", pre.Len, len(payload))
	}
	data := make([]byte, len(pre.Prefix)+len(payload))
	copy(data, pre.Prefix)
	copy(data[len(pre.Prefix):], payload)
	return EncodeCheck(data), nil
}
