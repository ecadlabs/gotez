package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
)

type Encoder interface {
	EncodeTZ(ctx *Context) ([]byte, error)
}

var encoderType = reflect.TypeOf((*Encoder)(nil)).Elem()

func encodeInt(out io.Writer, val reflect.Value, path ErrorPath) error {
	k := val.Kind()
	switch k {
	case reflect.Int8, reflect.Uint8:
		var buf [1]byte
		if k == reflect.Int8 {
			buf[0] = byte(val.Int())
		} else {
			buf[0] = byte(val.Uint())
		}
		_, err := out.Write(buf[:])
		return wrapError(err, path)

	case reflect.Int16, reflect.Uint16:
		var buf [2]byte
		if k == reflect.Int16 {
			be.PutUint16(buf[:], uint16(val.Int()))
		} else {
			be.PutUint16(buf[:], uint16(val.Uint()))
		}
		_, err := out.Write(buf[:])
		return wrapError(err, path)

	case reflect.Int32, reflect.Uint32:
		var buf [4]byte
		if k == reflect.Int32 {
			be.PutUint32(buf[:], uint32(val.Int()))
		} else {
			be.PutUint32(buf[:], uint32(val.Uint()))
		}
		_, err := out.Write(buf[:])
		return wrapError(err, path)

	case reflect.Int64, reflect.Uint64:
		var buf [8]byte
		if k == reflect.Int64 {
			be.PutUint64(buf[:], uint64(val.Int()))
		} else {
			be.PutUint64(buf[:], val.Uint())
		}
		_, err := out.Write(buf[:])
		return wrapError(err, path)

	default:
		panic("gotez: unhandled type")
	}
}

func encodeBuiltin(out io.Writer, val reflect.Value, ctx *Context, path ErrorPath) error {
	typ := val.Type()
	k := typ.Kind()
	switch {
	case k == reflect.Bool:
		var tmp [1]byte
		if val.Bool() {
			tmp[0] = 255
		}
		_, err := out.Write(tmp[:])
		return wrapError(err, path)

	case k >= reflect.Int8 && k <= reflect.Int64 || k >= reflect.Uint8 && k <= reflect.Uint64:
		return encodeInt(out, val, path)

	case k == reflect.String:
		_, err := out.Write([]byte(val.String()))
		return wrapError(err, path)

	case k == reflect.Array || k == reflect.Slice:
		if typ.Elem().Kind() == reflect.Uint8 {
			_, err := out.Write(val.Bytes())
			return wrapError(err, path)
		}
		path := append(path, nil)
		sel := &path[len(path)-1]
		for i := 0; i < val.Len(); i++ {
			*sel = IndexSelector(i)
			if err := encodeValue(out, val.Index(i), ctx, nil, path); err != nil {
				return err
			}
		}
		return nil

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
			*sel = (*FieldSelector)(&f)
			if err := encodeValue(out, val.Field(i), ctx, fl, path); err != nil {
				return err
			}
		}
		return nil

	default:
		return &Error{path, fmt.Errorf("unsupported type %v", typ)}
	}
}

func encodeValue(out io.Writer, val reflect.Value, ctx *Context, fl []flag, path ErrorPath) error {
	if len(fl) != 0 {
		if _, ok := fl[0].(flDynamic); ok {
			fl = fl[1:]
			var buf bytes.Buffer
			if err := encodeValue(&buf, val, ctx, fl, path); err != nil {
				return err
			}
			tmp := buf.Bytes()
			// write length
			var sz [4]byte
			be.PutUint32(sz[:], uint32(len(tmp)))
			if _, err := out.Write(sz[:]); err != nil {
				return wrapError(err, path)
			}
			_, err := out.Write(tmp)
			return err
		} else if _, ok := fl[0].(flOptional); ok {
			fl = fl[1:]
			if val.Kind() != reflect.Pointer {
				return &Error{path, fmt.Errorf("optional attribute on a non pointer field %v", val.Type())}
			}

			var buf [1]byte
			if !val.IsNil() {
				buf[0] = 255
				if _, err := out.Write(buf[:]); err != nil {
					return wrapError(err, path)
				}
				return encodeValue(out, val, ctx, fl, path)
			} else {
				_, err := out.Write(buf[:])
				return wrapError(err, path)
			}
		}
	}

	// dereference
	for val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return &Error{path, errors.New("nil pointer")}
		}
		val = val.Elem()
		if t := val.Type(); t.Kind() == reflect.Array && t.Elem().Kind() == reflect.Uint8 {
			// special case for the pointer to a byte array
			_, err := out.Write(val.Bytes())
			return wrapError(err, path)
		}
	}

	// concrete type
	if val.Kind() != reflect.Interface {
		// user type
		if wantsVal := val.Type().Implements(encoderType); wantsVal || reflect.PtrTo(val.Type()).Implements(encoderType) && val.CanAddr() {
			var enc Encoder
			if wantsVal {
				enc = val.Interface().(Encoder)
			} else {
				enc = val.Addr().Interface().(Encoder)
			}
			tmp, err := enc.EncodeTZ(ctx)
			if err, ok := err.(*Error); ok {
				return &Error{
					Path: append(path, err.Path...),
					Err:  err.Err,
				}
			}
			_, err = out.Write(tmp)
			return wrapError(err, path)
		}
		return encodeBuiltin(out, val, ctx, path)
	}

	ok, err := ctx.enums().tryEncode(out, val, ctx, path)
	if err != nil {
		return err
	}
	if !ok {
		return &Error{path, fmt.Errorf("unknown interface type %v", val.Type())}
	}
	return nil
}

func Encode(out io.Writer, v any, opt ...Option) error {
	if v == nil {
		return errors.New("gotez: nil interface")
	}
	val := reflect.ValueOf(v)
	ctx, flags := applyOptions(opt)
	return encodeValue(out, val, ctx, flags, ErrorPath{})
}
