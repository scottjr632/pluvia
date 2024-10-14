package context

import (
	"context"

	"github.com/pluvia/pluvia/logging"
	"github.com/pluvia/pluvia/options"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Context struct {
	ctx context.Context
	pl  *pulumi.Context
	log logging.Logger
}

func WithLogger(log logging.Logger) options.OptionFn[*Context] {
	return func(ctx *Context) *Context {
		ctx.log = log
		return ctx
	}
}

func WithContext(ctx context.Context) options.OptionFn[*Context] {
	return func(c *Context) *Context {
		c.ctx = ctx
		return c
	}
}

func New(pl *pulumi.Context, opts ...options.OptionFn[*Context]) *Context {
	c := &Context{
		ctx: context.Background(),
		pl:  pl,
		log: logging.NewNoopLogger(),
	}
	options.Apply(c)
	return c
}

func (c *Context) Ctx() context.Context {
	return c.ctx
}

func (c *Context) Context() context.Context {
	return c.ctx
}

func (c *Context) Pl() *pulumi.Context {
	return c.pl
}

func (c *Context) Pulumi() *pulumi.Context {
	return c.pl
}

func (c *Context) Log() logging.Logger {
	return c.log
}

func (c *Context) Logger() logging.Logger {
	return c.log
}
