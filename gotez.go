package gotez

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var ErrBuffer = errors.New("gotez: buffer is too short")

type Decoder interface {
	DecodeTZ(data []byte, ctx *Context) (rest []byte, err error)
}

type Context struct {
	typeReg *TypeRegistry
	enumReg *EnumRegistry
}

func applyOptions(opts []DecodeOption) (*Context, []flag) {
	var ctx Context
	flags := make([]flag, 0, len(opts))
	for _, fn := range opts {
		fn(&flags, &ctx)
	}
	return &ctx, flags
}

type DecodeOption func(fl *[]flag, opt *Context)

func Types(tr *TypeRegistry) func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		c.typeReg = tr
	}
}

func Enums(er *EnumRegistry) func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		c.enumReg = er
	}
}

func Ctx(ctx *Context) func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		*c = *ctx
	}
}

func Dynamic() func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		*fl = append(*fl, flDynamic{})
	}
}

func Optional() func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		*fl = append(*fl, flOptional{})
	}
}

func (ctx *Context) types() *TypeRegistry {
	if ctx.typeReg != nil {
		return ctx.typeReg
	}
	return defaultTypeRegistry
}

func (ctx *Context) enums() *EnumRegistry {
	if ctx.enumReg != nil {
		return ctx.enumReg
	}
	return defaultEnumRegistry
}

var be = binary.BigEndian

func decodeInt(data []byte, out reflect.Value) (rest []byte, err error) {
	k := out.Kind()
	switch k {
	case reflect.Int8, reflect.Uint8:
		if len(data) < 1 {
			return nil, ErrBuffer
		}
		if k == reflect.Int8 {
			out.SetInt(int64(data[0]))
		} else {
			out.SetUint(uint64(data[0]))
		}
		return data[1:], nil

	case reflect.Int16, reflect.Uint16:
		if len(data) < 2 {
			return nil, ErrBuffer
		}
		v := be.Uint16(data)
		if k == reflect.Int16 {
			out.SetInt(int64(v))
		} else {
			out.SetUint(uint64(v))
		}
		return data[2:], nil

	case reflect.Int32, reflect.Uint32:
		if len(data) < 4 {
			return nil, ErrBuffer
		}
		v := be.Uint32(data)
		if k == reflect.Int32 {
			out.SetInt(int64(v))
		} else {
			out.SetUint(uint64(v))
		}
		return data[4:], nil

	case reflect.Int64, reflect.Uint64:
		if len(data) < 8 {
			return nil, ErrBuffer
		}
		v := be.Uint64(data)
		if k == reflect.Int64 {
			out.SetInt(int64(v))
		} else {
			out.SetUint(uint64(v))
		}
		return data[8:], nil

	default:
		panic("gotez: unhandled type")
	}
}

type flag interface {
	flag()
}

type flOmit struct{}
type flOptional struct{}
type flDynamic struct{}
type flConst string

func (flOmit) flag()     {}
func (flOptional) flag() {}
func (flDynamic) flag()  {}
func (flConst) flag()    {}

func parseTag(tag string) []flag {
	opt := strings.Split(tag, ",")
	out := make([]flag, 0, len(opt))
	for _, o := range opt {
		switch {
		case o == "omit" || o == "-":
			out = append(out, flOmit{})
		case o == "dynamic" || o == "dyn":
			out = append(out, flDynamic{})
		case o == "optional" || o == "opt":
			out = append(out, flOptional{})
		case strings.HasPrefix(o, "const="):
			out = append(out, flConst(strings.SplitN(opt[0], "=", 2)[1]))
		}
	}
	return out
}

func decodeBuiltin(data []byte, out reflect.Value, ctx *Context) (rest []byte, err error) {
	typ := out.Type()
	k := typ.Kind()
	switch {
	case k == reflect.Bool:
		if len(data) < 1 {
			return nil, ErrBuffer
		}
		out.SetBool(data[0] != 0)
		return data[1:], nil

	case k >= reflect.Int8 && k <= reflect.Int64 || k >= reflect.Uint8 && k <= reflect.Uint64:
		return decodeInt(data, out)

	case k == reflect.Array:
		l := typ.Len()
		if typ.Elem().Kind() == reflect.Uint8 {
			if len(data) < l {
				return nil, ErrBuffer
			}
			reflect.Copy(out, reflect.ValueOf(data[:l]))
			return data[l:], nil
		} else {
			for i := 0; i < l; i++ {
				if data, err = decodeValue(data, out.Index(i), ctx, nil); err != nil {
					break
				}
			}
			return data, err
		}

	case k == reflect.Slice:
		if typ.Elem().Kind() == reflect.Uint8 {
			out.Set(reflect.ValueOf(data))
			return data[len(data):], nil
		} else {
			s := reflect.MakeSlice(typ, 0, 0)
			for len(data) != 0 {
				el := reflect.New(typ.Elem()).Elem()
				if data, err = decodeValue(data, el, ctx, nil); err != nil {
					break
				}
				s = reflect.Append(s, el)
			}
			out.Set(s)
			return data, err
		}

	case k == reflect.Struct:
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			if !f.IsExported() {
				continue
			}
			fl := parseTag(f.Tag.Get("tz"))
			if len(fl) != 0 {
				if _, ok := fl[0].(flOmit); ok {
					continue
				}
			}
			field := out.Field(i)
			if data, err = decodeValue(data, field, ctx, fl); err != nil {
				return nil, err
			}
		}
		return data, nil

	default:
		return nil, fmt.Errorf("gotez: unsupported type %v", k)
	}
}

