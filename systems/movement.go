package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// Movement ...
type Movement struct{}

// NewMovement ...
func NewMovement() ecs.System {
	return &Movement{}
}

// Process ...
func (s *Movement) Process(entityManager *ecs.EntityManager) {
	if rl.WindowShouldClose() {
		ecs.ShouldEngineStop = true
		return
	}
	if ecs.ShouldEnginePause {
		return
	}
	for _, e := range entityManager.FilterBy("position", "velocity") {
		position := e.Get("position").(*components.Position)
		velocity := e.Get("velocity").(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
}

// Setup ...
func (s *Movement) Setup() {}

// Teardown ...
func (s *Movement) Teardown() {}
