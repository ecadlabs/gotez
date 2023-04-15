package encoding

import (
	"strings"
)

type variable struct {
	key   any
	value any
}

type Context struct {
	enumReg *EnumRegistry
	typeReg *TypeRegistry
	vars    []variable
}

func NewContext() *Context {
	return &Context{}
}

func (ctx *Context) clone() *Context {
	out := *ctx
	return &out
}

func (ctx *Context) Get(key any) any {
	for i := len(ctx.vars) - 1; i >= 0; i-- {
		if ctx.vars[i].key == key {
			return ctx.vars[i].value
		}
	}
	return nil
}

func (ctx *Context) Set(key any, val any) *Context {
	out := *ctx
	out.vars = append(out.vars, variable{key: key, value: val})
	return &out
}

func (ctx *Context) Enums(er *EnumRegistry) {
	ctx.enumReg = er
}

func (ctx *Context) Types(tr *TypeRegistry) {
	ctx.typeReg = tr
}

func applyOptions(opts []Option) (*Context, []flag) {
	ctx := Context{}
	flags := make([]flag, 0, len(opts))
	for _, fn := range opts {
		fn(&flags, &ctx)
	}
	return &ctx, flags
}

type Option func(fl *[]flag, opt *Context)

func Enums(er *EnumRegistry) func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		c.enumReg = er
	}
}

func Types(tr *TypeRegistry) func(*[]flag, *Context) {
	return func(fl *[]flag, c *Context) {
		c.typeReg = tr
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

func (ctx *Context) enums() *EnumRegistry {
	if ctx.enumReg != nil {
		return ctx.enumReg
	}
	return defaultEnumRegistry
}

func (ctx *Context) types() *TypeRegistry {
	if ctx.typeReg != nil {
		return ctx.typeReg
	}
	return defaultTypeRegistry
}

type flag interface {
	flag()
}

type flOmit struct{}
type flOptional struct{}
type flDynamic struct{}

func (flOmit) flag()     {}
func (flOptional) flag() {}
func (flDynamic) flag()  {}

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
		}
	}
	return out
}
