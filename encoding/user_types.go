package encoding

import (
	"fmt"
	"reflect"
	"sync"
)

// TypeRegistry stores interface type constructors (decoders)
type TypeRegistry struct {
	types map[reflect.Type]any
	mtx   sync.RWMutex
}

// NewTypeRegistry returns new empty TypeRegistry
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		types: make(map[reflect.Type]any),
	}
}

var (
	errorType = reflect.TypeOf((*error)(nil)).Elem()
	bytesType = reflect.TypeOf([]byte(nil))
)

// RegisterType registers user interface type. The argument is a constructor function of type `func([]byte) (UserType, []byte, error)`.
// It panics if any other type is passed
func (r *TypeRegistry) RegisterType(fn any) {
	ft := reflect.TypeOf(fn)
	if ft.Kind() != reflect.Func {
		panic(fmt.Sprintf("gotez: function expected: %v", ft))
	}
	if ft.NumIn() != 1 || ft.In(0) != bytesType ||
		ft.NumOut() != 3 || ft.Out(1) != bytesType || ft.Out(2) != errorType {
		panic(fmt.Sprintf("gotez: invalid signature: %v", ft))
	}
	t := ft.Out(0)
	if t.Kind() != reflect.Interface {
		panic(fmt.Sprintf("gotez: user type must be an interface: %v", t))
	}
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.types[t]; ok {
		panic(fmt.Sprintf("gotez: duplicate user type: %v", t))
	}
	r.types[t] = fn
}

func (r *TypeRegistry) tryDecode(t reflect.Type, data []byte) (reflect.Value, []byte, error) {
	r.mtx.RLock()
	f, ok := r.types[t]
	r.mtx.RUnlock()
	if !ok {
		return reflect.Value{}, nil, nil
	}
	out := reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(data)})
	if !out[2].IsNil() {
		return reflect.Value{}, nil, fmt.Errorf("gotez: %w", out[2].Interface().(error))
	}
	return out[0], out[1].Interface().([]byte), nil
}

// RegisterType registers user interface type in the global registry
func RegisterType[T any](fn func([]byte) (T, []byte, error)) {
	defaultTypeRegistry.RegisterType(fn)
}

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

type Variants[T any] map[uint8]T

type Enum[T any] struct {
	Variants Variants[T]
	Default  T
}

// RegisterEnum registers enum type in the global registry
func RegisterEnum[T any](enum *Enum[T]) {
	defaultEnumRegistry.RegisterEnum(enum.Variants, enum.Default)
}

// NewEnumRegistry returns new empty EnumRegistry
func NewEnumRegistry() *EnumRegistry {
	return &EnumRegistry{
		types: make(map[reflect.Type]*enumData),
	}
}

var (
	defaultTypeRegistry = NewTypeRegistry()
	defaultEnumRegistry = NewEnumRegistry()
)
