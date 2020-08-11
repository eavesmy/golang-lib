package context

import "context"
import "time"

type Ctx struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func New() *Ctx {
	return &Ctx{
		ctx: context.Background(),
	}
}

func (c *Ctx) Set(k, v interface{}) *Ctx {
	c.ctx = context.WithValue(c.ctx, k, v)
	return c
}

func (c *Ctx) Get(k interface{}) interface{} {
	return c.ctx.Value(k)
}

func (c *Ctx) SetDeadline(t time.Time) {
	c.ctx, c.cancel = context.WithDeadline(c.ctx, t)
}

func (c *Ctx) Cancel() {
	if c.cancel == nil {
		c.ctx, c.cancel = context.WithCancel(c.ctx)
	}
	c.cancel()
}

func (c *Ctx) SetTimeout(t time.Duration) {
	c.ctx, c.cancel = context.WithTimeout(c.ctx, t)
}
