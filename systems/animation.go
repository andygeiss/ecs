package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// Animation ...
type Animation struct {
	frameCounter int64
}

// NewAnimation ...
func NewAnimation() ecs.System {
	return &Animation{}
}

// Process ...
func (s *Animation) Process(entityManager *ecs.EntityManager) {
	if rl.WindowShouldClose() {
		ecs.ShouldEngineStop = true
		return
	}
	if ecs.ShouldEnginePause {
		return
	}
	for _, e := range entityManager.FilterBy("animation") {
		animation := e.Get("animation").(*components.Animation)
		if s.frameCounter % (10 - animation.Speed) == 0 {
			animation.SpriteIndex++
			if animation.SpriteIndex == animation.SpriteCount {
				animation.SpriteIndex = 0
			}
		}
	}
	s.frameCounter++
}

// Setup ...
func (s *Animation) Setup() {}

// Teardown ...
func (s *Animation) Teardown() {}
