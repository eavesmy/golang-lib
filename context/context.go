package context

import (
	"context"
	"time"
)

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

func (c *Ctx) Background() *Ctx {
	return New()
}

func (c *Ctx) TODO() *Ctx {
	return New()
}

func (c *Ctx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *Ctx) Done() <-chan struct{} {
	return nil
}

func (c *Ctx) Err() error {
	return nil
}

func (c *Ctx) Value(key interface{}) interface{} {
	return nil
}
