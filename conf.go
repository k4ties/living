package living

import (
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Config is world.EntityConfig implementation for Living.
type Config struct {
	EntityType world.EntityType
	Handler    Handler
	Drops      []Drop
	MaxHealth  float64
}

// Apply applies Living entity data to world.EntityData.
func (conf Config) Apply(data *world.EntityData) {
	if conf.EntityType == nil {
		panic("entity type can't be nil")
	}
	if conf.Handler == nil {
		conf.Handler = NopHandler{}
	}

	hm := entity.NewHealthManager(conf.MaxHealth, conf.MaxHealth)
	data.Data = &livingData{
		drops:      conf.Drops,
		entityType: conf.EntityType,
		handler:    conf.Handler,

		HealthManager: hm,
		mc: &entity.MovementComputer{
			Gravity:           0.08,
			Drag:              0.02,
			DragBeforeGravity: true,
		},

		speed: protocol.AbilityBaseWalkSpeed,
	}
}
