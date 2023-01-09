package components

import "github.com/andygeiss/ecs/core"

// Velocity ...
type Velocity struct {
	ID string  `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}

func (a *Velocity) Mask() uint64 {
	return MaskVelocity
}

func (a *Velocity) WithX(x float32) *Velocity {
	a.X = x
	return a
}

func (a *Velocity) WithY(y float32) *Velocity {
	a.Y = y
	return a
}

// NewVelocity ...
func NewVelocity() core.Component {
	return &Velocity{
		ID: "velocity",
	}
}
