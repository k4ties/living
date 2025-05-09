package living

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
)

// NopLivingType is no-operation implementation of world.EntityType.
type NopLivingType struct{}

// Open ...
func (NopLivingType) Open(tx *world.Tx, handle *world.EntityHandle, data *world.EntityData) world.Entity {
	l := &Living{
		livingData: data.Data.(*livingData),
		tx:         tx,
		handle:     handle,
		data:       data,
	}

	return l
}

func (NopLivingType) EncodeEntity() string                        { panic("implement me") }
func (NopLivingType) BBox(world.Entity) cube.BBox                 { return cube.BBox{} }
func (NopLivingType) DecodeNBT(map[string]any, *world.EntityData) {}
func (NopLivingType) EncodeNBT(*world.EntityData) map[string]any  { return map[string]any{} }
