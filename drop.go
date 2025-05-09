package living

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"math/rand/v2"
)

type Drop struct {
	Item     world.Item
	Max, Min int
}

func (drop Drop) Stack() item.Stack {
	c := rand.IntN(drop.Max-drop.Min) + drop.Min
	if c == 0 {
		return item.Stack{}
	}

	return item.NewStack(drop.Item, c)
}
