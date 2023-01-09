package systems

import (
	"github.com/andygeiss/ecs/core"
	"github.com/andygeiss/ecs/example/components"
)

// Movement ...
type Movement struct {
	err error
}

func (a *Movement) Error() (err error) {
	return a.err
}

func (a *Movement) Setup() {}

func (a *Movement) Process(em core.EntityManager) (state int) {
	for _, e := range em.FilterByMask(components.MaskPosition | components.MaskVelocity) {
		position := e.Get(components.MaskPosition).(*components.Position)
		velocity := e.Get(components.MaskVelocity).(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
	return core.StateEngineContinue
}

func (a *Movement) Teardown() {}

// NewMovement ...
func NewMovement(em core.EntityManager) core.System {
	return &Movement{}
}
