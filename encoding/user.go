package encoding

import (
	"fmt"
	"reflect"
	"sync"
)

// TypeRegistry stores uses interface type constructors (decoders)
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
	ctxType   = reflect.TypeOf((*Context)(nil))
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

func (r *TypeRegistry) RegisterType(fn any) {
	ft := reflect.TypeOf(fn)
	if ft.Kind() != reflect.Func {
		panic(fmt.Sprintf("gotez: function expected: %v", ft))
	}
	if ft.NumIn() != 2 ||
		ft.In(0).Kind() != reflect.Slice ||
		ft.In(0).Elem().Kind() != reflect.Uint8 ||
		ft.In(1) != ctxType ||
		ft.NumOut() != 3 ||
		ft.Out(1).Kind() != reflect.Slice ||
		ft.Out(1).Elem().Kind() != reflect.Uint8 ||
		ft.Out(2) != errorType {
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

func (r *TypeRegistry) tryDecode(t reflect.Type, data []byte, ctx *Context) (reflect.Value, []byte, error) {
	r.mtx.RLock()
	f, ok := r.types[t]
	r.mtx.RUnlock()
	if !ok {
		return reflect.Value{}, nil, nil
	}
	out := reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(data), reflect.ValueOf(ctx)})
	if !out[2].IsNil() {
		return reflect.Value{}, nil, out[2].Interface().(error)
	}
	return out[0], out[1].Interface().([]byte), nil
}

var defaultTypeRegistry = NewTypeRegistry()

// RegisterType registers user interface type in the global registry
func RegisterType[T any](decoder func([]byte, *Context) (T, []byte, error)) {
	defaultTypeRegistry.RegisterType(decoder)
}
