package systems

import (
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/example/components"
)

// Collision ...
type Collision struct {
	err    error
	height float32
	width  float32
}

func (a *Collision) Error() error {
	return a.err
}

func (a *Collision) Setup() {}

func (a *Collision) Process(em core.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		if position.X >= a.width || position.X <= 0 {
			velocity.X = -velocity.X
		}
		if position.Y >= a.height || position.Y <= 0 {
			velocity.Y = -velocity.Y
		}
	}
	return core.StateEngineContinue
}

func (a *Collision) Teardown() {}

func (a *Collision) WithHeight(height int) *Collision {
	a.height = float32(height)
	return a
}

func (a *Collision) WithWidth(width int) *Collision {
	a.width = float32(width)
	return a
}

// NewCollision ...
func NewCollision() core.System {
	return &Collision{}
}
