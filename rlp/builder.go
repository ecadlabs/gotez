package rlp

import (
	"math/big"
	"unsafe"
)

type Marshaler interface {
	MarshalRLP(b *Builder)
}

type Builder struct {
	buffer []byte
	err    error
}

func NewBuilder(buf []byte) *Builder {
	return &Builder{
		buffer: buf,
	}
}

func encodeUint(v uint64) []byte {
	var buf [8]byte
	i := 8
	for {
		i--
		buf[i] = byte(v & 0xff)
		v >>= 8
		if v == 0 {
			break
		}
	}
	return buf[i:]
}

func (b *Builder) writeString(data []byte) {
	if len(data) == 0 {
		return
	}
	switch {
	case len(data) == 1 && data[0] < 0x80:
		b.buffer = append(b.buffer, data[0])
	case len(data) < 56:
		b.buffer = append(b.buffer, 0x80+byte(len(data)))
		b.buffer = append(b.buffer, data...)
	default:
		l := encodeUint(uint64(len(data)))
		b.buffer = append(b.buffer, 0xb7+byte(len(l)))
		b.buffer = append(b.buffer, l...)
		b.buffer = append(b.buffer, data...)
	}
}

func (b *Builder) writeList(data []byte) {
	if len(data) == 0 {
		return
	}
	switch {
	case len(data) < 56:
		b.buffer = append(b.buffer, 0xc0+byte(len(data)))
		b.buffer = append(b.buffer, data...)
	default:
		l := encodeUint(uint64(len(data)))
		b.buffer = append(b.buffer, 0xf7+byte(len(l)))
		b.buffer = append(b.buffer, l...)
		b.buffer = append(b.buffer, data...)
	}
}

func (b *Builder) AddString(data []byte) {
	if b.err != nil {
		return
	}
	b.writeString(data)
}

func (b *Builder) AddUint(v uint64) {
	b.writeString(encodeUint(v))
}

func (b *Builder) AddBigInt(v *big.Int) {
	var buf []byte
	if v.Sign() == 0 {
		tmp := [1]byte{0}
		buf = tmp[:]
	} else {
		buf = v.Bytes()
	}
	b.writeString(buf)
}

func (b *Builder) AddBigIntLE(val *big.Int, size int) {
	if b.err != nil {
		return
	}
	var words []big.Word
	if val.Sign() == 0 {
		tmp := [1]big.Word{0}
		words = tmp[:]
	} else {
		words = val.Bits()
	}
	trim := false
	if size == 0 {
		size = len(words) * wordBytes
		trim = true
	}
	buf := make([]byte, size)
	i := 0
Loop:
	for j, x := range words {
		for range wordBytes {
			if i == len(buf) {
				panic("rlp: buffer is too short")
			}
			buf[i] = byte(x & 0xff)
			i++
			x >>= 8
			if j == len(words)-1 && x == 0 {
				break Loop
			}
		}
	}
	if trim {
		buf = buf[:i]
	}
	b.writeString(buf)
}

func (b *Builder) AddList(child func(*Builder)) {
	if b.err != nil {
		return
	}
	cb := Builder{
		buffer: make([]byte, 0, cap(b.buffer)),
	}
	child(&cb)
	if cb.err != nil {
		b.err = cb.err
		return
	}
	b.writeList(cb.buffer)
}

func (b *Builder) AddElem(raw []byte) {
	b.buffer = append(b.buffer, raw...)
}

func (b *Builder) Add(elem Marshaler) {
	elem.MarshalRLP(b)
}

func AddUintLE[T ~uint8 | ~uint16 | ~uint32 | ~uint64](b *Builder, v T) {
	if b.err != nil {
		return
	}
	var buf [8]byte
	i := 0
	t := uint64(v)
	for t != 0 {
		buf[i] = byte(t & 0xff)
		t >>= 8
		i++
	}
	b.writeString(buf[:unsafe.Sizeof(v)])
}

func (b *Builder) Bytes() ([]byte, error) {
	buf := b.buffer
	if buf == nil {
		buf = make([]byte, 0)
	}
	return buf, b.err
}

func Marshal[T Marshaler](v T) ([]byte, error) {
	var b Builder
	b.Add(v)
	return b.Bytes()
}
