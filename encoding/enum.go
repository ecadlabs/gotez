package encoding

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"
)

type enumData struct {
	def      reflect.Type
	variants map[uint8]reflect.Type
}

// EnumRegistry stores variant types
type EnumRegistry struct {
	types map[reflect.Type]*enumData
	mtx   sync.RWMutex
}

func (v *EnumRegistry) RegisterEnum(variants, def any) {
	val := reflect.ValueOf(variants)
	t := val.Type()
	if t.Kind() != reflect.Map || t.Key().Kind() != reflect.Uint8 || t.Elem().Kind() != reflect.Interface {
		panic(fmt.Sprintf("gotez: map[uint8]interface{...} expected: %v", t))
	}
	iftype := t.Elem()
	out := enumData{
		variants: make(map[uint8]reflect.Type, val.Len()),
	}
	iter := val.MapRange()
	for iter.Next() {
		t := iter.Value().Type()
		if !t.Implements(iftype) {
			panic(fmt.Sprintf("gotez: type %v doesn't implement %v", t, iftype))
		}
		out.variants[uint8(iter.Key().Uint())] = iter.Value().Elem().Type()
	}
	if def != nil {
		t := reflect.TypeOf(def)
		if !t.Implements(iftype) {
			panic(fmt.Sprintf("gotez: type %v doesn't implement %v", t, iftype))
		}
		out.def = t
	}
	v.mtx.Lock()
	defer v.mtx.Unlock()
	if _, ok := v.types[t]; ok {
		panic(fmt.Sprintf("gotez: duplicate enum type: %v", iftype))
	}
	v.types[iftype] = &out
}

func (e *EnumRegistry) tryDecode(t reflect.Type, data []byte, ctx *Context, path ErrorPath) (reflect.Value, []byte, error) {
	e.mtx.RLock()
	enum, ok := e.types[t]
	e.mtx.RUnlock()
	if !ok {
		return reflect.Value{}, nil, nil
	}
	if len(data) < 1 {
		return reflect.Value{}, nil, &Error{path, ErrBuffer(1)}
	}
	tag := data[0]
	data = data[1:]
	variant, ok := enum.variants[tag]
	if !ok {
		if enum.def != nil {
			variant = enum.def
		} else {
			return reflect.Value{}, nil, &Error{path, fmt.Errorf("unknown tag %d", tag)}
		}
	}
	path = append(path, TypeSelector{Type: variant})
	val := reflect.New(variant).Elem()
	data, err := decodeValue(data, val, ctx, nil, path)
	return val, data, err
}

func (e *EnumRegistry) tryEncode(out io.Writer, v reflect.Value, ctx *Context, path ErrorPath) (bool, error) {
	e.mtx.RLock()
	enum, ok := e.types[v.Type()]
	e.mtx.RUnlock()
	if !ok {
		return false, nil
	}
	var tag *uint8
	el := v.Elem()
	if !el.IsValid() {
		return false, &Error{path, errors.New("invalid value")}
	}
	for t, typ := range enum.variants {
		if typ == el.Type() {
			tag = &t
			break
		}
	}
	if tag == nil {
		return false, &Error{path, fmt.Errorf("unknown enum variant %v", el.Type())}
	}
	path = append(path, TypeSelector{Type: el.Type()})
	if _, err := out.Write([]byte{*tag}); err != nil {
		return false, wrapError(err, path)
	}
	return true, encodeValue(out, el, ctx, nil, path)
}

func (e *EnumRegistry) ListVariants(typ any) []any {
	t := reflect.TypeOf(typ)
	if t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Interface {
		panic("gotez: pointer to an interface expected")
	}
	t = t.Elem()
	e.mtx.RLock()
	enum, ok := e.types[t]
	e.mtx.RUnlock()
	if !ok {
		return nil
	}
	out := make([]any, 0, len(enum.variants))
	for _, typ := range enum.variants {
		v := reflect.New(typ).Elem().Interface()
		out = append(out, v)
	}
	return out
}

type Variants[T any] map[uint8]T

type Enum[T any] struct {
	Variants Variants[T]
	Default  T
}

// RegisterEnum registers enum type in the global registry
func RegisterEnum[T any](enum *Enum[T]) {
	defaultEnumRegistry.RegisterEnum(enum.Variants, enum.Default)
}

func ListVariants[T any]() []T {
	var typ T
	variants := defaultEnumRegistry.ListVariants(&typ)
	out := make([]T, len(variants))
	for i, v := range variants {
		out[i] = v.(T)
	}
	return out
}

// NewEnumRegistry returns new empty EnumRegistry
func NewEnumRegistry() *EnumRegistry {
	return &EnumRegistry{
		types: make(map[reflect.Type]*enumData),
	}
}

var defaultEnumRegistry = NewEnumRegistry()
