package encoding

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

type ErrBuffer int

func (err ErrBuffer) Error() string {
	return fmt.Sprintf("buffer is too short, at least %d byte(s) were expected", err)
}

type Decoder interface {
	DecodeTZ(data []byte, ctx *Context) (rest []byte, err error)
}

var be = binary.BigEndian

func decodeInt(data []byte, out reflect.Value, path ErrorPath) (rest []byte, err error) {
	k := out.Kind()
	switch k {
	case reflect.Int8, reflect.Uint8:
		if len(data) < 1 {
			return nil, &Error{path, ErrBuffer(1)}
		}
		if k == reflect.Int8 {
			out.SetInt(int64(data[0]))
		} else {
			out.SetUint(uint64(data[0]))
		}
		return data[1:], nil

	case reflect.Int16, reflect.Uint16:
		if len(data) < 2 {
			return nil, &Error{path, ErrBuffer(2)}
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
			return nil, &Error{path, ErrBuffer(4)}
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
			return nil, &Error{path, ErrBuffer(8)}
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

func decodeBuiltin(data []byte, out reflect.Value, ctx *Context, path ErrorPath) (rest []byte, err error) {
	typ := out.Type()
	k := typ.Kind()
	switch {
	case k == reflect.Bool:
		if len(data) < 1 {
			return nil, &Error{path, ErrBuffer(1)}
		}
		out.SetBool(data[0] != 0)
		return data[1:], nil

	case k >= reflect.Int8 && k <= reflect.Int64 || k >= reflect.Uint8 && k <= reflect.Uint64:
		return decodeInt(data, out, path)

	case k == reflect.String:
		out.Set(reflect.ValueOf(data).Convert(typ))
		return data[len(data):], nil

	case k == reflect.Array:
		l := typ.Len()
		if typ.Elem().Kind() == reflect.Uint8 {
			if len(data) < l {
				return nil, &Error{path, ErrBuffer(l)}
			}
			reflect.Copy(out, reflect.ValueOf(data[:l]))
			return data[l:], nil
		} else {
			path := append(path, nil)
			sel := &path[len(path)-1]
			for i := 0; i < l; i++ {
				*sel = IndexSelector(i)
				if data, err = decodeValue(data, out.Index(i), ctx, nil, path); err != nil {
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
			path := append(path, nil)
			sel := &path[len(path)-1]
			s := reflect.MakeSlice(typ, 0, 0)
			var i int
			for len(data) != 0 {
				*sel = IndexSelector(i)
				tmp := reflect.New(typ.Elem()).Elem()
				s = reflect.Append(s, tmp)
				el := s.Index(s.Len() - 1) // don't lose data even in case of error
				if data, err = decodeValue(data, el, ctx, nil, path); err != nil {
					break
				}
				i += 1
			}
			out.Set(s)
			return data, err
		}

	case k == reflect.Struct:
		path := append(path, nil)
		sel := &path[len(path)-1]
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
			*sel = (*FieldSelector)(&f)
			if data, err = decodeValue(data, field, ctx, fl, path); err != nil {
				return nil, err
			}
		}
		return data, nil

	default:
		return nil, &Error{path, fmt.Errorf("unsupported type %v", typ)}
	}
}

var decoderType = reflect.TypeOf((*Decoder)(nil)).Elem()

func decodeValue(data []byte, out reflect.Value, ctx *Context, fl []flag, path ErrorPath) ([]byte, error) {
	if len(fl) != 0 {
		if _, ok := fl[0].(flDynamic); ok {
			fl = fl[1:]
			// get length
			if len(data) < 4 {
				return nil, &Error{path, ErrBuffer(4)}
			}
			ln := be.Uint32(data)
			data = data[4:]
			if len(data) < int(ln) {
				return nil, &Error{path, ErrBuffer(int(ln))}
			}
			tmp := data[:ln]
			data = data[ln:]
			if _, err := decodeValue(tmp, out, ctx, fl, path); err != nil {
				return nil, err
			}
			return data, nil
		} else if _, ok := fl[0].(flOptional); ok {
			fl = fl[1:]
			if out.Kind() != reflect.Pointer {
				return nil, &Error{path, fmt.Errorf("optional attribute on a non pointer field %v", out.Type())}
			}
			// get flag
			if len(data) < 1 {
				return nil, &Error{path, ErrBuffer(1)}
			}
			some := data[0] != 0
			data = data[1:]
			if some {
				return decodeValue(data, out, ctx, fl, path)
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
				return nil, &Error{path, ErrBuffer(l)}
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
		var dec Decoder
		if out.Type().Implements(decoderType) {
			dec = out.Interface().(Decoder)
		} else if reflect.PtrTo(out.Type()).Implements(decoderType) && out.CanAddr() {
			dec = out.Addr().Interface().(Decoder)
		}
		if dec != nil {
			rest, err := dec.DecodeTZ(data, ctx)
			if err, ok := err.(*Error); ok {
				return nil, &Error{
					Path: append(path, err.Path...),
					Err:  err.Err,
				}
			}
			return rest, wrapError(err, path)
		}
		return decodeBuiltin(data, out, ctx, path)
	}

	// decode enum
	val, rest, err := ctx.enums().tryDecode(out.Type(), data, ctx, path)
	if val.IsValid() {
		out.Set(val)
		return rest, err
	}
	return nil, &Error{path, fmt.Errorf("unsupported interface type %v", out.Type())}
}

func Decode(data []byte, v any, opt ...Option) (rest []byte, err error) {
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
	return decodeValue(data, out, ctx, flags, ErrorPath{})
}
