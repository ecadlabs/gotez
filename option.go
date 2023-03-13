package gotez

import (
	"bytes"
	"fmt"

	"github.com/ecadlabs/gotez/encoding"
)

type Option[T any] struct {
	some  bool
	value T
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		some:  true,
		value: val,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

func (op Option[T]) Unwrap() T {
	if !op.some {
		panic(fmt.Sprintf("called `Unwrap()` on a `None` value of type %T", op))
	}
	return op.value
}

func (op Option[T]) UnwrapUnchecked() T {
	return op.value
}

func (op Option[T]) IsSome() bool { return op.some }
func (op Option[T]) IsNone() bool { return !op.some }

func (op Option[T]) Or(val Option[T]) Option[T] {
	if op.some {
		return op
	}
	return val
}

func (op Option[T]) OrElse(f func() Option[T]) Option[T] {
	if op.some {
		return op
	}
	return f()
}

func (op Option[T]) UnwrapOr(def T) T {
	if op.some {
		return op.value
	}
	return def
}

func (op Option[T]) UnwrapOrElse(f func() T) T {
	if op.some {
		return op.value
	}
	return f()
}

func (op Option[T]) UnwrapOrZero() T {
	if op.some {
		return op.value
	}
	var t T
	return t
}

func (op *Option[T]) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if len(data) < 1 {
		return nil, encoding.ErrBuffer
	}
	out := Option[T]{
		some: data[0] != 0,
	}
	data = data[1:]

	if out.some {
		data, err = encoding.Decode(data, &out.value, encoding.Ctx(ctx))
		if err != nil {
			return nil, err
		}
	}
	*op = out
	return data, nil
}

func (op *Option[T]) EncodeTZ(ctx *encoding.Context) ([]byte, error) {
	var buf bytes.Buffer
	if op.IsSome() {
		buf.WriteByte(255)
		val := op.Unwrap()
		if err := encoding.Encode(&buf, &val, encoding.Ctx(ctx)); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	} else {
		return []byte{0}, nil
	}
}

func (op Option[T]) String() string {
	if op.some {
		return fmt.Sprintf("Some(%v)", op.value)
	}
	return "None"
}
