package systems_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ecs/components"
	"github.com/andygeiss/ecs/systems"
	"testing"
)

func TestMovement_Process_Position_Should_Not_Be_Changed_With_Velocity_Y_0(t *testing.T) {
	em := ecs.NewEntityManager()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 0},
		},
	}
	em.Add(player)
	s := systems.NewMovement()
	s.Process(em)
	assert.That(t, player.Components[0].(*components.Position).X, is.Equal(float32(0)))
	assert.That(t, player.Components[0].(*components.Position).Y, is.Equal(float32(0)))
}

func TestMovement_Process_Position_Should_Be_Changed_To_Y_1_With_Velocity_Y_1(t *testing.T) {
	em := ecs.NewEntityManager()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 1},
		},
	}
	em.Add(player)
	s := systems.NewMovement()
	s.Process(em)
	assert.That(t, player.Components[0].(*components.Position).X, is.Equal(float32(0)))
	assert.That(t, player.Components[0].(*components.Position).Y, is.Equal(float32(1)))
}

func TestMovement_Process_Position_Should_Be_Changed_To_Y_Minus_1_With_Velocity_Y_Minus_1(t *testing.T) {
	em := ecs.NewEntityManager()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: -1},
		},
	}
	em.Add(player)
	s := systems.NewMovement()
	s.Process(em)
	assert.That(t, player.Components[0].(*components.Position).X, is.Equal(float32(0)))
	assert.That(t, player.Components[0].(*components.Position).Y, is.Equal(float32(-1)))
}