var (
	decoderType = reflect.TypeOf((*Decoder)(nil)).Elem()
)

func decodeValue(data []byte, out reflect.Value, ctx *Context, fl []flag) ([]byte, error) {
	var dec func(data []byte) ([]byte, error)
	dec = func(data []byte) ([]byte, error) {
		if len(fl) != 0 {
			if _, ok := fl[0].(flDynamic); ok {
				fl = fl[1:]
				// get length
				if len(data) < 4 {
					return nil, ErrBuffer
				}
				ln := be.Uint32(data)
				data = data[4:]
				if len(data) < int(ln) {
					return nil, ErrBuffer
				}
				tmp := data[:ln]
				data = data[ln:]
				if _, err := dec(tmp); err != nil {
					return nil, err
				}
				return data, nil
			} else if _, ok := fl[0].(flOptional); ok {
				fl = fl[1:]
				if out.Kind() != reflect.Pointer {
					return nil, fmt.Errorf("gotez: optional attribute on a non pointer field %v", out.Type())
				}
				// get flag
				if len(data) < 1 {
					return nil, ErrBuffer
				}
				some := data[0] != 0
				data = data[1:]
				if some {
					return dec(data)
				} else {
					return data, nil
				}
			}
		}

		// out must be a non pointer
		for out.Kind() == reflect.Pointer {
			el := out.Type().Elem()
			if el.Kind() == reflect.Array && el.Elem().Kind() == reflect.Uint8 {
				// special case for the pointer to a byte array
				l := el.Len()
				if len(data) < l {
					return nil, ErrBuffer
				}
				out.Set(reflect.ValueOf(data[:l]).Convert(out.Type()))
				return data[l:], nil
			}
			if out.IsNil() {
				out.Set(reflect.New(el))
			}
			out = out.Elem()
		}

		// concrete type
		if out.Kind() != reflect.Interface {
			// user type
			if reflect.PtrTo(out.Type()).Implements(decoderType) && out.CanAddr() {
				dec := out.Addr().Interface().(Decoder)
				return dec.DecodeTZ(data, ctx)
			}
			return decodeBuiltin(data, out, ctx)
		}

		// user interface type
		val, rest, err := ctx.types().tryDecode(out.Type(), data)
		if err != nil {
			return nil, err
		}
		if !val.IsValid() {
			// decode enum
			if val, rest, err = ctx.enums().tryDecode(out.Type(), data, ctx); err != nil {
				return nil, err
			}
			if !val.IsValid() {
				return nil, fmt.Errorf("gotez: unsupported interface type %v", out.Type())
			}
		}
		out.Set(val)
		return rest, nil
	}

	rest, err := dec(data)
	if err != nil {
		return nil, err
	}
	if len(fl) != 0 {
		if cv, ok := fl[0].(flConst); ok {
			switch out.Kind() {
			case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				val, err := strconv.ParseUint(string(cv), 0, 64)
				if err != nil {
					return nil, fmt.Errorf("gotez: %w", err)
				}
				if out.Uint() != val {
					return nil, fmt.Errorf("gotez: const field is expected to be %d, got %d", val, out.Uint())
				}

			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				val, err := strconv.ParseInt(string(cv), 0, 64)
				if err != nil {
					return nil, fmt.Errorf("gotez: %w", err)
				}
				if out.Int() != val {
					return nil, fmt.Errorf("gotez: const field is expected to be %d, got %d", val, out.Uint())
				}
			}
		}
	}
	return rest, nil
}

func Decode(data []byte, v any, opt ...DecodeOption) (rest []byte, err error) {
	if v == nil {
		return nil, errors.New("gotez: nil interface")
	}
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("gotez: pointer expected: %v", val.Type())
	}
	if val.IsNil() {
		return nil, errors.New("gotez: nil pointer")
	}
	out := val.Elem()
	ctx, flags := applyOptions(opt)
	return decodeValue(data, out, ctx, flags)
}
