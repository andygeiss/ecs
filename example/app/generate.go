package app

import (
	"fmt"
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/example/components"
	"math/rand"
)

// Generate ...
func Generate(cfg *Config) []*core.Entity {
	out := make([]*core.Entity, cfg.NumberOfEntities)
	for i := range out {
		out[i] = core.NewEntity(fmt.Sprintf("%d", i), []core.Component{
			components.NewPosition().(*components.Position).
				WithX(rand.Float32() * float32(cfg.Width)).
				WithY(rand.Float32() * float32(cfg.Height)),
			components.NewSize().(*components.Size).
				WithWidth(3).
				WithHeight(3),
			components.NewVelocity().(*components.Velocity).
				WithX(rand.Float32() * 10).
				WithY(rand.Float32() * 10),
		})
	}
	return out
}
