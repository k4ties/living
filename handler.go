package living

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type Context = event.Context[*Living]

// Handler is used to handle main Living actions.
type Handler interface {
	// HandleTick handles the entity's tick.
	HandleTick(ctx *Context)
	// HandleHurt handles when entity being hurt.
	HandleHurt(ctx *Context, damage float64, immune bool, immunity *time.Duration, src world.DamageSource)
}

type NopHandler struct{}

func (NopHandler) HandleTick(*Context) {}
func (NopHandler) HandleHurt(*Context, float64, bool, *time.Duration, world.DamageSource) {
}
