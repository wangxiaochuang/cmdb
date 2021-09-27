package backbone

import (
    "context"
)

type ccContext struct {
    ctx context.Context
}

type CCContextInterface interface {
    WithCancel() (context.Context, context.CancelFunc)
}

func newCCContext() CCContextInterface {
    return &ccContext{
        ctx: context.Background(),
    }
}

// WithCancel
func (c *ccContext) WithCancel() (context.Context, context.CancelFunc) {
    return context.WithCancel(c.ctx)
}


