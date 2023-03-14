package encoding

import (
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

func (e *EnumRegistry) tryDecode(t reflect.Type, data []byte, ctx *Context) (reflect.Value, []byte, error) {
	e.mtx.RLock()
	enum, ok := e.types[t]
	e.mtx.RUnlock()
	if !ok {
		return reflect.Value{}, nil, nil
	}
	if len(data) < 1 {
		return reflect.Value{}, nil, ErrBuffer
	}
	tag := data[0]
	data = data[1:]
	variant, ok := enum.variants[tag]
	if !ok {
		if enum.def != nil {
			variant = enum.def
		} else {
			return reflect.Value{}, nil, fmt.Errorf("gotez: unknown tag %d", tag)
		}
	}
	val := reflect.New(variant).Elem()
	data, err := decodeValue(data, val, ctx, nil)
	return val, data, err
}

func (e *EnumRegistry) tryEncode(out io.Writer, v reflect.Value, ctx *Context) (bool, error) {
	e.mtx.RLock()
	enum, ok := e.types[v.Type()]
	e.mtx.RUnlock()
	if !ok {
		return false, nil
	}
	var tag *uint8
	el := v.Elem()
	for t, typ := range enum.variants {
		if typ == el.Type() {
			tag = &t
			break
		}
	}
	if tag == nil {
		return false, fmt.Errorf("gotez: unknown enum variant %v", el.Type())
	}
	if _, err := out.Write([]byte{*tag}); err != nil {
		return false, err
	}
	return true, encodeValue(out, el, ctx, nil)
}

func (e *EnumRegistry) ForEach(typ any, cb func(tag uint8, v any)) {
	t := reflect.TypeOf(typ)
	if t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Interface {
		panic("gotez: pointer to an interface expected")
	}
	t = t.Elem()
	e.mtx.RLock()
	enum, ok := e.types[t]
	e.mtx.RUnlock()
	if !ok {
		return
	}
	for tag, typ := range enum.variants {
		v := reflect.New(typ).Elem().Interface()
		cb(tag, v)
	}
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

func ForEachInEnum[T any](cb func(tag uint8, v T)) {
	var typ T
	defaultEnumRegistry.ForEach(&typ, func(tag uint8, v any) { cb(tag, v.(T)) })
}

// NewEnumRegistry returns new empty EnumRegistry
func NewEnumRegistry() *EnumRegistry {
	return &EnumRegistry{
		types: make(map[reflect.Type]*enumData),
	}
}

var defaultEnumRegistry = NewEnumRegistry()
