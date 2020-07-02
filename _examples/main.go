package main

import (
	"fmt"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/_examples/engine"
	"github.com/andygeiss/ecs/_examples/plugins"
	"math/rand"
)

const (
	Width  = 800
	Height = 600
)

func generateEntities(num int) []*ecs.Entity {
	out := make([]*ecs.Entity, num)
	for i := range out {
		out[i] = ecs.NewEntity(fmt.Sprintf("%d", i), []ecs.Component{
			engine.NewPosition(rand.Int31()%Width, rand.Int31()%Height),
			engine.NewSize(10, 10),
			engine.NewVelocity(rand.Int31()%10, rand.Int31()%10),
		})
	}
	return out
}

func run() {
	em := ecs.NewEntityManager()
	em.Add(generateEntities(1000)...)
	sm := ecs.NewSystemManager()
	sm.Add(
		engine.NewMovement(),
		engine.NewCollision(Width, Height),
		engine.NewRendering(Width, Height, "ECS with SDL Demo",
			plugins.ShowEngineStats(em),
		),
	)
	ecs.Run(em, sm)
}

func main() {
	ecs.Main(func() {
		run()
	})
}
