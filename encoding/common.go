package encoding

import "strings"

type Context struct {
	enumReg *EnumRegistry
}

func applyOptions(opts []Option) (*Context, []flag) {
	var ctx Context
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
