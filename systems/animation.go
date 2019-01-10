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
		if s.frameCounter % 4 == 0 {
			animation := e.Get("animation").(*components.Animation)
			animation.Index++
			if animation.Index == animation.Count {
				animation.Index = 0
			}
		}
	}
	s.frameCounter++
}

// Setup ...
func (s *Animation) Setup() {}

// Teardown ...
func (s *Animation) Teardown() {}
